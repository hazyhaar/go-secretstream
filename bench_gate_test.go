//go:build !race

// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55_test

import (
	"slices"
	"testing"
	"time"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher55"
)

// measureMedianThroughput exécute plusieurs séries de mesures pour éliminer la variance
// thermique et les fluctuations de charge CPU en retenant la médiane des débits.
func measureMedianThroughput(t *testing.T, fn func() error, size int, itersPerRun int, runs int) float64 {
	throughputs := make([]float64, runs)
	totalBytes := float64(size * itersPerRun)
	for r := 0; r < runs; r++ {
		start := time.Now()
		for i := 0; i < itersPerRun; i++ {
			if err := fn(); err != nil {
				t.Fatal(err)
			}
		}
		elapsed := time.Since(start)
		throughputs[r] = (totalBytes / (1024 * 1024)) / elapsed.Seconds()
	}
	slices.Sort(throughputs)
	return throughputs[runs/2]
}

// TestBenchGate_PerformanceFloor valide mécaniquement les seuils de performance plancher
// (Règle ARCHTIME n°6) avec filtre médian sur 5 runs pour interdire toute régression silencieuse.
func TestBenchGate_PerformanceFloor(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping bench gate in short mode")
	}
	// Le plancher est calibré sur le chemin AVX2 ; sur le repli scalaire le
	// gate n'a pas de sens et le dit (la recette exige le chemin SIMD par
	// SECRETSTREAM_REQUIRE_SIMD=1, voir TestSIMDSelectionNotSilent).
	if !sgoi.SIMDActive() {
		t.Skip("bench gate : plancher calibré SIMD, chemin scalaire actif — gate non applicable")
	}

	key, nonce := cmpKeyNonce()
	ad := []byte("HEADER_AD")
	const size = 1024 * 1024 // 1 Mo
	pt := cmpPT(size)
	dstCT := make([]byte, size)
	var mac [16]byte

	// Préchauffage (warmup 5 itérations) pour stabilisation du gouverneur CPU et des caches
	for i := 0; i < 5; i++ {
		if err := sgoi.LockDst(dstCT, mac[:], key, nonce, ad, pt); err != nil {
			t.Fatal(err)
		}
	}

	const (
		runs        = 5
		itersPerRun = 80
		floorThresh = 2000.0 // Mo/s plancher de sécurité (nominal 2 400 Mo/s, réf 8164952)
	)

	// 1. Banc Seal 1 Mo (Médiane 5 runs, plancher 2 000 Mo/s, 0 alloc — Réf 8164952)
	mbsSeal := measureMedianThroughput(t, func() error {
		return sgoi.LockDst(dstCT, mac[:], key, nonce, ad, pt)
	}, size, itersPerRun, runs)

	t.Logf("BenchGate Seal 1MB: %.2f MB/s (floor: %.2f MB/s, median 5 runs, ref 8164952)", mbsSeal, floorThresh)
	if mbsSeal < floorThresh {
		t.Errorf("FAIL BENCH GATE: Seal 1MB throughput %.2f MB/s < %.2f MB/s floor", mbsSeal, floorThresh)
	}

	// 2. Banc Open 1 Mo (Médiane 5 runs, plancher 2 000 Mo/s, 0 alloc — Fusion 1-pass)
	dstPT := make([]byte, size)
	mbsOpen := measureMedianThroughput(t, func() error {
		return sgoi.UnlockDst(dstPT, key, nonce, mac[:], ad, dstCT)
	}, size, itersPerRun, runs)

	t.Logf("BenchGate Open 1MB: %.2f MB/s (floor: %.2f MB/s, median 5 runs, 1-pass fused)", mbsOpen, floorThresh)
	if mbsOpen < floorThresh {
		t.Errorf("FAIL BENCH GATE: Open 1MB throughput %.2f MB/s < %.2f MB/s floor", mbsOpen, floorThresh)
	}
}
