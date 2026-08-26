// SPDX-License-Identifier: Apache-2.0

// Pure-Go libsodium secretstream reader (ported from wal-g fork, Apache-2.0).
package lsstream

import (
	"fmt"
	"io"
	"sync"
)

// Reader decrypts WAL-G / libsodium framing:
// non-final chunks are exactly chunkSize plaintext (+ ABytes wire);
// last chunk carries TAG_FINAL and may be shorter.
type Reader struct {
	io.Reader

	state      *streamState
	in         []byte
	out        []byte
	outIdx     int
	outLen     int
	onceHeader sync.Once
	key        []byte
	headerErr  error
	stickyErr  error
	done       bool
}

// NewReader creates a Reader. key is copied.
func NewReader(reader io.Reader, key []byte) io.Reader {
	k := make([]byte, len(key))
	copy(k, key)
	return &Reader{
		Reader: reader,
		in:     make([]byte, chunkSize+ABytes),
		out:    make([]byte, chunkSize),
		key:    k,
	}
}

func (r *Reader) readHeader() {
	header := make([]byte, HeaderBytes)
	if _, err := io.ReadFull(r.Reader, header); err != nil {
		r.headerErr = fmt.Errorf("lsstream reader: header read need %d: %w", HeaderBytes, err)
		return
	}
	st, err := initPull(r.key, header)
	if err != nil {
		r.headerErr = fmt.Errorf("lsstream reader: header reject: %w", err)
		return
	}
	r.state = st
}

// Read implements io.Reader. Never returns (0, nil).
func (r *Reader) Read(p []byte) (n int, err error) {
	if r.stickyErr != nil {
		return 0, r.stickyErr
	}
	r.onceHeader.Do(r.readHeader)
	if r.headerErr != nil {
		r.stickyErr = r.headerErr
		return 0, r.headerErr
	}
	for {
		if r.outIdx < r.outLen {
			n = copy(p, r.out[r.outIdx:r.outLen])
			r.outIdx += n
			return n, nil
		}
		if r.done {
			return 0, io.EOF
		}
		if len(p) >= chunkSize {
			plainLen, rerr := r.readNextChunkTo(p)
			if plainLen > 0 {
				return plainLen, nil
			}
			if rerr != nil {
				r.stickyErr = rerr
				return 0, rerr
			}
		}
		if err = r.readNextChunk(); err != nil {
			r.stickyErr = err
			return 0, err
		}
	}
}

func (r *Reader) readNextChunkTo(dst []byte) (int, error) {
	n, err := io.ReadFull(r.Reader, r.in)
	if err == io.EOF || n == 0 {
		if r.done {
			return 0, io.EOF
		}
		return 0, io.ErrUnexpectedEOF
	}
	if err != nil && err != io.ErrUnexpectedEOF {
		return 0, fmt.Errorf("lsstream reader: chunk read buf=%d got=%d: %w", len(r.in), n, err)
	}
	plainLen, tag, perr := r.state.pullTo(r.in[:n], dst)
	if perr != nil {
		return 0, fmt.Errorf("lsstream reader: decrypt wire_len=%d: %w", n, perr)
	}
	// Libsodium allows TAG_FINAL on a full-sized last data chunk (no trailing empty FINAL).
	// WAL-G historically emitted MESSAGE full + empty FINAL; both are accepted.
	if tag == TagFinal {
		r.done = true
	}
	return plainLen, nil
}

func (r *Reader) readNextChunk() error {
	plainLen, err := r.readNextChunkTo(r.out)
	if err == nil {
		r.outIdx = 0
		r.outLen = plainLen
	}
	return err
}
