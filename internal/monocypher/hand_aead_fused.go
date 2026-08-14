package monocypher

// Hand: Crypto_aead_write — boucle fusionnée 1-pass (Fused Streaming AEAD)
// Traite le chiffrement ChaCha20 et l'absorption Poly1305 par blocs de 256 octets
// pour maximiser la localité de cache L1 et masquer le temps de calcul Poly1305.
func Crypto_aead_write(ctx *Crypto_aead_ctx, cipher_text []byte, mac []byte, ad []byte, ad_size uint64, plain_text []byte, text_size uint64) {
	var _arr_v7 [64]uint8
	v7 := _arr_v7[:]

	// 1. Dérivation de la clé Poly1305 et clé de rekey (Bloc 0)
	Crypto_chacha20_djb(v7, nil, uint64(64), ctx.Key[:], ctx.Nonce[:], ctx.Counter)

	// 2. Initialisation de l'état Poly1305
	var polyCtx Crypto_poly1305_ctx
	Crypto_poly1305_init(&polyCtx, v7)

	// 3. Absorption de l'Additional Data (AD) et de son alignement sur 16 octets
	if ad_size > 0 {
		Crypto_poly1305_update(&polyCtx, ad, ad_size)
		gap := Gap(ad_size, uint64(16))
		if gap > 0 {
			Crypto_poly1305_update(&polyCtx, zero, gap)
		}
	}

	currCtr := ctx.Counter + 1
	offset := uint64(0)

	// 4. Boucle micro-entrelacée instruction-par-instruction (ChaCha20 + Poly1305 simultanés)
	if hasAVX2() && text_size >= 256 {
		numChunks := text_size / 256
		currCtr = aead_interleaved_write_simd(ctx, &polyCtx, cipher_text, plain_text, numChunks, currCtr)
		offset += numChunks * 256
	}

	// 5. Traitement du reliquat de payload (< 256 octets)
	remaining := text_size - offset
	if remaining > 0 {
		cOff := int(offset)
		currCtr = Crypto_chacha20_djb(cipher_text[cOff:], plain_text[cOff:], remaining, ctx.Key[:], ctx.Nonce[:], currCtr)
		Crypto_poly1305_update(&polyCtx, cipher_text[cOff:], remaining)
	}

	// 6. Alignement du ciphertext sur 16 octets
	gapText := Gap(text_size, uint64(16))
	if gapText > 0 {
		Crypto_poly1305_update(&polyCtx, zero, gapText)
	}

	// 7. Absorption des tailles ad_size (8o) et text_size (8o) en Little Endian
	var _arr_sizes [16]uint8
	sizes := _arr_sizes[:]
	Store64_le(sizes, ad_size)
	Store64_le(sizes[8:], text_size)
	Crypto_poly1305_update(&polyCtx, sizes, uint64(16))

	// 8. Finalisation du tag MAC
	Crypto_poly1305_final(&polyCtx, mac)

	// 9. Mise à jour de la clé de contexte pour le rekeying et hygiène mémoire
	for i := 0; i < 32; i++ {
		ctx.Key[i] = v7[32+i]
	}
	ctx.Counter = currCtr
	clear(v7)
	clear(sizes)
}
