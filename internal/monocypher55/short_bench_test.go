// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55_test

import (
	"testing"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher55"
)

// BenchmarkStratum_AEAD_LockDst_32 mesure le fast-path AEAD sur un payload de 32 octets.
func BenchmarkStratum_AEAD_LockDst_32(b *testing.B) {
	key, nonce := keyNonce()
	pt := mkPT(32, "i%251")
	ad := []byte("ad")
	dst := make([]byte, len(pt))
	var mac [16]byte
	b.SetBytes(32)
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		if err := sgoi.LockDst(dst, mac[:], key, nonce, ad, pt); err != nil {
			b.Fatal(err)
		}
	}
}
