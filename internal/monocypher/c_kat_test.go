package monocypher_test

import (
	"bytes"
	"encoding/hex"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher"
)

func findC2simdRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("caller")
	}
	// .../pkg/monocypher55 -> /devhoros
	dev := filepath.Join(filepath.Dir(file), "..", "..")
	c2 := filepath.Join(dev, "c2simd")
	if _, err := os.Stat(filepath.Join(c2, "sgoiter")); err == nil {
		return c2
	}
	t.Skip("c2simd root not found from ", file)
	return ""
}

func parseDriverOut(out []byte) map[string]map[string]string {
	res := map[string]map[string]string{}
	var cur string
	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "NAME=") {
			cur = strings.TrimPrefix(line, "NAME=")
			res[cur] = map[string]string{}
			continue
		}
		if cur == "" {
			continue
		}
		if i := strings.IndexByte(line, '='); i > 0 {
			res[cur][line[:i]] = line[i+1:]
		}
	}
	return res
}

func TestAEAD_VsMonocypherCGCC(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc missing")
	}
	c2 := findC2simdRoot(t)
	amalg := filepath.Join(c2, "spec", "c_sources", "upstream", "monocypher", "4.0.2", "monocypher.c")
	hdr := filepath.Join(c2, "spec", "c_sources", "upstream", "monocypher", "4.0.2")
	driver := filepath.Join(c2, "sgoiter", "testdata", "c_kat", "aead_kat_driver.c")
	if _, err := os.Stat(amalg); err != nil {
		t.Skip("monocypher.c missing")
	}
	tmp := t.TempDir()
	bin := filepath.Join(tmp, "aead_kat_driver")
	cmd := exec.Command("gcc", "-O0", "-I", hdr, "-o", bin, driver, amalg)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("gcc: %v\n%s", err, out)
	}
	out, err := exec.Command(bin, "all").CombinedOutput()
	if err != nil {
		t.Fatalf("driver: %v\n%s", err, out)
	}
	gotC := parseDriverOut(out)
	key, nonce := keyNonce()

	type tc struct {
		name string
		ad   []byte
		pt   []byte
	}
	cases := []tc{
		{"pt0_ad_empty", nil, nil},
		{"pt1_ad_empty", nil, []byte{0x41}},
		{"pt36_header", []byte("HEADER"), []byte("HELLO MONOCYPHER SGOITER AEAD CGO=0!")},
		{"pt64_ad_empty", nil, mkPT(64, "i%251")},
		{"pt65_ad_empty", nil, mkPT(65, "i%251")},
		{"pt129_ad_empty", nil, mkPT(129, "i%251")},
		{"pt193_ad_empty", nil, mkPT(193, "i%251")},
		{"pt1024_header1kb", []byte("HEADER 1KB"), mkPT(1024, "(i*17+3)%251")},
		{"pt4096_ad_empty", nil, mkPT(4096, "i%251")},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			block, ok := gotC[c.name]
			if !ok {
				t.Fatalf("C missing %s", c.name)
			}
			ctS, macS, err := sgoi.AEADLock(key, nonce, c.ad, c.pt)
			if err != nil {
				t.Fatal(err)
			}
			ctC, _ := hex.DecodeString(block["CT"])
			macC, _ := hex.DecodeString(block["MAC"])
			if !bytes.Equal(macS, macC) {
				t.Fatalf("mac\nsgoi %x\nC    %s", macS, block["MAC"])
			}
			if !bytes.Equal(ctS, ctC) {
				t.Fatalf("ct mismatch n=%d", len(c.pt))
			}
		})
	}
}

func TestChacha20IETF_VsC(t *testing.T) {
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc missing")
	}
	c2 := findC2simdRoot(t)
	amalg := filepath.Join(c2, "spec", "c_sources", "upstream", "monocypher", "4.0.2", "monocypher.c")
	hdr := filepath.Join(c2, "spec", "c_sources", "upstream", "monocypher", "4.0.2")
	driver := filepath.Join(c2, "sgoiter", "testdata", "c_kat", "aead_kat_driver.c")
	tmp := t.TempDir()
	bin := filepath.Join(tmp, "aead_kat_driver")
	cmd := exec.Command("gcc", "-O0", "-I", hdr, "-o", bin, driver, amalg)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("gcc: %v\n%s", err, out)
	}
	out, err := exec.Command(bin, "chacha20_ietf").CombinedOutput()
	if err != nil {
		t.Fatalf("driver: %v\n%s", err, out)
	}
	gotC := parseDriverOut(out)
	block, ok := gotC["chacha20_ietf"]
	if !ok {
		t.Fatalf("C missing chacha20_ietf")
	}
	key := make([]byte, 32)
	for i := 0; i < 32; i++ { key[i] = byte(i + 1) }
	nonce := make([]byte, 12)
	for i := 0; i < 12; i++ { nonce[i] = byte(i + 5) }
	pt := make([]byte, 64)
	for i := 0; i < 64; i++ { pt[i] = byte(i % 251) }
	ctS := make([]byte, 64)
	ctrS := sgoi.Crypto_chacha20_ietf(ctS, pt, 64, key, nonce, 0x1000)

	ctC, _ := hex.DecodeString(block["CT"])
	if !bytes.Equal(ctS, ctC) {
		t.Fatalf("ietf ct mismatch:\nsgoi %x\nC    %s", ctS, block["CT"])
	}
	if sprintfHex := strings.ToLower(hex.EncodeToString([]byte{byte(ctrS >> 24), byte(ctrS >> 16), byte(ctrS >> 8), byte(ctrS)})); sprintfHex != strings.ToLower(block["CTR"]) {
		t.Fatalf("ietf ctr mismatch:\nsgoi %08x\nC    %s", ctrS, block["CTR"])
	}
}

func TestFeIsOdd_WipeCheck(t *testing.T) {
	// f is an odd field element (1)
	f := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	got := sgoi.Fe_isodd(f)
	if got != 1 {
		t.Fatalf("Fe_isodd(1) = %d, expected 1 (C1 wipe-before-return bug)", got)
	}
}
