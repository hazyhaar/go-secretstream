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

func cBuild(t *testing.T, dir, name, src string) string {
	t.Helper()
	amalg, hdr := findMonocypherC(t)
	p := filepath.Join(dir, name+".c")
	bin := filepath.Join(dir, name+".bin")
	os.WriteFile(p, []byte(src), 0644)
	if out, err := exec.Command("gcc", "-O0", "-I", hdr, "-o", bin, p, amalg).CombinedOutput(); err != nil {
		t.Fatalf("gcc %s: %v %s", name, err, out)
	}
	return bin
}

func TestRound3_ScalarbaseRandom(t *testing.T) {
	dir := t.TempDir()
	bin := cBuild(t, dir, "scbase", `
#include "monocypher.h"
#include <stdio.h>
int main(int argc, char**argv){
  uint8_t sc[32], pt[32];
  FILE*f=fopen(argv[1],"rb"); fread(sc,1,32,f); fclose(f);
  crypto_eddsa_scalarbase(pt, sc);
  fwrite(pt,1,32,stdout);
}`)
	for n := 0; n < 32; n++ {
		h := sha256.Sum256([]byte(fmt.Sprintf("sc-%d", n)))
		sc := append([]byte(nil), h[:]...)
		sc[31] &= 0x0f
		var pt [32]byte
		sgoi.Crypto_eddsa_scalarbase(pt[:], sc)
		scPath := filepath.Join(dir, "scalar.bin")
		os.WriteFile(scPath, sc, 0600)
		cout, err := exec.Command(bin, scPath).Output()
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(pt[:], cout) {
			t.Fatalf("n=%d\ngo %x\nc  %x", n, pt[:], cout)
		}
	}
}

func TestRound3_KeyPairSignMany(t *testing.T) {
	dir := t.TempDir()
	binKP := cBuild(t, dir, "kp", `
#include "monocypher.h"
#include <stdio.h>
int main(int argc,char**argv){
  uint8_t seed[32],sk[64],pk[32];
  FILE*f=fopen(argv[1],"rb"); fread(seed,1,32,f); fclose(f);
  crypto_eddsa_key_pair(sk,pk,seed);
  fwrite(pk,1,32,stdout); fwrite(sk,1,64,stdout);
}`)
	binVer := cBuild(t, dir, "ver", `
#include "monocypher.h"
#include <stdio.h>
int main(int argc,char**argv){
  uint8_t pk[32],sig[64],msg[256];
  FILE*f; size_t n;
  f=fopen(argv[1],"rb"); fread(pk,1,32,f); fclose(f);
  f=fopen(argv[2],"rb"); fread(sig,1,64,f); fclose(f);
  f=fopen(argv[3],"rb"); n=fread(msg,1,256,f); fclose(f);
  printf("%d\n", crypto_eddsa_check(sig,pk,msg,n));
}`)
	binSign := cBuild(t, dir, "sign", `
#include "monocypher.h"
#include <stdio.h>
int main(int argc,char**argv){
  uint8_t sk[64],sig[64],msg[256];
  FILE*f; size_t n;
  f=fopen(argv[1],"rb"); fread(sk,1,64,f); fclose(f);
  f=fopen(argv[2],"rb"); n=fread(msg,1,256,f); fclose(f);
  crypto_eddsa_sign(sig,sk,msg,n);
  fwrite(sig,1,64,stdout);
}`)
	for n := 0; n < 16; n++ {
		seed := sha256.Sum256([]byte(fmt.Sprintf("seed-%d", n)))
		var sk [64]byte
		var pk [32]byte
		sgoi.Crypto_eddsa_key_pair(sk[:], pk[:], append([]byte(nil), seed[:]...))
		os.WriteFile(filepath.Join(dir, "seed"), seed[:], 0600)
		cout, _ := exec.Command(binKP, filepath.Join(dir, "seed")).Output()
		if !bytes.Equal(pk[:], cout[:32]) {
			t.Fatalf("pk n=%d go %x c %x", n, pk[:], cout[:32])
		}
		if !bytes.Equal(sk[:], cout[32:96]) {
			t.Fatalf("sk n=%d", n)
		}
		msg := []byte(fmt.Sprintf("msg-%d-night", n))
		var sig [64]byte
		sgoi.Crypto_eddsa_sign(sig[:], sk[:], msg, uint64(len(msg)))
		if sgoi.Crypto_eddsa_check(sig[:], pk[:], msg, uint64(len(msg))) != 0 {
			t.Fatalf("go verify n=%d", n)
		}
		os.WriteFile(filepath.Join(dir, "pk"), pk[:], 0600)
		os.WriteFile(filepath.Join(dir, "sig"), sig[:], 0600)
		os.WriteFile(filepath.Join(dir, "msg"), msg, 0600)
		vout, _ := exec.Command(binVer, filepath.Join(dir, "pk"), filepath.Join(dir, "sig"), filepath.Join(dir, "msg")).Output()
		if string(bytes.TrimSpace(vout)) != "0" {
			t.Fatalf("C verify go-sig n=%d → %s", n, vout)
		}
		os.WriteFile(filepath.Join(dir, "sk"), sk[:], 0600)
		csig, _ := exec.Command(binSign, filepath.Join(dir, "sk"), filepath.Join(dir, "msg")).Output()
		if sgoi.Crypto_eddsa_check(csig, pk[:], msg, uint64(len(msg))) != 0 {
			t.Fatalf("go verify c-sig n=%d", n)
		}
		if !bytes.Equal(sig[:], csig) {
			t.Fatalf("sig bytes differ n=%d", n)
		}
	}
}
