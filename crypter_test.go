package secretstream

import (
	"bytes"
	"crypto/rand"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMockCrypterFromKey_ShouldReturnErrorOnEmptyKey(t *testing.T) {
	tests := map[string]struct {
		key string
	}{
		"empty": {key: ""},
		"short": {key: "short_key"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := CrypterFromKey(test.key, KeyTransformNone).setup()
			assert.Error(t, err, "no error on short key")
		})
	}
}

func TestMockCrypterFromKeyPath_ShouldReturnErrorOnNonExistentFile(t *testing.T) {
	assert.Error(t, CrypterFromKeyPath("", KeyTransformNone).setup(), "no error on non-existent key path")
}

func TestMockCrypterFromKeyPath_ShouldErrorIfTransformFails(t *testing.T) {
	type TestCase struct {
		key       string
		transform string
	}
	testcases := []TestCase{
		{key: "2e4af6d03c7f73f4a80b0594dee2b4bcd11300bafb8a", transform: KeyTransformHex},
		{key: "invalid hex", transform: KeyTransformHex},
		{key: "invalid base64", transform: KeyTransformBase64},
		{key: "DBXYo+QaYKCLSNad+m27jl2UHtW4Htm9pStJv1ujjKPB2N5fmitOFw==", transform: KeyTransformBase64},
	}
	for _, tc := range testcases {
		assert.Error(t, CrypterFromKey(tc.key, tc.transform).setup(), "no error on invalid encoding")
	}
}

func EncryptionCycle(t *testing.T, crypter *Crypter) {
	secret := strings.Repeat(" so very secret thing ", 1000)
	reader, writer := io.Pipe()
	encrypt, err := crypter.Encrypt(writer)
	assert.NoErrorf(t, err, "encryption error: %v", err)
	decrypt, err := crypter.Decrypt(reader)
	assert.NoErrorf(t, err, "decryption error: %v", err)
	go func() {
		_, _ = encrypt.Write([]byte(secret))
		_ = encrypt.Close()
	}()
	decrypted, err := io.ReadAll(decrypt)
	assert.NoErrorf(t, err, "decryption read error: %v", err)
	assert.Equal(t, secret, string(decrypted), "decrypted text not equals to open text")
	var buf [8]byte
	n, err := decrypt.Read(buf[:])
	assert.Equal(t, 0, n, "decryptor should not read any more data after ReadAll")
	assert.ErrorIs(t, err, io.EOF)
}

func TestEncryptionCycleFromKey(t *testing.T) {
	type TestCase struct {
		keyInline    string
		keyTransform string
	}
	testcases := []TestCase{
		{keyInline: "TEST_LIBSODIUM_KEY_______", keyTransform: KeyTransformNone},
		{keyInline: "4c0829fdfe7ae1987918edc585b1a90556d901eaea963c7625bb5734576dfb59", keyTransform: KeyTransformHex},
		{keyInline: "jv81yb3v3gNePrY0JmJ4q2j2NrqcM7tDYSHFoZ0tTIw=", keyTransform: KeyTransformBase64},
	}
	for _, tc := range testcases {
		EncryptionCycle(t, CrypterFromKey(tc.keyInline, tc.keyTransform))
	}
}

func TestEncryptionCycleFromKeyPath(t *testing.T) {
	type TestCase struct {
		keyPath      string
		keyTransform string
	}
	testcases := []TestCase{
		{keyPath: "./testdata/testKey", keyTransform: KeyTransformNone},
		{keyPath: "./testdata/testKeyHex", keyTransform: KeyTransformHex},
		{keyPath: "./testdata/testKeyB64", keyTransform: KeyTransformBase64},
	}
	for _, tc := range testcases {
		EncryptionCycle(t, CrypterFromKeyPath(tc.keyPath, tc.keyTransform))
	}
}

func TestSecretstreamRoundTripSizes(t *testing.T) {
	key := make([]byte, KeyBytes)
	_, err := rand.Read(key)
	require.NoError(t, err)
	sizes := []int{0, 1, 100, ChunkSize - 1, ChunkSize, ChunkSize + 1, ChunkSize*2 + 50, ChunkSize * 3}
	for _, sz := range sizes {
		t.Run(itoa(sz), func(t *testing.T) {
			plain := make([]byte, sz)
			_, _ = rand.Read(plain)
			var buf bytes.Buffer
			w := NewWriter(&buf, key)
			_, err := w.Write(plain)
			require.NoError(t, err)
			require.NoError(t, w.Close())
			out, err := io.ReadAll(NewReader(bytes.NewReader(buf.Bytes()), key))
			require.NoError(t, err)
			require.Equal(t, plain, out)
		})
	}
}

func TestSecretstreamWrongKey(t *testing.T) {
	key := bytes.Repeat([]byte{1}, KeyBytes)
	bad := bytes.Repeat([]byte{2}, KeyBytes)
	var buf bytes.Buffer
	w := NewWriter(&buf, key)
	_, _ = w.Write([]byte("hello secretstream"))
	_ = w.Close()
	_, err := io.ReadAll(NewReader(bytes.NewReader(buf.Bytes()), bad))
	require.Error(t, err)
}

func TestSecretstreamBitFlip(t *testing.T) {
	key := bytes.Repeat([]byte{7}, KeyBytes)
	var buf bytes.Buffer
	w := NewWriter(&buf, key)
	_, _ = w.Write(bytes.Repeat([]byte("x"), 100))
	_ = w.Close()
	raw := buf.Bytes()
	raw[len(raw)/2] ^= 0xff
	_, err := io.ReadAll(NewReader(bytes.NewReader(raw), key))
	require.Error(t, err)
}

func TestChunkSizeConstant(t *testing.T) {
	require.Equal(t, 8192, ChunkSize)
	require.Equal(t, 24, HeaderBytes)
	require.Equal(t, 17, ABytes)
}

func TestCrossWriteGoReadC(t *testing.T) {
	py := lookupCrossOracle(t)
	if py == "" {
		t.Skip("no libsodium cross oracle (install python3-nacl or set SECRETSTREAM_ORACLE)")
	}
	key := bytes.Repeat([]byte{0xab}, KeyBytes)
	plain := bytes.Repeat([]byte("cross-go-c-"), 900)
	var buf bytes.Buffer
	w := NewWriter(&buf, key)
	_, err := w.Write(plain)
	require.NoError(t, err)
	require.NoError(t, w.Close())
	dir := t.TempDir()
	cipherPath := filepath.Join(dir, "c.bin")
	keyPath := filepath.Join(dir, "key.bin")
	require.NoError(t, os.WriteFile(cipherPath, buf.Bytes(), 0o600))
	require.NoError(t, os.WriteFile(keyPath, key, 0o600))
	out, err := exec.Command(py, "-c", crossDecryptPy, cipherPath, keyPath).CombinedOutput()
	if err != nil {
		t.Fatalf("oracle failed: %v %s", err, out)
	}
	require.Equal(t, plain, out)
}

func lookupCrossOracle(t *testing.T) string {
	t.Helper()
	if p := os.Getenv("SECRETSTREAM_ORACLE"); p != "" {
		return p
	}
	if p := os.Getenv("WALG_LIBSODIUM_ORACLE"); p != "" {
		return p
	}
	if _, err := exec.LookPath("python3"); err != nil {
		return ""
	}
	if err := exec.Command("python3", "-c", "import nacl.bindings").Run(); err != nil {
		// try dedicated venv from dogfood
		venv := "/inference/venvs/walg-oracle/bin/python3"
		if err := exec.Command(venv, "-c", "import nacl.bindings").Run(); err == nil {
			return venv
		}
		return ""
	}
	return "python3"
}

const crossDecryptPy = `
import sys
from nacl.bindings.crypto_secretstream import (
    crypto_secretstream_xchacha20poly1305_state,
    crypto_secretstream_xchacha20poly1305_init_pull,
    crypto_secretstream_xchacha20poly1305_pull,
    crypto_secretstream_xchacha20poly1305_TAG_FINAL,
)
path, keypath = sys.argv[1], sys.argv[2]
key = open(keypath, "rb").read()
data = open(path, "rb").read()
hdr, rest = data[:24], data[24:]
state = crypto_secretstream_xchacha20poly1305_state()
crypto_secretstream_xchacha20poly1305_init_pull(state, hdr, key)
out = bytearray()
off = 0
CHUNK = 8192 + 17
while off < len(rest):
    remaining = len(rest) - off
    take = CHUNK if remaining > CHUNK else remaining
    piece = rest[off:off+take]
    off += take
    m, tag = crypto_secretstream_xchacha20poly1305_pull(state, bytes(piece), None)
    out.extend(m)
    if tag == crypto_secretstream_xchacha20poly1305_TAG_FINAL:
        break
sys.stdout.buffer.write(out)
`

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b [16]byte
	i := len(b)
	for n > 0 {
		i--
		b[i] = byte('0' + n%10)
		n /= 10
	}
	return string(b[i:])
}
