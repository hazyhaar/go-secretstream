// Modified in the lateos-ai/wal-g fork. Pure-Go secretstream writer (no CGO).

package secretstream

import (
	"fmt"
	"io"
	"sync"
)

// Writer wraps ordinary writer with libsodium encryption.
// Close is idempotent. Write after Close returns an error.
type Writer struct {
	io.Writer

	state *streamState

	in []byte

	wireBuf []byte

	inIdx int

	// Header is written lazily on first Write/Close so Pipe consumers can
	// attach the reader before any bytes are produced.
	onceHeader sync.Once

	key []byte

	headerErr error

	closed   bool
	closeErr error
}

// NewWriter creates Writer from ordinary writer and key.
// key is copied; the caller's slice is not retained.
func NewWriter(writer io.Writer, key []byte) io.WriteCloser {
	k := make([]byte, len(key))
	copy(k, key)
	return &Writer{
		Writer:  writer,
		in:      make([]byte, ChunkSize),
		wireBuf: make([]byte, ChunkSize+ABytes),
		key:     k,
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

// Write implements io.Writer
func (writer *Writer) Write(p []byte) (n int, err error) {
	if writer.closed {
		return 0, fmt.Errorf("secretstream writer: write after close")
	}
	writer.onceHeader.Do(writer.writeHeader)
	if writer.headerErr != nil {
		return 0, writer.headerErr
	}
	for n != len(p) {
		// Zero-copy fast path: when inIdx == 0 and remaining bytes >= ChunkSize,
		// encrypt directly from p slice into wireBuf without copying to writer.in
		if writer.inIdx == 0 && len(p)-n >= ChunkSize {
			if err = writer.writeNextChunkFrom(p[n:n+ChunkSize], false); err != nil {
				return
			}
			n += ChunkSize
			continue
		}
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

func (writer *Writer) writeNextChunkFrom(m []byte, last bool) (err error) {
	tag := byte(TagMessage)
	if last {
		tag = TagFinal
	}
	wireLen, err := writer.state.pushTo(m, tag, writer.wireBuf)
	if err != nil {
		return fmt.Errorf("secretstream writer: push failed (plain_len=%d final=%v): %w", len(m), last, err)
	}
	n, err := writer.Writer.Write(writer.wireBuf[:wireLen])
	if err != nil {
		return fmt.Errorf("secretstream writer: wire write failed (wire_len=%d final=%v): %w", wireLen, last, err)
	}
	if n != wireLen {
		return fmt.Errorf("secretstream writer: wire short write (%d/%d final=%v)", n, wireLen, last)
	}
	return nil
}

func (writer *Writer) writeNextChunk(last bool) (err error) {
	err = writer.writeNextChunkFrom(writer.in[:writer.inIdx], last)
	if err == nil {
		writer.inIdx = 0
	}
	return
}

// Close flushes final chunk and wipes state. Idempotent.
func (writer *Writer) Close() error {
	if writer.closed {
		return writer.closeErr
	}
	writer.closed = true
	writer.onceHeader.Do(writer.writeHeader)
	if writer.headerErr != nil {
		writer.closeErr = writer.headerErr
		return writer.closeErr
	}
	err := writer.writeNextChunk(true)
	if writer.state != nil {
		writer.state.wipe()
	}
	memzero(writer.key)
	memzero(writer.in)
	if err != nil {
		writer.closeErr = fmt.Errorf("secretstream writer: close write failed: %w", err)
	}
	if closer, ok := writer.Writer.(io.Closer); ok {
		if cerr := closer.Close(); writer.closeErr == nil {
			writer.closeErr = cerr
		}
	}
	return writer.closeErr
}
