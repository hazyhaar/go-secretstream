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

func cScalarbase(t *testing.T, sc0 byte) string {
	t.Helper()
	amalg, hdr := findMonocypherC(t)
	dir := t.TempDir()
	src := filepath.Join(dir, "t.c")
	bin := filepath.Join(dir, "t")
	os.WriteFile(src, []byte(fmt.Sprintf(`
#include "monocypher.h"
#include <stdio.h>
int main(){
  uint8_t sc[32]={0}; sc[0]=%d; uint8_t pt[32];
  crypto_eddsa_scalarbase(pt, sc);
  for(int i=0;i<32;i++) printf("%%02x", pt[i]);
  printf("\n");
}
`, sc0)), 0644)
	if out, err := exec.Command("gcc", "-O0", "-I", hdr, "-o", bin, src, amalg).CombinedOutput(); err != nil {
		t.Fatalf("gcc %v %s", err, out)
	}
	cout, err := exec.Command(bin).Output()
	if err != nil {
		t.Fatal(err)
	}
	return string(bytes.TrimSpace(cout))
}

func TestRound_ScalarbaseMatrix(t *testing.T) {
	for _, sc0 := range []byte{0, 1, 2, 3, 7, 255} {
		sc0 := sc0
		t.Run(fmt.Sprintf("sc0=%d", sc0), func(t *testing.T) {
			var sc [32]byte
			sc[0] = sc0
			var pt [32]byte
			sgoi.Crypto_eddsa_scalarbase(pt[:], sc[:])
			want := cScalarbase(t, sc0)
			got := fmt.Sprintf("%x", pt[:])
			if got != want {
				t.Fatalf("mismatch\ngo %s\nc  %s", got, want)
			}
		})
	}
}

func TestRound_KeyPairSignVerify_C(t *testing.T) {
	seed := bytes.Repeat([]byte{9}, 32)
	var sk [64]byte
	var pk [32]byte
	sgoi.Crypto_eddsa_key_pair(sk[:], pk[:], append([]byte(nil), seed...))
	// C key_pair
	amalg, hdr := findMonocypherC(t)
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "seed"), seed, 0600)
	src := filepath.Join(dir, "k.c")
	bin := filepath.Join(dir, "k")
	os.WriteFile(src, []byte(fmt.Sprintf(`
#include "monocypher.h"
#include <stdio.h>
int main(){
  uint8_t seed[32], sk[64], pk[32];
  FILE*f=fopen("%s/seed","rb"); fread(seed,1,32,f); fclose(f);
  crypto_eddsa_key_pair(sk, pk, seed);
  for(int i=0;i<32;i++) printf("%%02x", pk[i]); printf("\n");
  for(int i=0;i<64;i++) printf("%%02x", sk[i]); printf("\n");
}
`, dir)), 0644)
	exec.Command("gcc", "-O0", "-I", hdr, "-o", bin, src, amalg).Run()
	cout, _ := exec.Command(bin).Output()
	lines := bytes.Split(bytes.TrimSpace(cout), []byte("\n"))
	if fmt.Sprintf("%x", pk[:]) != string(lines[0]) {
		t.Fatalf("pk\ngo %x\nc  %s", pk[:], lines[0])
	}
	if fmt.Sprintf("%x", sk[:]) != string(lines[1]) {
		t.Fatalf("sk\ngo %x\nc  %s", sk[:], lines[1])
	}
	msg := []byte("night-round")
	var sig [64]byte
	sgoi.Crypto_eddsa_sign(sig[:], sk[:], msg, uint64(len(msg)))
	if sgoi.Crypto_eddsa_check(sig[:], pk[:], msg, uint64(len(msg))) != 0 {
		t.Fatal("go verify failed")
	}
	// C verify
	os.WriteFile(filepath.Join(dir, "pk"), pk[:], 0600)
	os.WriteFile(filepath.Join(dir, "sig"), sig[:], 0600)
	os.WriteFile(filepath.Join(dir, "msg"), msg, 0600)
	os.WriteFile(filepath.Join(dir, "v.c"), []byte(fmt.Sprintf(`
#include "monocypher.h"
#include <stdio.h>
int main(){
  uint8_t pk[32],sig[64],msg[32];
  FILE*f;
  f=fopen("%s/pk","rb"); fread(pk,1,32,f); fclose(f);
  f=fopen("%s/sig","rb"); fread(sig,1,64,f); fclose(f);
  f=fopen("%s/msg","rb"); fread(msg,1,%d,f); fclose(f);
  printf("%%d\n", crypto_eddsa_check(sig,pk,msg,%d));
}
`, dir, dir, dir, len(msg), len(msg))), 0644)
	exec.Command("gcc", "-O0", "-I", hdr, "-o", filepath.Join(dir, "v"), filepath.Join(dir, "v.c"), amalg).Run()
	vout, _ := exec.Command(filepath.Join(dir, "v")).Output()
	if string(bytes.TrimSpace(vout)) != "0" {
		t.Fatalf("C verify=%s", vout)
	}
}
