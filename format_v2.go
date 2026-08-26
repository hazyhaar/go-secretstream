// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55

import "encoding/binary"

//go:generate go run ./internal/cmd/formatgen

func isMagicV2(b []byte) bool {
	return len(b) >= 8 &&
		b[0] == magicV2[0] && b[1] == magicV2[1] &&
		b[2] == magicV2[2] && b[3] == magicV2[3] &&
		b[4] == magicV2[4] && b[5] == magicV2[5] &&
		b[6] == magicV2[6] && b[7] == magicV2[7]
}

func writeHeaderV2(dst []byte, nonce *[24]byte) {
	copy(dst[0:8], magicV2[:])
	binary.BigEndian.PutUint16(dst[8:10], VersionV2)
	binary.BigEndian.PutUint16(dst[10:12], FlagsV2)
	copy(dst[12:36], nonce[:])
}

func frameNonceV2(nonce *[24]byte, seq uint64) [12]byte {
	var n [12]byte
	copy(n[0:4], nonce[16:20])
	binary.BigEndian.PutUint64(n[4:12], seq)
	return n
}

func frameNonceV1(nonce *[24]byte, seq uint64) [12]byte {
	var n [12]byte
	base := binary.BigEndian.Uint64(nonce[16:24])
	binary.BigEndian.PutUint64(n[4:12], base^seq)
	return n
}

// bindChunkADv2 assemble
// AD = "SS55-v2\x00" || seq_be64 || tag || len(ad_appelant)_be32 || ad_appelant.
// Le préfixe de 21 octets tient dans prefix, sans allocation si ad est vide.
func bindChunkADv2(prefix *[adPrefixV2Len]byte, ext *[]byte, seq uint64, tag byte, ad []byte) []byte {
	copy(prefix[:8], magicV2[:])
	binary.BigEndian.PutUint64(prefix[8:16], seq)
	prefix[16] = tag
	binary.BigEndian.PutUint32(prefix[17:21], uint32(len(ad)))
	if len(ad) == 0 {
		return prefix[:]
	}
	need := adPrefixV2Len + len(ad)
	if cap(*ext) < need {
		*ext = make([]byte, need)
	}
	buf := (*ext)[:need]
	copy(buf[:adPrefixV2Len], prefix[:])
	copy(buf[adPrefixV2Len:], ad)
	*ext = buf
	return buf
}
