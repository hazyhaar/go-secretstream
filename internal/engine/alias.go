// SPDX-License-Identifier: Apache-2.0 OR MIT

package engine

import "unsafe"

// anyOverlap and inexactOverlap follow golang.org/x/crypto/internal/alias
// (AnyOverlap / InexactOverlap). They are reimplemented locally: this module
// must not import that internal package.

func anyOverlap(x, y []byte) bool {
	return len(x) > 0 && len(y) > 0 &&
		uintptr(unsafe.Pointer(&x[0])) <= uintptr(unsafe.Pointer(&y[len(y)-1])) &&
		uintptr(unsafe.Pointer(&y[0])) <= uintptr(unsafe.Pointer(&x[len(x)-1]))
}

func inexactOverlap(x, y []byte) bool {
	if len(x) == 0 || len(y) == 0 || &x[0] == &y[0] {
		return false
	}
	return anyOverlap(x, y)
}

func checkAEADAlias(dst, payload, ad, mac []byte) error {
	if inexactOverlap(dst, payload) {
		return errInexactOverlap
	}
	if anyOverlap(dst, ad) {
		return errOverlapAD
	}
	if anyOverlap(dst, mac) {
		return errOverlapMAC
	}
	return nil
}
