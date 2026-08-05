// Pure-Go secretstream writer (no CGO). WAL-G framing: 8192-byte chunks.

package secretstream

import (
	"fmt"
	"io"
	"sync"
)

// Writer encrypts plaintext into a secretstream (header + tagged chunks).
type Writer struct {
	io.Writer

	state *streamState

	in []byte

	inIdx int

	onceHeader sync.Once

	key []byte

	headerErr error
}

// NewWriter creates an encrypting WriteCloser. key must be KeyBytes long.
// Close must be called to emit TAG_FINAL.
func NewWriter(writer io.Writer, key []byte) io.WriteCloser {
	return &Writer{
		Writer: writer,
		in:     make([]byte, ChunkSize),
		key:    key,
	}
}

func (writer *Writer) writeHeader() {
	header := make([]byte, HeaderBytes)
	st, err := initPush(writer.key, header)
	if err != nil {
		writer.headerErr = fmt.Errorf("secretstream writer: init_push failed: %w", err)
		return
	}
	if _, err := writer.Writer.Write(header); err != nil {
		writer.headerErr = fmt.Errorf("secretstream writer: header write failed: %w", err)
		return
	}
	writer.state = st
}

// Write implements io.Writer.
func (writer *Writer) Write(p []byte) (n int, err error) {
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

func (writer *Writer) writeNextChunk(last bool) (err error) {
	tag := byte(TagMessage)
	if last {
		tag = TagFinal
	}
	wire, err := writer.state.push(writer.in[:writer.inIdx], tag)
	if err != nil {
		return fmt.Errorf("secretstream writer: push failed (plain_len=%d final=%v): %w", writer.inIdx, last, err)
	}
	if _, err = writer.Writer.Write(wire); err != nil {
		return fmt.Errorf("secretstream writer: wire write failed (wire_len=%d final=%v): %w", len(wire), last, err)
	}
	writer.inIdx = 0
	return
}

// Close implements io.Closer — emits the final tagged chunk.
func (writer *Writer) Close() (err error) {
	if closer, ok := writer.Writer.(io.Closer); ok {
		defer closer.Close()
	}
	writer.onceHeader.Do(writer.writeHeader)
	if writer.headerErr != nil {
		return writer.headerErr
	}
	return writer.writeNextChunk(true)
}
