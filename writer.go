// Pure-Go secretstream writer (no CGO). WAL-G framing.
// SoT: lateos-ai/wal-g internal/crypto/libsodium.

package secretstream

import (
	"fmt"
	"io"
	"sync"
)

// Writer encrypts plaintext into a secretstream.
// Close is idempotent. Write after Close returns an error.
type Writer struct {
	io.Writer
	state      *streamState
	in         []byte
	inIdx      int
	onceHeader sync.Once
	key        []byte
	headerErr  error
	closed     bool
	closeErr   error
}

// NewWriter creates an encrypting WriteCloser. key is copied. Close emits TAG_FINAL.
func NewWriter(writer io.Writer, key []byte) io.WriteCloser {
	k := make([]byte, len(key))
	copy(k, key)
	return &Writer{
		Writer: writer,
		in:     make([]byte, ChunkSize),
		key:    k,
	}
}

func (writer *Writer) writeHeader() {
	header := make([]byte, HeaderBytes)
	st, err := initPush(writer.key, header)
	if err != nil {
		writer.headerErr = fmt.Errorf("secretstream writer: init_push failed: %w", err)
		return
	}
	n, err := writer.Writer.Write(header)
	if err != nil {
		writer.headerErr = fmt.Errorf("secretstream writer: header write failed: %w", err)
		return
	}
	if n != len(header) {
		writer.headerErr = fmt.Errorf("secretstream writer: header short write (%d/%d)", n, len(header))
		return
	}
	writer.state = st
}

// Write implements io.Writer.
func (writer *Writer) Write(p []byte) (n int, err error) {
	if writer.closed {
		return 0, fmt.Errorf("secretstream writer: write after close")
	}
	writer.onceHeader.Do(writer.writeHeader)
	if writer.headerErr != nil {
		return 0, writer.headerErr
	}
	for n != len(p) {
		count := copy(writer.in[writer.inIdx:], p[n:])
		writer.inIdx += count
		n += count
		if writer.inIdx == len(writer.in) {
			if err = writer.writeNextChunk(false); err != nil {
				return
			}
		}
	}
	return
}

func (writer *Writer) writeNextChunk(last bool) error {
	tag := byte(TagMessage)
	if last {
		tag = TagFinal
	}
	wire, err := writer.state.push(writer.in[:writer.inIdx], tag)
	if err != nil {
		return fmt.Errorf("secretstream writer: push failed (plain_len=%d final=%v): %w", writer.inIdx, last, err)
	}
	n, err := writer.Writer.Write(wire)
	if err != nil {
		return fmt.Errorf("secretstream writer: wire write failed (wire_len=%d final=%v): %w", len(wire), last, err)
	}
	if n != len(wire) {
		return fmt.Errorf("secretstream writer: wire short write (%d/%d final=%v)", n, len(wire), last)
	}
	writer.inIdx = 0
	return nil
}

// Close implements io.Closer. Idempotent. Scrubs stream state and key copy.
func (writer *Writer) Close() error {
	if writer.closed {
		return writer.closeErr
	}
	writer.closed = true
	writer.onceHeader.Do(writer.writeHeader)
	if writer.headerErr != nil {
		writer.closeErr = writer.headerErr
	} else {
		writer.closeErr = writer.writeNextChunk(true)
	}
	if writer.state != nil {
		writer.state.wipe()
		writer.state = nil
	}
	memzero(writer.key)
	memzero(writer.in)
	if closer, ok := writer.Writer.(io.Closer); ok {
		if cerr := closer.Close(); writer.closeErr == nil {
			writer.closeErr = cerr
		}
	}
	return writer.closeErr
}
