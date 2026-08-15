package monocypher55

// Hand: Crypto_x25519_dirty_small — calcul X25519 non canonique sur scalaire de pile (évite la mutation globale).
func Crypto_x25519_dirty_small(public_key []byte, secret_key []byte) {
	var _arr_scalar [32]byte
	scalar := _arr_scalar[:]
	Crypto_eddsa_trim_scalar(scalar, secret_key)
	Add_xl(scalar, secret_key[0])
	Scalarmult(public_key, scalar, dirty_base_point[:], 0x100)
	clear(scalar)
}
