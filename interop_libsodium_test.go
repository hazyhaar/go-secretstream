package secretstream55_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/hazyhaar/go-secretstream"
)

func packageDir(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("caller")
	}
	return filepath.Dir(file)
}

func driverPath(t *testing.T) string {
	t.Helper()
	p := filepath.Join(packageDir(t), "testdata", "libsodium_interop", "bin", "driver_secretstream")
	if _, err := os.Stat(p); err != nil {
		t.Skip("libsodium_interop_unavailable: driver missing — run make interop-driver")
	}
	return p
}

func goldenDir(t *testing.T) string {
	return filepath.Join(packageDir(t), "testdata", "libsodium_interop", "golden")
}

type goldenEntry struct {
	ID           string `json:"id"`
	Bytes        int    `json:"bytes"`
	SHA256Plain  string `json:"sha256_plain"`
	SHA256Wire   string `json:"sha256_wire"`
}

func loadManifest(t *testing.T) []goldenEntry {
	t.Helper()
	p := filepath.Join(goldenDir(t), "manifest.json")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Skip("libsodium_interop_unavailable: no goldens — run make interop-goldens")
	}
	var m []goldenEntry
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	return m
}

func sha256File(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

// TestInterop_Golden_CWire_GoDecrypt: C-produced wire → Go lsstream decrypt.
func TestInterop_Golden_CWire_GoDecrypt(t *testing.T) {
	for _, e := range loadManifest(t) {
		e := e
		t.Run(e.ID, func(t *testing.T) {
			g := goldenDir(t)
			key, err := os.ReadFile(filepath.Join(g, e.ID+".key"))
			if err != nil {
				t.Fatal(err)
			}
			wire, err := os.ReadFile(filepath.Join(g, e.ID+".wire"))
			if err != nil {
				t.Fatal(err)
			}
			plain, err := os.ReadFile(filepath.Join(g, e.ID+".plain"))
			if err != nil {
				t.Fatal(err)
			}
			if sha256File(t, filepath.Join(g, e.ID+".plain")) != e.SHA256Plain {
				t.Fatal("plain sha mismatch manifest")
			}
			r, err := secretstream55.NewLibsodiumDecryptor(bytes.NewReader(wire), key)
			if err != nil {
				t.Fatal(err)
			}
			got, err := io.ReadAll(r)
			if err != nil {
				t.Fatalf("go decrypt: %v", err)
			}
			if !bytes.Equal(got, plain) {
				t.Fatalf("plain mismatch n=%d got=%d", e.Bytes, len(got))
			}
		})
	}
}

// TestInterop_GoEncrypt_CDecrypt: Go encrypt → C pull.
func TestInterop_GoEncrypt_CDecrypt(t *testing.T) {
	drv := driverPath(t)
	dir := t.TempDir()
	key := bytes.Repeat([]byte{0x5a}, 32)
	sizes := []int{0, 1, 64, 8192, 8193, 20000}
	for _, sz := range sizes {
		sz := sz
		t.Run(itoa(sz), func(t *testing.T) {
			plain := make([]byte, sz)
			for i := range plain {
				plain[i] = byte((i * 3) % 251)
			}
			var wireBuf bytes.Buffer
			w, err := secretstream55.NewLibsodiumEncryptor(&wireBuf, key)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := w.Write(plain); err != nil {
				t.Fatal(err)
			}
			if err := w.Close(); err != nil {
				t.Fatal(err)
			}
			keyf := filepath.Join(dir, "k.bin")
			wiref := filepath.Join(dir, "w.bin")
			outf := filepath.Join(dir, "o.bin")
			if err := os.WriteFile(keyf, key, 0o600); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(wiref, wireBuf.Bytes(), 0o600); err != nil {
				t.Fatal(err)
			}
			cmd := exec.Command(drv, "pull-stream", keyf, wiref, outf)
			if out, err := cmd.CombinedOutput(); err != nil {
				t.Fatalf("c pull: %v\n%s", err, out)
			}
			got, err := os.ReadFile(outf)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(got, plain) {
				t.Fatalf("c decrypt mismatch n=%d", sz)
			}
		})
	}
}

// TestInterop_CEncrypt_GoDecrypt live via driver push.
func TestInterop_CEncrypt_GoDecrypt(t *testing.T) {
	drv := driverPath(t)
	dir := t.TempDir()
	key := bytes.Repeat([]byte{0x7e}, 32)
	plain := make([]byte, 12345)
	for i := range plain {
		plain[i] = byte(i)
	}
	keyf := filepath.Join(dir, "k.bin")
	plainf := filepath.Join(dir, "p.bin")
	wiref := filepath.Join(dir, "w.bin")
	_ = os.WriteFile(keyf, key, 0o600)
	_ = os.WriteFile(plainf, plain, 0o600)
	if out, err := exec.Command(drv, "push-stream", keyf, plainf, wiref).CombinedOutput(); err != nil {
		t.Fatalf("c push: %v\n%s", err, out)
	}
	wire, _ := os.ReadFile(wiref)
	r, err := secretstream55.NewLibsodiumDecryptor(bytes.NewReader(wire), key)
	if err != nil {
		t.Fatal(err)
	}
	got, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, plain) {
		t.Fatal("mismatch")
	}
}

// TestInterop_MACTamper_FailClosed flips one MAC byte.
func TestInterop_MACTamper_FailClosed(t *testing.T) {
	key := bytes.Repeat([]byte{1}, 32)
	var buf bytes.Buffer
	w, _ := secretstream55.NewLibsodiumEncryptor(&buf, key)
	_, _ = w.Write([]byte("tamper-me"))
	_ = w.Close()
	wire := buf.Bytes()
	if len(wire) < 24+17 {
		t.Fatal("wire too short")
	}
	// flip last byte (inside MAC of final chunk)
	wire[len(wire)-1] ^= 0xff
	r, err := secretstream55.NewLibsodiumDecryptor(bytes.NewReader(wire), key)
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.ReadAll(r)
	if err == nil {
		t.Fatal("expected MAC failure")
	}
	// C side
	drv := driverPath(t)
	dir := t.TempDir()
	keyf := filepath.Join(dir, "k.bin")
	wiref := filepath.Join(dir, "w.bin")
	outf := filepath.Join(dir, "o.bin")
	_ = os.WriteFile(keyf, key, 0o600)
	_ = os.WriteFile(wiref, wire, 0o600)
	if err := exec.Command(drv, "pull-stream", keyf, wiref, outf).Run(); err == nil {
		t.Fatal("expected C pull failure")
	}
}

func TestInterop_GoEncrypt_GoDecrypt_Large(t *testing.T) {
	key := make([]byte, 32)
	for i := range key {
		key[i] = byte(i + 9)
	}
	plain := make([]byte, 200000)
	for i := range plain {
		plain[i] = byte(i % 251)
	}
	var buf bytes.Buffer
	w, err := secretstream55.NewLibsodiumEncryptor(&buf, key)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := io.Copy(w, bytes.NewReader(plain)); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	r, err := secretstream55.NewLibsodiumDecryptor(bytes.NewReader(buf.Bytes()), key)
	if err != nil {
		t.Fatal(err)
	}
	got, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, plain) {
		t.Fatal("large mismatch")
	}
}
