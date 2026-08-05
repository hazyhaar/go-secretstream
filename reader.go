// Modified in the lateos-ai/wal-g fork. Pure-Go secretstream reader (no CGO).

package secretstream

import (
	"fmt"
	"io"
	"sync"
)

// Reader wraps ordinary reader with libsodium decryption.
//
// WAL-G framing: non-final chunks are exactly ChunkSize plaintext (+ ABytes wire);
// the last chunk carries TAG_FINAL and may be shorter. TAG_FINAL on a full-sized
// wire chunk is rejected (legacy premature-end). Empty non-final short chunks are
// not demuxed by this framer (fixed read size).
//
// Read never returns (0, nil). After MAC failure abandon the Reader (state is not
// advanced on mismatch). Header bytes are unauthenticated alone (libsodium design).
type Reader struct {
	io.Reader

	state *streamState

	in []byte

	out []byte

	outIdx int

	outLen int

	// Header is read lazily on first Read so Pipe producers can write after attach.
	onceHeader sync.Once

	key []byte

	headerErr error

	done bool
}

// NewReader creates Reader from ordinary reader and key.
// key is copied; the caller's slice is not retained.
func NewReader(reader io.Reader, key []byte) io.Reader {
	k := make([]byte, len(key))
	copy(k, key)
	return &Reader{
		Reader: reader,
		in:     make([]byte, ChunkSize+ABytes),
		out:    make([]byte, ChunkSize),
		key:    k,
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
		// Zero-copy fast path: when caller's buffer p can hold a full chunk (>= ChunkSize),
		// decrypt directly into p without copying to reader.out
		if len(p) >= ChunkSize {
			plainLen, rerr := reader.readNextChunkTo(p)
			if plainLen > 0 {
				return plainLen, nil
			}
			if rerr != nil {
				return 0, rerr
			}
		}
		if err = reader.readNextChunk(); err != nil {
			return 0, err
		}
	}
}

func (reader *Reader) readNextChunkTo(dst []byte) (int, error) {
	n, err := io.ReadFull(reader.Reader, reader.in)
	if err == io.EOF || n == 0 {
		if reader.done {
			return 0, io.EOF
		}
		return 0, io.ErrUnexpectedEOF
	}
	if err != nil && err != io.ErrUnexpectedEOF {
		return 0, fmt.Errorf("secretstream reader: chunk read failed (buf=%d got=%d): %w", len(reader.in), n, err)
	}
	plainLen, tag, perr := reader.state.pullTo(reader.in[:n], dst)
	if perr != nil {
		return 0, fmt.Errorf("secretstream reader: decrypt chunk (wire_len=%d): %w", n, perr)
	}
	// Legacy parity: FINAL on a full-sized wire chunk is "premature end".
	if tag == TagFinal && err != io.ErrUnexpectedEOF {
		return 0, fmt.Errorf("secretstream reader: TAG_FINAL on full wire chunk (%d bytes) — framing anomaly (premature end)", n)
	}
	if tag == TagFinal {
		reader.done = true
	}
	return plainLen, nil
}

func (reader *Reader) readNextChunk() error {
	plainLen, err := reader.readNextChunkTo(reader.out)
	if err == nil {
		reader.outIdx = 0
		reader.outLen = plainLen
	}
	return err
}
