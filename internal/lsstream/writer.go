// SPDX-License-Identifier: Apache-2.0

// Pure-Go libsodium secretstream writer (ported from wal-g fork, Apache-2.0).
package lsstream

import (
	"fmt"
	"io"
	"sync"
)

// Writer encrypts with crypto_secretstream_xchacha20poly1305 wire layout.
// Close is idempotent. Write after Close returns an error.
type Writer struct {
	io.Writer

	state      *streamState
	in         []byte
	wireBuf    []byte
	inIdx      int
	onceHeader sync.Once
	key        []byte
	headerErr  error
	stickyErr  error
	closed     bool
	closeErr   error
}

// NewWriter creates a Writer. key is copied.
func NewWriter(writer io.Writer, key []byte) io.WriteCloser {
	k := make([]byte, len(key))
	copy(k, key)
	return &Writer{
		Writer:  writer,
		in:      make([]byte, chunkSize),
		wireBuf: make([]byte, chunkSize+ABytes),
		key:     k,
	}
}

func (w *Writer) writeHeader() {
	header := make([]byte, HeaderBytes)
	st, err := initPush(w.key, header)
	if err != nil {
		w.headerErr = fmt.Errorf("lsstream writer: init_push: %w", err)
		return
	}
	n, err := w.Writer.Write(header)
	if err != nil {
		w.headerErr = fmt.Errorf("lsstream writer: header write: %w", err)
		return
	}
	if n != len(header) {
		w.headerErr = fmt.Errorf("lsstream writer: header short write (%d/%d)", n, len(header))
		return
	}
	w.state = st
}

// Write implements io.Writer.
func (w *Writer) Write(p []byte) (n int, err error) {
	if w.closed {
		return 0, fmt.Errorf("lsstream writer: write after close")
	}
	if w.stickyErr != nil {
		return 0, w.stickyErr
	}
	w.onceHeader.Do(w.writeHeader)
	if w.headerErr != nil {
		w.stickyErr = w.headerErr
		return 0, w.headerErr
	}
	for n != len(p) {
		if w.inIdx == 0 && len(p)-n >= chunkSize {
			if err = w.writeNextChunkFrom(p[n:n+chunkSize], false); err != nil {
				w.stickyErr = err
				return
			}
			n += chunkSize
			continue
		}
		count := copy(w.in[w.inIdx:], p[n:])
		w.inIdx += count
		n += count
		if w.inIdx == len(w.in) {
			if err = w.writeNextChunk(false); err != nil {
				w.stickyErr = err
				return
			}
		}
	}
	return
}

func (w *Writer) writeNextChunkFrom(m []byte, last bool) error {
	tag := byte(TagMessage)
	if last {
		tag = TagFinal
	}
	wireLen, err := w.state.pushTo(m, tag, w.wireBuf)
	if err != nil {
		return fmt.Errorf("lsstream writer: push plain_len=%d final=%v: %w", len(m), last, err)
	}
	n, err := w.Writer.Write(w.wireBuf[:wireLen])
	if err != nil {
		return fmt.Errorf("lsstream writer: wire write wire_len=%d final=%v: %w", wireLen, last, err)
	}
	if n != wireLen {
		return fmt.Errorf("lsstream writer: wire short write (%d/%d final=%v)", n, wireLen, last)
	}
	return nil
}

func (w *Writer) writeNextChunk(last bool) error {
	err := w.writeNextChunkFrom(w.in[:w.inIdx], last)
	if err == nil {
		w.inIdx = 0
	}
	return err
}

// Close implements io.Closer.
func (w *Writer) Close() error {
	if w.closed {
		return w.closeErr
	}
	w.closed = true
	if w.stickyErr != nil {
		w.closeErr = w.stickyErr
	} else {
		w.onceHeader.Do(w.writeHeader)
		if w.headerErr != nil {
			w.closeErr = w.headerErr
		} else {
			w.closeErr = w.writeNextChunk(true)
			if w.closeErr != nil {
				w.stickyErr = w.closeErr
			}
		}
	}
	if w.state != nil {
		w.state.wipe()
		w.state = nil
	}
	memzero(w.key)
	memzero(w.in)
	if closer, ok := w.Writer.(io.Closer); ok {
		if cerr := closer.Close(); w.closeErr == nil {
			w.closeErr = cerr
		}
	}
	return w.closeErr
}
