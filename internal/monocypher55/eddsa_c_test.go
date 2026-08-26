// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher55"
)

func TestEdDSA_ScalarbaseVsC(t *testing.T) {
	amalg, hdr := findMonocypherC(t)
	dir := t.TempDir()
	src := filepath.Join(dir, "t.c")
	bin := filepath.Join(dir, "t")
	os.WriteFile(src, []byte(`
#include "monocypher.h"
#include <stdio.h>
int main(){
  uint8_t sc[32]={2}; uint8_t pt[32];
  crypto_eddsa_scalarbase(pt, sc);
  for(int i=0;i<32;i++) printf("%02x", pt[i]);
  printf("\n");
  return 0;
}
`), 0644)
	if out, err := exec.Command("gcc", "-O0", "-I", hdr, "-o", bin, src, amalg).CombinedOutput(); err != nil {
		t.Fatalf("gcc %v %s", err, out)
	}
	cout, err := exec.Command(bin).Output()
	if err != nil {
		t.Fatal(err)
	}
	var sc [32]byte
	sc[0] = 2
	var pt [32]byte
	sgoi.Crypto_eddsa_scalarbase(pt[:], sc[:])
	ghex := fmt.Sprintf("%x", pt[:])
	got := string(bytes.TrimSpace(cout))
	if ghex != got {
		t.Fatalf("scalarbase mismatch\ngo %s\nc  %s", ghex, got)
	}
}
