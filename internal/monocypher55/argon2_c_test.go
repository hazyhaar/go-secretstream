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

func TestArgon2_VsC_ThreeConfigs(t *testing.T) {
	type cfg struct {
		name                  string
		algo                  uint32
		blocks, passes, lanes uint32
		hashSize              uint32
		pass, salt, key, ad   []byte
	}
	cases := []cfg{
		{name: "i_m8_p1", algo: sgoi.Crypto_argon2_i, blocks: 8, passes: 1, lanes: 1, hashSize: 32,
			pass: []byte("password"), salt: []byte("saltsaltsaltsalt")},
		{name: "d_m8_p2", algo: sgoi.Crypto_argon2_d, blocks: 8, passes: 2, lanes: 1, hashSize: 32,
			pass: []byte("pw"), salt: bytes.Repeat([]byte{1}, 16)},
		{name: "id_ad_key", algo: sgoi.Crypto_argon2_id, blocks: 16, passes: 1, lanes: 1, hashSize: 32,
			pass: []byte("secret"), salt: bytes.Repeat([]byte{2}, 16),
			key: []byte("keykeykeykeykeykeykeykeykeykeyke"), ad: []byte("ad-data")},
	}
	dir := t.TempDir()
	amalg, hdr := findMonocypherC(t)
	src := filepath.Join(dir, "argon.c")
	os.WriteFile(src, []byte(`
#include "monocypher.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int main(int argc, char**argv){
  uint32_t algo=atoi(argv[1]), blocks=atoi(argv[2]), passes=atoi(argv[3]), lanes=atoi(argv[4]), hs=atoi(argv[5]);
  FILE*f; size_t psz,ssz,ksz=0,asz=0;
  uint8_t pass[256],salt[64],key[64],ad[64],hash[64];
  f=fopen(argv[6],"rb"); psz=fread(pass,1,256,f); fclose(f);
  f=fopen(argv[7],"rb"); ssz=fread(salt,1,64,f); fclose(f);
  if(argc>8 && argv[8][0]){ f=fopen(argv[8],"rb"); ksz=fread(key,1,64,f); fclose(f); }
  if(argc>9 && argv[9][0]){ f=fopen(argv[9],"rb"); asz=fread(ad,1,64,f); fclose(f); }
  crypto_argon2_config c={.algorithm=algo,.nb_blocks=blocks,.nb_passes=passes,.nb_lanes=lanes};
  crypto_argon2_inputs in={.pass=pass,.salt=salt,.pass_size=(uint32_t)psz,.salt_size=(uint32_t)ssz};
  crypto_argon2_extras ex=crypto_argon2_no_extras;
  if(ksz){ ex.key=key; ex.key_size=(uint32_t)ksz; }
  if(asz){ ex.ad=ad; ex.ad_size=(uint32_t)asz; }
  void*wa=malloc(blocks*1024);
  crypto_argon2(hash, hs, wa, c, in, ex);
  for(uint32_t i=0;i<hs;i++) printf("%02x", hash[i]);
  printf("\n"); free(wa);
}
`), 0644)
	bin := filepath.Join(dir, "argon.bin")
	if out, err := exec.Command("gcc", "-O0", "-I", hdr, "-o", bin, src, amalg).CombinedOutput(); err != nil {
		t.Fatalf("gcc: %v %s", err, out)
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			work := make([]byte, int(tc.blocks)*1024)
			hash := make([]byte, tc.hashSize)
			cfg := sgoi.Crypto_argon2_config{Algorithm: tc.algo, Nb_blocks: tc.blocks, Nb_passes: tc.passes, Nb_lanes: tc.lanes}
			in := sgoi.Crypto_argon2_inputs{Pass: tc.pass, Salt: tc.salt, Pass_size: uint32(len(tc.pass)), Salt_size: uint32(len(tc.salt))}
			ex := sgoi.Crypto_argon2_no_extras
			if len(tc.key) > 0 {
				ex.Key, ex.Key_size = tc.key, uint32(len(tc.key))
			}
			if len(tc.ad) > 0 {
				ex.Ad, ex.Ad_size = tc.ad, uint32(len(tc.ad))
			}
			sgoi.Crypto_argon2(hash, tc.hashSize, work, cfg, in, ex)

			passP := filepath.Join(dir, "pass")
			saltP := filepath.Join(dir, "salt")
			os.WriteFile(passP, tc.pass, 0600)
			os.WriteFile(saltP, tc.salt, 0600)
			args := []string{fmt.Sprint(tc.algo), fmt.Sprint(tc.blocks), fmt.Sprint(tc.passes), fmt.Sprint(tc.lanes), fmt.Sprint(tc.hashSize), passP, saltP}
			keyP, adP := "", ""
			if len(tc.key) > 0 {
				keyP = filepath.Join(dir, "key")
				os.WriteFile(keyP, tc.key, 0600)
			}
			if len(tc.ad) > 0 {
				adP = filepath.Join(dir, "ad")
				os.WriteFile(adP, tc.ad, 0600)
			}
			args = append(args, keyP, adP)
			cout, err := exec.Command(bin, args...).Output()
			if err != nil {
				t.Fatal(err)
			}
			want := string(bytes.TrimSpace(cout))
			got := fmt.Sprintf("%x", hash)
			if got != want {
				t.Fatalf("mismatch\ngo %s\nc  %s", got, want)
			}
		})
	}
}
