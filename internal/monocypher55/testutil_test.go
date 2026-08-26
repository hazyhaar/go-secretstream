// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

func findMonocypherC(t *testing.T) (amalg, hdr string) {
	t.Helper()
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("Compilateur gcc non disponible")
	}
	// 1. Variable d'environnement prioritaire (portabilité conteneurs / hôtes externes)
	if env := os.Getenv("MONOCYPHER_C_SRC"); env != "" {
		if stat, err := os.Stat(env); err == nil {
			if stat.IsDir() {
				amalgFile := filepath.Join(env, "monocypher.c")
				if _, err := os.Stat(amalgFile); err == nil {
					return amalgFile, env
				}
			} else {
				return env, filepath.Dir(env)
			}
		}
	}

	// 2. Découverte relative depuis le fichier source actuel
	_, file, _, ok := runtime.Caller(0)
	if ok {
		candidates := []string{
			filepath.Join(filepath.Dir(file), "..", "..", "c2simd", "spec", "c_sources", "upstream", "monocypher", "4.0.2"),
			filepath.Join(filepath.Dir(file), "testdata", "monocypher"),
		}
		for _, c := range candidates {
			amalgFile := filepath.Join(c, "monocypher.c")
			if _, err := os.Stat(amalgFile); err == nil {
				return amalgFile, c
			}
		}
	}

	t.Skip("Oracle C Monocypher non disponible (définir MONOCYPHER_C_SRC ou fournir spec/c_sources)")
	return "", ""
}

func findC2simdRoot(t *testing.T) string {
	t.Helper()
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("Compilateur gcc non disponible")
	}
	if env := os.Getenv("C2SIMD_ROOT"); env != "" {
		if _, err := os.Stat(filepath.Join(env, "sgoiter")); err == nil {
			return env
		}
	}
	_, file, _, ok := runtime.Caller(0)
	if ok {
		c2 := filepath.Join(filepath.Dir(file), "..", "..", "c2simd")
		if _, err := os.Stat(filepath.Join(c2, "sgoiter")); err == nil {
			return c2
		}
	}
	t.Skip("Racine c2simd non disponible (définir C2SIMD_ROOT)")
	return ""
}

func keyNonce() (key, nonce []byte) {
	key = make([]byte, 32)
	nonce = make([]byte, 24)
	for i := range key {
		key[i] = byte(i + 1)
	}
	for i := range nonce {
		nonce[i] = byte(i + 10)
	}
	return key, nonce
}

func mkPT(n int, pattern string) []byte {
	pt := make([]byte, n)
	switch pattern {
	case "zero":
		// leave zeros
	case "i%251":
		for i := range pt {
			pt[i] = byte(i % 251)
		}
	case "(i*17+3)%251":
		for i := range pt {
			pt[i] = byte((i*17 + 3) % 251)
		}
	default:
		for i := range pt {
			pt[i] = byte((i*17 + 3) % 251)
		}
	}
	return pt
}
