package monocypher55

// xorStreamShort masque dst = src XOR ks sur n octets (n <= 64), par mots de
// 8 octets puis reliquat, sans test de nil par octet. src == nil => dst = ks.
func xorStreamShort(dst []byte, src []byte, ks []byte, n uint64) {
	if src == nil {
		copy(dst[:n], ks[:n])
		return
	}
	i := uint64(0)
	for ; i+8 <= n; i += 8 {
		Store64_le(dst[i:], Load64_le(src[i:])^Load64_le(ks[i:]))
	}
	for ; i < n; i++ {
		dst[i] = src[i] ^ ks[i]
	}
}

// Hand: Crypto_aead_write — boucle fusionnée 1-pass (Fused Streaming AEAD)
// Traite le chiffrement ChaCha20 et l'absorption Poly1305 par blocs de 256 octets
// pour maximiser la localité de cache L1 et masquer le temps de calcul Poly1305.
func Crypto_aead_write(ctx *Crypto_aead_ctx, cipher_text []byte, mac []byte, ad []byte, ad_size uint64, plain_text []byte, text_size uint64) {
	// Fast-path pour petites tailles (<= 64 octets) : dérivation unifiée ChaCha20 en un seul appel
	if text_size <= 64 {
		var _arr_buf [128]uint8
		buf := _arr_buf[:]
		derivSize := uint64(64)
		if text_size > 0 {
			derivSize = 128
		}
		if derivSize == 128 && hasAVX2() {
			// Dérivation des 2 blocs (compteurs ctr et ctr+1) en une seule passe
			// de rounds 256-bit : 2 lanes de bloc côte à côte dans des Uint32x8.
			chacha20_deriv2_simd256(buf, ctx.Key[:], ctx.Nonce[:], ctx.Counter)
		} else {
			Crypto_chacha20_djb(buf, nil, derivSize, ctx.Key[:], ctx.Nonce[:], ctx.Counter)
		}

		// Masquage direct du payload court par XOR avec le bloc 1
		xorStreamShort(cipher_text, plain_text, buf[64:], text_size)

		// Initialisation et absorption Poly1305
		var polyCtx Crypto_poly1305_ctx
		Crypto_poly1305_init(&polyCtx, buf[:64])
		if ad_size > 0 {
			Crypto_poly1305_update(&polyCtx, ad, ad_size)
			if gap := Gap(ad_size, 16); gap > 0 {
				Crypto_poly1305_update(&polyCtx, zero, gap)
			}
		}
		if text_size > 0 {
			Crypto_poly1305_update(&polyCtx, cipher_text[:text_size], text_size)
			if gapText := Gap(text_size, 16); gapText > 0 {
				Crypto_poly1305_update(&polyCtx, zero, gapText)
			}
		}
		var _arr_sizes [16]uint8
		sizes := _arr_sizes[:]
		Store64_le(sizes, ad_size)
		Store64_le(sizes[8:], text_size)
		Crypto_poly1305_update(&polyCtx, sizes, 16)
		Crypto_poly1305_final(&polyCtx, mac)

		// Rekeying et avancement du compteur
		for i := 0; i < 32; i++ {
			ctx.Key[i] = buf[32+i]
		}
		if text_size > 0 {
			ctx.Counter += 2
		} else {
			ctx.Counter += 1
		}
		clear(buf)
		clear(sizes)
		polyCtx = Crypto_poly1305_ctx{}
		return
	}

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
	polyCtx = Crypto_poly1305_ctx{}
}

// Hand: Crypto_aead_read — boucle fusionnée 1-pass (Fused Streaming AEAD Decryption)
// avec contrat strict WIPE-ON-FAILURE (effacement de plain_text si le MAC est invalide).
func Crypto_aead_read(ctx *Crypto_aead_ctx, plain_text []byte, mac []byte, ad []byte, ad_size uint64, cipher_text []byte, text_size uint64) int {
	// Fast-path pour petites tailles (<= 64 octets)
	if text_size <= 64 {
		var _arr_buf [128]uint8
		buf := _arr_buf[:]
		derivSize := uint64(64)
		if text_size > 0 {
			derivSize = 128
		}
		if derivSize == 128 && hasAVX2() {
			// Même dérivation vectorielle 2 blocs que la voie write.
			chacha20_deriv2_simd256(buf, ctx.Key[:], ctx.Nonce[:], ctx.Counter)
		} else {
			Crypto_chacha20_djb(buf, nil, derivSize, ctx.Key[:], ctx.Nonce[:], ctx.Counter)
		}

		// Initialisation et absorption Poly1305 sur cipher_text
		var polyCtx Crypto_poly1305_ctx
		Crypto_poly1305_init(&polyCtx, buf[:64])
		if ad_size > 0 {
			Crypto_poly1305_update(&polyCtx, ad, ad_size)
			if gap := Gap(ad_size, 16); gap > 0 {
				Crypto_poly1305_update(&polyCtx, zero, gap)
			}
		}
		if text_size > 0 {
			Crypto_poly1305_update(&polyCtx, cipher_text[:text_size], text_size)
			if gapText := Gap(text_size, 16); gapText > 0 {
				Crypto_poly1305_update(&polyCtx, zero, gapText)
			}
		}
		var _arr_sizes [16]uint8
		sizes := _arr_sizes[:]
		Store64_le(sizes, ad_size)
		Store64_le(sizes[8:], text_size)
		Crypto_poly1305_update(&polyCtx, sizes, 16)
		var expectedMac [16]byte
		Crypto_poly1305_final(&polyCtx, expectedMac[:])

		diff := Crypto_verify16(mac, expectedMac[:])
		if diff == 0 {
			// Tag authentique : déchiffrement et rekeying
			if plain_text != nil && cipher_text != nil {
				xorStreamShort(plain_text, cipher_text, buf[64:], text_size)
			}
			copy(ctx.Key[:], buf[32:64])
			if text_size > 0 {
				ctx.Counter += 2
			} else {
				ctx.Counter += 1
			}
		} else {
			// WIPE-ON-FAILURE : effacement systématique du plaintext
			if plain_text != nil && text_size > 0 {
				clear(plain_text[:text_size])
			}
		}

		clear(buf)
		clear(sizes)
		clear(expectedMac[:])
		polyCtx = Crypto_poly1305_ctx{}
		return diff
	}

	var _arr_v7 [64]uint8
	v7 := _arr_v7[:]

	// 1. Dérivation de la clé Poly1305 et clé de rekey (Bloc 0)
	Crypto_chacha20_djb(v7, nil, uint64(64), ctx.Key[:], ctx.Nonce[:], ctx.Counter)

	// 2. Initialisation de l'état Poly1305
	var polyCtx Crypto_poly1305_ctx
	Crypto_poly1305_init(&polyCtx, v7)

	// 3. Absorption de l'Additional Data (AD)
	if ad_size > 0 {
		Crypto_poly1305_update(&polyCtx, ad, ad_size)
		if gap := Gap(ad_size, uint64(16)); gap > 0 {
			Crypto_poly1305_update(&polyCtx, zero, gap)
		}
	}

	currCtr := ctx.Counter + 1
	offset := uint64(0)

	// 4. Boucle micro-entrelacée 1-pass (ChaCha20 déchiffrement + Poly1305 absorption simultanés)
	// Contrainte in-place : le cœur SIMD absorbe le Poly1305 du chunk c-1 depuis
	// `cipher` APRÈS que l'itération c-1 l'a recouvert de plaintext. En alias exact
	// (plain_text == cipher_text, garanti par Monocypher C), ce recouvrement corromprait
	// l'absorption : tout le payload passe alors par le chemin scalaire, qui absorbe
	// le ciphertext AVANT d'écrire le plaintext.
	inPlace := len(plain_text) > 0 && len(cipher_text) > 0 && &plain_text[0] == &cipher_text[0]
	if hasAVX2() && text_size >= 256 && !inPlace {
		numChunks := text_size / 256
		currCtr = aead_interleaved_read_simd(ctx, &polyCtx, plain_text, cipher_text, numChunks, currCtr)
		offset += numChunks * 256
	}

	// 5. Traitement du reliquat (< 256 octets)
	remaining := text_size - offset
	if remaining > 0 {
		cOff := int(offset)
		Crypto_poly1305_update(&polyCtx, cipher_text[cOff:], remaining)
		currCtr = Crypto_chacha20_djb(plain_text[cOff:], cipher_text[cOff:], remaining, ctx.Key[:], ctx.Nonce[:], currCtr)
	}

	// 6. Alignement du ciphertext sur 16 octets
	if gapText := Gap(text_size, uint64(16)); gapText > 0 {
		Crypto_poly1305_update(&polyCtx, zero, gapText)
	}

	// 7. Absorption des tailles
	var _arr_sizes [16]uint8
	sizes := _arr_sizes[:]
	Store64_le(sizes, ad_size)
	Store64_le(sizes[8:], text_size)
	Crypto_poly1305_update(&polyCtx, sizes, uint64(16))

	// 8. Calcul et vérification du tag MAC
	var expectedMac [16]byte
	Crypto_poly1305_final(&polyCtx, expectedMac[:])
	diff := Crypto_verify16(mac, expectedMac[:])

	if diff == 0 {
		// Tag authentique : mise à jour de la clé de rekeying et du compteur
		for i := 0; i < 32; i++ {
			ctx.Key[i] = v7[32+i]
		}
		ctx.Counter = currCtr
	} else {
		// WIPE-ON-FAILURE : effacement complet du plaintext déchiffré spéculativement
		if plain_text != nil && text_size > 0 {
			clear(plain_text[:text_size])
		}
	}

	clear(v7)
	clear(sizes)
	clear(expectedMac[:])
	polyCtx = Crypto_poly1305_ctx{}
	return diff
}
