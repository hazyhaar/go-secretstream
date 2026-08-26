// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55

// Hand: front fails on do-while + COPY/WIPE macros (crypto_elligator_key_pair).
// Algorithm: monocypher 4.0.2.

func Crypto_elligator_key_pair(hidden, secret_key, seed []byte) {
	var pk [32]byte
	var buf [64]byte
	copy(buf[32:], seed[:32])
	for {
		Crypto_chacha20_djb(buf[:], nil, 64, buf[32:], Zero[:], 0)
		Crypto_x25519_dirty_fast(pk[:], buf[:])
		if Crypto_elligator_rev(buf[32:], pk[:], buf[32]) == 0 {
			break
		}
	}
	Crypto_wipe(seed, 32)
	copy(hidden[:32], buf[32:64])
	copy(secret_key[:32], buf[:32])
	for i := range buf {
		buf[i] = 0
	}
	for i := range pk {
		pk[i] = 0
	}
}
