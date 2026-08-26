// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher55"
)

const streamDriverC = `
#include "monocypher.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>

static const size_t sizes[] = {0, 1, 63, 64, 65, 4095, 4096, 4097, 65536};
static const int nsizes = 9;

static void hexout(const uint8_t *p, size_t n) {
	for (size_t i = 0; i < n; i++) printf("%02x", p[i]);
}

static int hexval(char c) {
	if (c >= '0' && c <= '9') return c - '0';
	if (c >= 'a' && c <= 'f') return c - 'a' + 10;
	if (c >= 'A' && c <= 'F') return c - 'A' + 10;
	return -1;
}

static size_t hexin(uint8_t *dst, size_t cap, const char *s) {
	size_t n = 0;
	while (s[0] && s[1] && n < cap) {
		int a = hexval(s[0]), b = hexval(s[1]);
		if (a < 0 || b < 0) break;
		dst[n++] = (uint8_t)((a << 4) | b);
		s += 2;
	}
	return n;
}

static void stripnl(char *s) {
	size_t n = strlen(s);
	while (n > 0 && (s[n - 1] == '\n' || s[n - 1] == '\r')) {
		s[--n] = 0;
	}
}

static void fill_pt(uint8_t *p, size_t n) {
	for (size_t i = 0; i < n; i++) p[i] = (uint8_t)(i % 251);
}

static void fill_ad(uint8_t *p, size_t n) {
	for (size_t i = 0; i < n; i++) p[i] = (uint8_t)(0xA0 + i);
}

static void init_ctx(crypto_aead_ctx *ctx, const char *kind,
                     const uint8_t key[32], const uint8_t *nonce) {
	if (strcmp(kind, "x") == 0) crypto_aead_init_x(ctx, key, nonce);
	else if (strcmp(kind, "djb") == 0) crypto_aead_init_djb(ctx, key, nonce);
	else if (strcmp(kind, "ietf") == 0) crypto_aead_init_ietf(ctx, key, nonce);
	else {
		fprintf(stderr, "bad init\n");
		exit(2);
	}
}

int main(int argc, char **argv) {
	if (argc < 4) {
		fprintf(stderr, "usage: %s write|read x|djb|ietf adsize\n", argv[0]);
		return 2;
	}
	const char *mode = argv[1];
	const char *kind = argv[2];
	int adsize = atoi(argv[3]);
	if (adsize < 0 || adsize > 64) return 2;

	uint8_t key[32];
	for (int i = 0; i < 32; i++) key[i] = (uint8_t)(i + 1);
	uint8_t nonce[24];
	for (int i = 0; i < 24; i++) nonce[i] = (uint8_t)(i + 10);

	crypto_aead_ctx ctx;
	init_ctx(&ctx, kind, key, nonce);

	uint8_t ad[64];
	fill_ad(ad, (size_t)adsize);
	uint8_t *pt = malloc(65536);
	uint8_t *ct = malloc(65536);
	uint8_t mac[16];
	if (!pt || !ct) return 3;

	if (strcmp(mode, "write") == 0) {
		for (int i = 0; i < nsizes; i++) {
			size_t n = sizes[i];
			fill_pt(pt, n);
			crypto_aead_write(&ctx, n ? ct : NULL, mac,
			                  adsize ? ad : NULL, (size_t)adsize,
			                  n ? pt : NULL, n);
			printf("MSG=%d AD=%d N=%zu\n", i, adsize, n);
			printf("CT=");
			hexout(ct, n);
			printf("\n");
			printf("MAC=");
			hexout(mac, 16);
			printf("\n");
		}
		free(pt);
		free(ct);
		return 0;
	}

	if (strcmp(mode, "read") == 0) {
		static char linebuf[200000];
		for (int i = 0; i < nsizes; i++) {
			size_t n = sizes[i];
			if (!fgets(linebuf, sizeof linebuf, stdin)) {
				fprintf(stderr, "eof ct %d\n", i);
				return 4;
			}
			stripnl(linebuf);
			if (hexin(ct, 65536, linebuf) != n) {
				fprintf(stderr, "ct len want %zu\n", n);
				return 5;
			}
			if (!fgets(linebuf, sizeof linebuf, stdin)) {
				fprintf(stderr, "eof mac %d\n", i);
				return 4;
			}
			stripnl(linebuf);
			if (hexin(mac, 16, linebuf) != 16) {
				fprintf(stderr, "mac len\n");
				return 5;
			}
			int st = crypto_aead_read(&ctx, n ? pt : NULL, mac,
			                          adsize ? ad : NULL, (size_t)adsize,
			                          n ? ct : NULL, n);
			printf("MSG=%d STATUS=%d N=%zu\n", i, st, n);
			printf("PT=");
			if (st == 0) hexout(pt, n);
			printf("\n");
		}
		free(pt);
		free(ct);
		return 0;
	}

	fprintf(stderr, "bad mode\n");
	return 2;
}
`

func streamSizes() []int {
	return []int{0, 1, 63, 64, 65, 4095, 4096, 4097, 65536}
}

func streamKey() []byte {
	key := make([]byte, 32)
	for i := range key {
		key[i] = byte(i + 1)
	}
	return key
}

func streamNonce24() []byte {
	nonce := make([]byte, 24)
	for i := range nonce {
		nonce[i] = byte(i + 10)
	}
	return nonce
}

func streamAD(n int) []byte {
	ad := make([]byte, n)
	for i := range ad {
		ad[i] = byte(0xA0 + i)
	}
	return ad
}

func compileStreamDriver(t *testing.T) string {
	t.Helper()
	amalg, hdr := findMonocypherC(t)
	dir := t.TempDir()
	src := filepath.Join(dir, "driver.c")
	bin := filepath.Join(dir, "driver")
	if err := os.WriteFile(src, []byte(streamDriverC), 0644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("gcc", "-O2", "-I", hdr, "-o", bin, src, amalg)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("gcc -O2: %v\n%s", err, out)
	}
	return bin
}

func parseStreamMsgs(out []byte) []map[string]string {
	var msgs []map[string]string
	var cur map[string]string
	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if strings.HasPrefix(fields[0], "MSG=") {
			cur = map[string]string{}
			msgs = append(msgs, cur)
		}
		if cur == nil {
			continue
		}
		for _, f := range fields {
			if i := strings.IndexByte(f, '='); i > 0 {
				cur[f[:i]] = f[i+1:]
			}
		}
	}
	return msgs
}

func TestAEADStream_VsMonocypherC(t *testing.T) {
	bin := compileStreamDriver(t)
	key := streamKey()
	n24 := streamNonce24()
	sizes := streamSizes()
	inits := []struct {
		name  string
		flag  string
		nonce []byte
		init  func(*sgoi.Crypto_aead_ctx, []byte, []byte)
	}{
		{"init_x", "x", n24, sgoi.Crypto_aead_init_x},
		{"init_djb", "djb", n24[:8], sgoi.Crypto_aead_init_djb},
		{"init_ietf", "ietf", n24[:12], sgoi.Crypto_aead_init_ietf},
	}
	for _, in := range inits {
		for _, adn := range []int{0, 13} {
			t.Run(fmt.Sprintf("%s/ad%d/write", in.name, adn), func(t *testing.T) {
				out, err := exec.Command(bin, "write", in.flag, strconv.Itoa(adn)).CombinedOutput()
				if err != nil {
					t.Fatalf("c write: %v\n%s", err, out)
				}
				parsed := parseStreamMsgs(out)
				if len(parsed) != len(sizes) {
					t.Fatalf("c write msgs=%d want %d\n%s", len(parsed), len(sizes), out)
				}
				var ctx sgoi.Crypto_aead_ctx
				in.init(&ctx, key, in.nonce)
				var adPtr []byte
				if adn > 0 {
					adPtr = streamAD(adn)
				}
				for i, n := range sizes {
					pt := mkPT(n, "i%251")
					var ptPtr []byte
					if n > 0 {
						ptPtr = pt
					}
					ct := make([]byte, n)
					mac := make([]byte, 16)
					sgoi.Crypto_aead_write(&ctx, ct, mac, adPtr, uint64(adn), ptPtr, uint64(n))
					ctC, err := hex.DecodeString(parsed[i]["CT"])
					if err != nil {
						t.Fatalf("msg %d ct hex: %v", i, err)
					}
					macC, err := hex.DecodeString(parsed[i]["MAC"])
					if err != nil {
						t.Fatalf("msg %d mac hex: %v", i, err)
					}
					if !bytes.Equal(mac, macC) {
						t.Fatalf("msg %d n=%d mac diverge\ngo %x\nc  %s", i, n, mac, parsed[i]["MAC"])
					}
					if !bytes.Equal(ct, ctC) {
						t.Fatalf("msg %d n=%d ct diverge (go %d c %d)", i, n, len(ct), len(ctC))
					}
				}
			})

			t.Run(fmt.Sprintf("%s/ad%d/go_reads_c", in.name, adn), func(t *testing.T) {
				out, err := exec.Command(bin, "write", in.flag, strconv.Itoa(adn)).CombinedOutput()
				if err != nil {
					t.Fatalf("c write: %v\n%s", err, out)
				}
				parsed := parseStreamMsgs(out)
				if len(parsed) != len(sizes) {
					t.Fatalf("c write msgs=%d want %d", len(parsed), len(sizes))
				}
				var ctx sgoi.Crypto_aead_ctx
				in.init(&ctx, key, in.nonce)
				var adPtr []byte
				if adn > 0 {
					adPtr = streamAD(adn)
				}
				for i, n := range sizes {
					ctC, err := hex.DecodeString(parsed[i]["CT"])
					if err != nil {
						t.Fatalf("msg %d ct hex: %v", i, err)
					}
					macC, err := hex.DecodeString(parsed[i]["MAC"])
					if err != nil {
						t.Fatalf("msg %d mac hex: %v", i, err)
					}
					pt := make([]byte, n)
					var ptPtr, ctPtr []byte
					if n > 0 {
						ptPtr = pt
						ctPtr = ctC
					}
					st := sgoi.Crypto_aead_read(&ctx, ptPtr, macC, adPtr, uint64(adn), ctPtr, uint64(n))
					if st != 0 {
						t.Fatalf("go read c msg %d n=%d rejected status=%d", i, n, st)
					}
					want := mkPT(n, "i%251")
					if !bytes.Equal(pt, want) {
						t.Fatalf("go read c msg %d n=%d plaintext diverge", i, n)
					}
				}
			})

			t.Run(fmt.Sprintf("%s/ad%d/c_reads_go", in.name, adn), func(t *testing.T) {
				var ctx sgoi.Crypto_aead_ctx
				in.init(&ctx, key, in.nonce)
				var adPtr []byte
				if adn > 0 {
					adPtr = streamAD(adn)
				}
				var stdin bytes.Buffer
				for _, n := range sizes {
					pt := mkPT(n, "i%251")
					var ptPtr []byte
					if n > 0 {
						ptPtr = pt
					}
					ct := make([]byte, n)
					mac := make([]byte, 16)
					sgoi.Crypto_aead_write(&ctx, ct, mac, adPtr, uint64(adn), ptPtr, uint64(n))
					fmt.Fprintf(&stdin, "%s\n%s\n", hex.EncodeToString(ct), hex.EncodeToString(mac))
				}
				cmd := exec.Command(bin, "read", in.flag, strconv.Itoa(adn))
				cmd.Stdin = &stdin
				out, err := cmd.CombinedOutput()
				if err != nil {
					t.Fatalf("c read: %v\n%s", err, out)
				}
				parsed := parseStreamMsgs(out)
				if len(parsed) != len(sizes) {
					t.Fatalf("c read msgs=%d want %d\n%s", len(parsed), len(sizes), out)
				}
				for i, n := range sizes {
					st, _ := strconv.Atoi(parsed[i]["STATUS"])
					if st != 0 {
						t.Fatalf("c read go msg %d n=%d status=%d", i, n, st)
					}
					ptC, err := hex.DecodeString(parsed[i]["PT"])
					if err != nil {
						t.Fatalf("msg %d pt hex: %v", i, err)
					}
					want := mkPT(n, "i%251")
					if !bytes.Equal(ptC, want) {
						t.Fatalf("c read go msg %d n=%d plaintext diverge", i, n)
					}
				}
			})
		}
	}
}

func TestAEADStream_Write64Allocs(t *testing.T) {
	key := streamKey()
	nonce := streamNonce24()
	var ctx sgoi.Crypto_aead_ctx
	sgoi.Crypto_aead_init_x(&ctx, key, nonce)
	pt := mkPT(64, "i%251")
	ct := make([]byte, 64)
	mac := make([]byte, 16)
	n := testing.AllocsPerRun(200, func() {
		sgoi.Crypto_aead_write(&ctx, ct, mac, nil, 0, pt, 64)
	})
	if n != 0 {
		t.Fatalf("Crypto_aead_write 64 allocs = %v want 0", n)
	}
}
