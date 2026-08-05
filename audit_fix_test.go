package secretstream

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriterCloseIdempotent(t *testing.T) {
	key := bytes.Repeat([]byte{0x11}, KeyBytes)
	var buf bytes.Buffer
	w := NewWriter(&buf, key)
	_, err := w.Write([]byte("once"))
	require.NoError(t, err)
	require.NoError(t, w.Close())
	nBefore := buf.Len()
	require.NoError(t, w.Close())
	require.NoError(t, w.Close())
	require.Equal(t, nBefore, buf.Len(), "second Close must not append another TAG_FINAL")

	out, err := io.ReadAll(NewReader(bytes.NewReader(buf.Bytes()), key))
	require.NoError(t, err)
	require.Equal(t, []byte("once"), out)
}

func TestWriterWriteAfterClose(t *testing.T) {
	key := bytes.Repeat([]byte{0x22}, KeyBytes)
	w := NewWriter(io.Discard, key)
	require.NoError(t, w.Close())
	_, err := w.Write([]byte("nope"))
	require.Error(t, err)
	require.Contains(t, err.Error(), "write after close")
}

func TestCoreEmptyMessageChunk(t *testing.T) {
	key := bytes.Repeat([]byte{0x33}, KeyBytes)
	header := make([]byte, HeaderBytes)
	pushSt, err := initPush(key, header)
	require.NoError(t, err)
	pullSt, err := initPull(key, header)
	require.NoError(t, err)

	wEmpty, err := pushSt.push(nil, TagMessage)
	require.NoError(t, err)
	require.Equal(t, ABytes, len(wEmpty))
	m, tag, err := pullSt.pull(wEmpty)
	require.NoError(t, err)
	require.Equal(t, byte(TagMessage), tag)
	require.Empty(t, m)

	wFinal, err := pushSt.push([]byte("after-empty"), TagFinal)
	require.NoError(t, err)
	m, tag, err = pullSt.pull(wFinal)
	require.NoError(t, err)
	require.Equal(t, byte(TagFinal), tag)
	require.Equal(t, []byte("after-empty"), m)
}

func TestReaderNeverReturnsZeroNil(t *testing.T) {
	key := bytes.Repeat([]byte{0x34}, KeyBytes)
	var buf bytes.Buffer
	w := NewWriter(&buf, key)
	_, _ = w.Write([]byte("z"))
	_ = w.Close()
	r := NewReader(bytes.NewReader(buf.Bytes()), key)
	p := make([]byte, 64)
	n, err := r.Read(p)
	require.NoError(t, err)
	require.Greater(t, n, 0)
	n, err = r.Read(p)
	require.Equal(t, 0, n)
	require.ErrorIs(t, err, io.EOF)
}

func TestCoreTagPushAndRekey(t *testing.T) {
	key := bytes.Repeat([]byte{0x44}, KeyBytes)
	header := make([]byte, HeaderBytes)
	pushSt, err := initPush(key, header)
	require.NoError(t, err)
	pullSt, err := initPull(key, header)
	require.NoError(t, err)

	w1, err := pushSt.push([]byte("a"), TagMessage)
	require.NoError(t, err)
	w2, err := pushSt.push([]byte("b"), TagPush)
	require.NoError(t, err)
	w3, err := pushSt.push([]byte("c"), TagRekey)
	require.NoError(t, err)
	w4, err := pushSt.push([]byte("d"), TagFinal)
	require.NoError(t, err)

	m, tag, err := pullSt.pull(w1)
	require.NoError(t, err)
	require.Equal(t, byte(TagMessage), tag)
	require.Equal(t, []byte("a"), m)

	m, tag, err = pullSt.pull(w2)
	require.NoError(t, err)
	require.Equal(t, byte(TagPush), tag)
	require.Equal(t, []byte("b"), m)

	m, tag, err = pullSt.pull(w3)
	require.NoError(t, err)
	require.Equal(t, byte(TagRekey), tag)
	require.Equal(t, []byte("c"), m)

	m, tag, err = pullSt.pull(w4)
	require.NoError(t, err)
	require.Equal(t, byte(TagFinal), tag)
	require.Equal(t, []byte("d"), m)
}

type shortWriter struct{}

func (shortWriter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	return 1, nil
}

func TestWriterRejectsShortWrite(t *testing.T) {
	key := bytes.Repeat([]byte{0x55}, KeyBytes)
	w := NewWriter(shortWriter{}, key)
	_, err := w.Write([]byte("x"))
	require.Error(t, err)
	require.Contains(t, err.Error(), "short write")
}

func TestCrypterErrorWrapping(t *testing.T) {
	c := CrypterFromKeyPath("/no/such/key/file", KeyTransformHex)
	err := c.setup()
	require.Error(t, err)
	require.ErrorIs(t, err, os.ErrNotExist)
	require.Contains(t, err.Error(), "unable to read key")
}
