// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55_test

import (
	"os"
	"testing"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher55"
)

// TestSIMDSelectionNotSilent rend la sélection du chemin vectoriel
// observable et, sous la recette (SECRETSTREAM_REQUIRE_SIMD=1), obligatoire :
// une suite verte sur le repli scalaire alors que le chemin SIMD était attendu
// est un faux vert, pas un succès.
func TestSIMDSelectionNotSilent(t *testing.T) {
	active := sgoi.SIMDActive()
	t.Logf("monocypher55 SIMD path selected: %v", active)
	if os.Getenv("SECRETSTREAM_REQUIRE_SIMD") == "1" && !active {
		t.Fatal("SECRETSTREAM_REQUIRE_SIMD=1 : chemin SIMD attendu, repli scalaire actif (build sans GOEXPERIMENT=simd, hors amd64, ou CPU sans AVX2)")
	}
}
