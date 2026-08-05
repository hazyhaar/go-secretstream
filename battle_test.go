package secretstream

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBattleRandomSizes(t *testing.T) {
	key := make([]byte, KeyBytes)
	_, err := rand.Read(key)
	require.NoError(t, err)
	sizes := []int{
		0, 1, 2, 15, 16, 17, 31, 32, 63, 64, 65,
		1000, 4095, 4096, 4097,
		ChunkSize - 17, ChunkSize - 1, ChunkSize, ChunkSize + 1, ChunkSize + 17,
		ChunkSize*2 - 1, ChunkSize * 2, ChunkSize*2 + 1,
		ChunkSize*3 + 123, ChunkSize*5 + 7,
		100_000,
	}
	for _, sz := range sizes {
		sz := sz
		t.Run(fmt.Sprintf("n=%d", sz), func(t *testing.T) {
			t.Parallel()
			plain := make([]byte, sz)
			_, _ = rand.Read(plain)
			var buf bytes.Buffer
			w := NewWriter(&buf, key)
			_, err := io.Copy(w, bytes.NewReader(plain))
			require.NoError(t, err)
			require.NoError(t, w.Close())
			out, err := io.ReadAll(NewReader(bytes.NewReader(buf.Bytes()), key))
			require.NoError(t, err)
			require.Equal(t, plain, out)
		})
	}
}

func TestBattlePartialWrites(t *testing.T) {
	key := bytes.Repeat([]byte{0x3c}, KeyBytes)
	plain := make([]byte, ChunkSize*3+500)
	_, _ = rand.Read(plain)
	var buf bytes.Buffer
	w := NewWriter(&buf, key)
	for i := 0; i < len(plain); {
		n := 1 + int(plain[i]%17)
		if i+n > len(plain) {
			n = len(plain) - i
		}
		_, err := w.Write(plain[i : i+n])
		require.NoError(t, err)
		i += n
	}
	require.NoError(t, w.Close())
	out, err := io.ReadAll(NewReader(bytes.NewReader(buf.Bytes()), key))
	require.NoError(t, err)
	require.Equal(t, plain, out)
}

func TestBattleTruncateCiphertext(t *testing.T) {
	key := bytes.Repeat([]byte{9}, KeyBytes)
	var buf bytes.Buffer
	w := NewWriter(&buf, key)
	_, _ = w.Write(bytes.Repeat([]byte("t"), 5000))
	_ = w.Close()
	raw := buf.Bytes()
	for cut := 1; cut < len(raw) && cut < 40; cut++ {
		_, err := io.ReadAll(NewReader(bytes.NewReader(raw[:len(raw)-cut]), key))
		require.Error(t, err, "cut %d should fail", cut)
	}
}

func TestBattleCorruptHeader(t *testing.T) {
	key := bytes.Repeat([]byte{1}, KeyBytes)
	var buf bytes.Buffer
	w := NewWriter(&buf, key)
	_, _ = w.Write([]byte("hdr"))
	_ = w.Close()
	raw := buf.Bytes()
	raw[3] ^= 0xff
	_, err := io.ReadAll(NewReader(bytes.NewReader(raw), key))
	require.Error(t, err)
	require.NotContains(t, err.Error(), "EOF")
	msg := err.Error()
	require.True(t,
		strings.Contains(msg, "MAC mismatch") ||
			strings.Contains(msg, "decrypt chunk") ||
			strings.Contains(msg, "header") ||
			strings.Contains(msg, "secretstream"),
		"opaque error: %v", err)
}

func TestBattleErrorsAreLoud(t *testing.T) {
	key := bytes.Repeat([]byte{4}, KeyBytes)
	var buf bytes.Buffer
	w := NewWriter(&buf, key)
	_, _ = w.Write([]byte("noise-payload-for-loud-errors"))
	_ = w.Close()
	raw := append([]byte(nil), buf.Bytes()...)

	_, err := io.ReadAll(NewReader(bytes.NewReader(raw[:HeaderBytes+2]), key))
	require.Error(t, err)
	require.NotEqual(t, io.EOF, err)
	require.True(t,
		strings.Contains(err.Error(), "chunk") || strings.Contains(err.Error(), "MAC") ||
			strings.Contains(err.Error(), "short") || strings.Contains(err.Error(), "decrypt"),
		"error class missing: %q", err.Error())

	bad := bytes.Repeat([]byte{5}, KeyBytes)
	_, err = io.ReadAll(NewReader(bytes.NewReader(raw), bad))
	require.Error(t, err)
	require.Contains(t, err.Error(), "MAC mismatch")

	_, err = io.ReadAll(NewReader(bytes.NewReader(nil), key))
	require.Error(t, err)
	require.Contains(t, err.Error(), "header")
}

func TestBattleCorruptEachChunkByte(t *testing.T) {
	key := bytes.Repeat([]byte{2}, KeyBytes)
	var buf bytes.Buffer
	w := NewWriter(&buf, key)
	_, _ = w.Write(bytes.Repeat([]byte("m"), ChunkSize+100))
	_ = w.Close()
	raw := buf.Bytes()
	for _, off := range []int{HeaderBytes, HeaderBytes + 10, len(raw) - 5, len(raw) / 2} {
		if off < 0 || off >= len(raw) {
			continue
		}
		mut := append([]byte(nil), raw...)
		mut[off] ^= 0x5a
		_, err := io.ReadAll(NewReader(bytes.NewReader(mut), key))
		require.Error(t, err, "offset %d", off)
	}
}

func TestBattleParallelIndependentStreams(t *testing.T) {
	const n = 20
	var wg sync.WaitGroup
	errCh := make(chan error, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := make([]byte, KeyBytes)
			binary.LittleEndian.PutUint32(key, uint32(i))
			_, _ = rand.Read(key[4:])
			plain := bytes.Repeat([]byte{byte(i)}, 3000+i*100)
			var buf bytes.Buffer
			w := NewWriter(&buf, key)
			if _, err := w.Write(plain); err != nil {
				errCh <- err
				return
			}
			if err := w.Close(); err != nil {
				errCh <- err
				return
			}
			out, err := io.ReadAll(NewReader(bytes.NewReader(buf.Bytes()), key))
			if err != nil {
				errCh <- err
				return
			}
			if !bytes.Equal(out, plain) {
				errCh <- fmt.Errorf("mismatch i=%d", i)
			}
		}(i)
	}
	wg.Wait()
	close(errCh)
	for err := range errCh {
		require.NoError(t, err)
	}
}

func TestBattleLargeStream(t *testing.T) {
	if testing.Short() {
		t.Skip("large")
	}
	key := bytes.Repeat([]byte{0xee}, KeyBytes)
	const n = 2 * 1024 * 1024
	plain := make([]byte, n)
	_, _ = rand.Read(plain)
	pr, pw := io.Pipe()
	errCh := make(chan error, 1)
	go func() {
		w := NewWriter(pw, key)
		_, err := io.Copy(w, bytes.NewReader(plain))
		if err != nil {
			errCh <- err
			_ = pw.CloseWithError(err)
			return
		}
		errCh <- w.Close()
		_ = pw.Close()
	}()
	out, err := io.ReadAll(NewReader(pr, key))
	require.NoError(t, err)
	require.NoError(t, <-errCh)
	require.Equal(t, plain, out)
}

func TestCrossWriteCReadGo(t *testing.T) {
	py := lookupCrossOracle(t)
	if py == "" {
		t.Skip("no oracle")
	}
	key := bytes.Repeat([]byte{0xcd}, KeyBytes)
	plain := bytes.Repeat([]byte("c-to-go-"), 1200)
	dir := t.TempDir()
	plainPath := filepath.Join(dir, "p.bin")
	keyPath := filepath.Join(dir, "k.bin")
	cipherPath := filepath.Join(dir, "c.bin")
	require.NoError(t, os.WriteFile(plainPath, plain, 0o600))
	require.NoError(t, os.WriteFile(keyPath, key, 0o600))
	out, err := exec.Command(py, "-c", crossEncryptPy, plainPath, keyPath, cipherPath).CombinedOutput()
	if err != nil {
		t.Fatalf("oracle encrypt: %v %s", err, out)
	}
	cipher, err := os.ReadFile(cipherPath)
	require.NoError(t, err)
	got, err := io.ReadAll(NewReader(bytes.NewReader(cipher), key))
	require.NoError(t, err)
	require.Equal(t, plain, got)
}

const crossEncryptPy = `
import sys
from nacl.bindings.crypto_secretstream import (
    crypto_secretstream_xchacha20poly1305_state,
    crypto_secretstream_xchacha20poly1305_init_push,
    crypto_secretstream_xchacha20poly1305_push,
    crypto_secretstream_xchacha20poly1305_TAG_MESSAGE,
    crypto_secretstream_xchacha20poly1305_TAG_FINAL,
)
plain_path, key_path, out_path = sys.argv[1], sys.argv[2], sys.argv[3]
key = open(key_path, "rb").read()
plain = open(plain_path, "rb").read()
st = crypto_secretstream_xchacha20poly1305_state()
hdr = crypto_secretstream_xchacha20poly1305_init_push(st, key)
parts = [hdr]
CHUNK = 8192
off = 0
while off < len(plain):
    end = min(off + CHUNK, len(plain))
    piece = plain[off:end]
    off = end
    last = off >= len(plain)
    tag = crypto_secretstream_xchacha20poly1305_TAG_FINAL if last else crypto_secretstream_xchacha20poly1305_TAG_MESSAGE
    parts.append(crypto_secretstream_xchacha20poly1305_push(st, piece, None, tag))
if len(plain) == 0:
    parts.append(crypto_secretstream_xchacha20poly1305_push(st, b"", None, crypto_secretstream_xchacha20poly1305_TAG_FINAL))
open(out_path, "wb").write(b"".join(parts))
`
