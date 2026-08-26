// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55

// SIMDActive rapporte si le chemin vectoriel AVX2 est effectivement
// sélectionné pour ce binaire sur cette machine : compilé sous
// goexperiment.simd && amd64 ET AVX2 présent au runtime. Faux dans tout
// autre cas (repli scalaire), ce qui permet à un consommateur ou à une
// recette de test d'exiger explicitement le chemin SIMD au lieu de tester le
// repli par accident.
func SIMDActive() bool {
	return hasAVX2()
}
