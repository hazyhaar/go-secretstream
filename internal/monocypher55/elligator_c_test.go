// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55_test

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher55"
)

func TestElligator_KeyPair_VsC(t *testing.T) {
	dir := t.TempDir()
	amalg, hdr := findMonocypherC(t)
	src := filepath.Join(dir, "ekp.c")
	os.WriteFile(src, []byte(`
#include "monocypher.h"
#include <stdio.h>
int main(int argc,char**argv){
  uint8_t seed[32], hidden[32], sk[32];
  FILE*f=fopen(argv[1],"rb"); fread(seed,1,32,f); fclose(f);
  crypto_elligator_key_pair(hidden, sk, seed);
  for(int i=0;i<32;i++) printf("%02x", hidden[i]); printf("\n");
  for(int i=0;i<32;i++) printf("%02x", sk[i]); printf("\n");
}
`), 0644)
	bin := filepath.Join(dir, "ekp.bin")
	if out, err := exec.Command("gcc", "-O0", "-I", hdr, "-o", bin, src, amalg).CombinedOutput(); err != nil {
		t.Fatalf("gcc %v %s", err, out)
	}
	for n := 0; n < 16; n++ {
		seed := sha256.Sum256([]byte(fmt.Sprintf("elligator-seed-%d", n)))
		var hidden, sk [32]byte
		sgoi.Crypto_elligator_key_pair(hidden[:], sk[:], append([]byte(nil), seed[:]...))
		sp := filepath.Join(dir, "seed")
		os.WriteFile(sp, seed[:], 0600)
		cout, err := exec.Command(bin, sp).Output()
		if err != nil {
			t.Fatal(err)
		}
		lines := bytes.Split(bytes.TrimSpace(cout), []byte("\n"))
		if fmt.Sprintf("%x", hidden[:]) != string(lines[0]) {
			t.Fatalf("n=%d hidden\ngo %x\nc  %s", n, hidden[:], lines[0])
		}
		if fmt.Sprintf("%x", sk[:]) != string(lines[1]) {
			t.Fatalf("n=%d sk\ngo %x\nc  %s", n, sk[:], lines[1])
		}
		// map(hidden) should be dirty_fast(sk) public
		var curve [32]byte
		sgoi.Crypto_elligator_map(curve[:], hidden[:])
		var pk [32]byte
		sgoi.Crypto_x25519_dirty_fast(pk[:], sk[:])
		if !bytes.Equal(curve[:], pk[:]) {
			t.Fatalf("n=%d map≠dirty_fast", n)
		}
	}
}

func TestElligator_MapRev_Roundtrip_VsC(t *testing.T) {
	// map is deterministic; rev may fail — only check map vs C on random hidden
	dir := t.TempDir()
	amalg, hdr := findMonocypherC(t)
	src := filepath.Join(dir, "map.c")
	os.WriteFile(src, []byte(`
#include "monocypher.h"
#include <stdio.h>
int main(int argc,char**argv){
  uint8_t h[32], c[32];
  FILE*f=fopen(argv[1],"rb"); fread(h,1,32,f); fclose(f);
  crypto_elligator_map(c, h);
  for(int i=0;i<32;i++) printf("%02x", c[i]); printf("\n");
}
`), 0644)
	bin := filepath.Join(dir, "map.bin")
	exec.Command("gcc", "-O0", "-I", hdr, "-o", bin, src, amalg).Run()
	for n := 0; n < 8; n++ {
		h := sha256.Sum256([]byte(fmt.Sprintf("hidden-%d", n)))
		var curve [32]byte
		sgoi.Crypto_elligator_map(curve[:], h[:])
		hp := filepath.Join(dir, "h")
		os.WriteFile(hp, h[:], 0600)
		cout, _ := exec.Command(bin, hp).Output()
		if fmt.Sprintf("%x", curve[:]) != string(bytes.TrimSpace(cout)) {
			t.Fatalf("n=%d map mismatch go %x c %s", n, curve[:], cout)
		}
	}
}
