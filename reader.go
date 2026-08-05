// Pure-Go secretstream reader (no CGO). WAL-G framing: 8192-byte chunks.

package secretstream

import (
	"fmt"
	"io"
	"sync"
)

// Reader decrypts a secretstream produced by Writer (or libsodium C with the
// same framing).
type Reader struct {
	io.Reader

	state *streamState

	in []byte

	out []byte

	outIdx int

	outLen int

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

// Read implements io.Reader.
func (reader *Reader) Read(p []byte) (n int, err error) {
	reader.onceHeader.Do(reader.readHeader)
	if reader.headerErr != nil {
		return 0, reader.headerErr
	}
	if reader.outIdx >= reader.outLen {
		if reader.done {
			return 0, io.EOF
		}
		if err = reader.readNextChunk(); err != nil {
			return
		}
	}
	n = copy(p, reader.out[reader.outIdx:reader.outLen])
	reader.outIdx += n
	return
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
	if tag == TagFinal && err != io.ErrUnexpectedEOF {
		return fmt.Errorf("secretstream reader: TAG_FINAL on full wire chunk (%d bytes) — framing anomaly (premature end)", n)
	}
	if tag == TagFinal {
		reader.done = true
	}
	if len(plain) > len(reader.out) {
		reader.out = make([]byte, len(plain))
	}
	copy(reader.out, plain)
	reader.outIdx = 0
	reader.outLen = len(plain)
	return nil
}
