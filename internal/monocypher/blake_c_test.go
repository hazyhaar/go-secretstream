package monocypher_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher"
)

func TestBlake2b_VsC_Seed(t *testing.T) {
	c2 := "/devhoros/c2simd"
	amalg := filepath.Join(c2, "spec/c_sources/upstream/monocypher/4.0.2/monocypher.c")
	hdr := filepath.Join(c2, "spec/c_sources/upstream/monocypher/4.0.2")
	dir := t.TempDir()
	src := filepath.Join(dir, "b.c")
	bin := filepath.Join(dir, "b")
	os.WriteFile(src, []byte(`
#include "monocypher.h"
#include <stdio.h>
int main(){
  uint8_t a[64]; for(int i=0;i<32;i++) a[i]=3;
  crypto_blake2b(a, 64, a, 32);
  for(int i=0;i<64;i++) printf("%02x", a[i]);
  printf("\n");
  return 0;
}
`), 0644)
	exec.Command("gcc", "-O0", "-I", hdr, "-o", bin, src, amalg).Run()
	cout, _ := exec.Command(bin).Output()
	var a [64]byte
	for i := 0; i < 32; i++ {
		a[i] = 3
	}
	sgoi.Crypto_blake2b(a[:], 64, a[:], 32)
	ghex := fmt.Sprintf("%x", a[:])
	got := string(bytes.TrimSpace(cout))
	if ghex != got {
		t.Fatalf("blake\ngo %s\nc  %s", ghex, got)
	}
}
