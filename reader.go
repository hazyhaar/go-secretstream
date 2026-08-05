// Pure-Go secretstream reader (no CGO). WAL-G framing: 8192-byte chunks.

package secretstream

import (
	"fmt"
	"io"
	"sync"
)

// Reader decrypts a secretstream produced by Writer (or libsodium C with the
// same framing: full chunks are ChunkSize plaintext + ABytes wire overhead;
// the last chunk carries TAG_FINAL and may be shorter).
//
// Empty non-final chunks (MESSAGE/PUSH with zero plaintext) are skipped so
// Read never returns (0, nil). After a MAC failure the stream is unusable;
// callers must abandon the Reader (state is not advanced on MAC mismatch).
//
// Header bytes are not authenticated on their own (libsodium design): corruption
// is detected at the first chunk MAC.
type Reader struct {
	io.Reader

	state *streamState

	in []byte

	out []byte

	outIdx int

	outLen int

	// Header is read lazily on the first Read so Pipe producers can write
	// the header after the reader is attached.
	onceHeader sync.Once

	key []byte

	headerErr error

	done bool
}

// NewReader creates a decrypting Reader. key must be KeyBytes long.
func NewReader(reader io.Reader, key []byte) io.Reader {
	return &Reader{
		Reader: reader,
		in:     make([]byte, ChunkSize+ABytes),
		out:    make([]byte, ChunkSize),
		key:    key,
	}
}

func (reader *Reader) readHeader() {
	header := make([]byte, HeaderBytes)
	if _, err := io.ReadFull(reader.Reader, header); err != nil {
		reader.headerErr = fmt.Errorf("secretstream reader: header read failed (need %d bytes): %w", HeaderBytes, err)
		return
	}
	st, err := initPull(reader.key, header)
	if err != nil {
		reader.headerErr = fmt.Errorf("secretstream reader: header reject: %w", err)
		return
	}
	reader.state = st
}

// Read implements io.Reader. Never returns (0, nil).
func (reader *Reader) Read(p []byte) (n int, err error) {
	reader.onceHeader.Do(reader.readHeader)
	if reader.headerErr != nil {
		return 0, reader.headerErr
	}
	for {
		if reader.outIdx < reader.outLen {
			n = copy(p, reader.out[reader.outIdx:reader.outLen])
			reader.outIdx += n
			return n, nil
		}
		if reader.done {
			return 0, io.EOF
		}
		if err = reader.readNextChunk(); err != nil {
			return 0, err
		}
		// empty non-final plaintext: loop for next chunk
	}
}

func (reader *Reader) readNextChunk() error {
	n, err := io.ReadFull(reader.Reader, reader.in)
	if err != nil && err != io.ErrUnexpectedEOF {
		return fmt.Errorf("secretstream reader: chunk read failed (buf=%d got=%d): %w", len(reader.in), n, err)
	}
	if n == 0 {
		return io.EOF
	}
	plain, tag, perr := reader.state.pull(reader.in[:n])
	if perr != nil {
		return fmt.Errorf("secretstream reader: decrypt chunk (wire_len=%d): %w", n, perr)
	}
	// WAL-G framing: TAG_FINAL only on a short (or empty) last wire chunk.
	// A full-sized FINAL is rejected — foreign producers that finalize on an
	// exact ChunkSize boundary need a trailing empty FINAL (libsodium-friendly)
	// or must use a non-WAL-G framer.
	if tag == TagFinal && err != io.ErrUnexpectedEOF {
		return fmt.Errorf("secretstream reader: TAG_FINAL on full wire chunk (%d bytes) — framing anomaly (premature end)", n)
	}
	if tag == TagFinal {
		reader.done = true
	}
	// plain fits ChunkSize by framing; grow only if a foreign full-chunk FINAL
	// path is ever relaxed.
	if len(plain) > len(reader.out) {
		reader.out = make([]byte, len(plain))
	}
	copy(reader.out, plain)
	reader.outIdx = 0
	reader.outLen = len(plain)
	return nil
}
