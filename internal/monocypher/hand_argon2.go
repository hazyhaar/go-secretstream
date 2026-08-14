package monocypher

import "unsafe"

// Hand: crypto_argon2 — monocypher 4.0.2 (front still trips on volatile u64* casts).
// Uses emitted helpers: Extended_hash, G_rounds, Copy_block, Xor_block, Blake_*.

const (
	Crypto_argon2_d  uint32 = 0
	Crypto_argon2_i  uint32 = 1
	Crypto_argon2_id uint32 = 2
)

// Crypto_argon2_no_extras — monocypher global.
var Crypto_argon2_no_extras = Crypto_argon2_extras{}

// Crypto_argon2 fills hash from password inputs using work_area (must be ≥ nb_blocks*1024 bytes).
// Matches monocypher: work_area is the sole block store (no hidden heap for blocks).
func Crypto_argon2(hash []byte, hash_size uint32, work_area []byte, config Crypto_argon2_config, inputs Crypto_argon2_inputs, extras Crypto_argon2_extras) {
	segment_size := config.Nb_blocks / config.Nb_lanes / 4
	lane_size := segment_size * 4
	nb_blocks := lane_size * config.Nb_lanes

	need := int(nb_blocks) * 1024
	if len(work_area) < need {
		// Fail closed: C requires a suitable work_area; do not silently allocate a second store.
		panic("monocypher55: Crypto_argon2 work_area too small")
	}
	if uintptr(unsafe.Pointer(unsafe.SliceData(work_area)))%8 != 0 {
		panic("monocypher55: Crypto_argon2 work_area must be 8-byte aligned")
	}
	// Blk is [128]uint64 — overlay on caller buffer (same layout as C blk*).
	blocks := unsafe.Slice((*Blk)(unsafe.Pointer(unsafe.SliceData(work_area))), int(nb_blocks))
	{
		var initial_hash [72]byte
		var ctx Crypto_blake2b_ctx
		Crypto_blake2b_init(&ctx, 64)
		Blake_update_32(&ctx, config.Nb_lanes)
		Blake_update_32(&ctx, hash_size)
		Blake_update_32(&ctx, config.Nb_blocks)
		Blake_update_32(&ctx, config.Nb_passes)
		Blake_update_32(&ctx, 0x13)
		Blake_update_32(&ctx, config.Algorithm)
		Blake_update_32_buf(&ctx, inputs.Pass, inputs.Pass_size)
		Blake_update_32_buf(&ctx, inputs.Salt, inputs.Salt_size)
		Blake_update_32_buf(&ctx, extras.Key, extras.Key_size)
		Blake_update_32_buf(&ctx, extras.Ad, extras.Ad_size)
		Crypto_blake2b_final(&ctx, initial_hash[:])

		var hash_area [1024]byte
		for l := uint32(0); l < config.Nb_lanes; l++ {
			for i := uint32(0); i < 2; i++ {
				Store32_le(initial_hash[64:], i)
				Store32_le(initial_hash[68:], l)
				Extended_hash(hash_area[:], 1024, initial_hash[:], 72)
				Load64_le_buf(blocks[l*lane_size+i].A[:], hash_area[:], 128)
			}
		}
		for i := range initial_hash {
			initial_hash[i] = 0
		}
		for i := range hash_area {
			hash_area[i] = 0
		}
	}

	constant_time := config.Algorithm != Crypto_argon2_d
	var tmp Blk
	for pass := uint32(0); pass < config.Nb_passes; pass++ {
		for slice := uint32(0); slice < 4; slice++ {
			pass_offset := uint32(0)
			if pass == 0 && slice == 0 {
				pass_offset = 2
			}
			slice_offset := slice * segment_size
			if slice == 2 && config.Algorithm == Crypto_argon2_id {
				constant_time = false
			}
			for segment := uint32(0); segment < config.Nb_lanes; segment++ {
				var index_block Blk
				index_ctr := uint32(1)
				for block := pass_offset; block < segment_size; block++ {
					lane_offset := segment * lane_size
					segStart := lane_offset + slice_offset
					curIdx := segStart + block
					var prevIdx uint32
					if block == 0 && slice_offset == 0 {
						prevIdx = segStart + lane_size - 1
					} else {
						prevIdx = segStart + block - 1
					}

					var index_seed uint64
					if constant_time {
						if block == pass_offset || (block%128) == 0 {
							for i := range index_block.A {
								index_block.A[i] = 0
							}
							index_block.A[0] = uint64(pass)
							index_block.A[1] = uint64(segment)
							index_block.A[2] = uint64(slice)
							index_block.A[3] = uint64(nb_blocks)
							index_block.A[4] = uint64(config.Nb_passes)
							index_block.A[5] = uint64(config.Algorithm)
							index_block.A[6] = uint64(index_ctr)
							index_ctr++
							Copy_block(&tmp, &index_block)
							G_rounds(&index_block)
							Xor_block(&index_block, &tmp)
							Copy_block(&tmp, &index_block)
							G_rounds(&index_block)
							Xor_block(&index_block, &tmp)
						}
						index_seed = index_block.A[block%128]
					} else {
						index_seed = blocks[prevIdx].A[0]
					}

					next_slice := ((slice + 1) % 4) * segment_size
					window_start := uint32(0)
					if pass != 0 {
						window_start = next_slice
					}
					nb_segments := slice
					if pass != 0 {
						nb_segments = 3
					}
					lane := uint64(segment)
					if !(pass == 0 && slice == 0) {
						lane = (index_seed >> 32) % uint64(config.Nb_lanes)
					}
					var window_size uint32
					if lane == uint64(segment) {
						window_size = nb_segments*segment_size + (block - 1)
					} else if block == 0 {
						window_size = nb_segments*segment_size + ^uint32(0)
					} else {
						window_size = nb_segments * segment_size
					}

					j1 := index_seed & 0xffffffff
					x := (j1 * j1) >> 32
					y := (uint64(window_size) * x) >> 32
					z := (uint64(window_size) - 1) - y
					ref := (uint64(window_start) + z) % uint64(lane_size)
					index := uint32(lane)*lane_size + uint32(ref)

					Copy_block(&tmp, &blocks[prevIdx])
					Xor_block(&tmp, &blocks[index])
					if pass == 0 {
						Copy_block(&blocks[curIdx], &tmp)
					} else {
						Xor_block(&blocks[curIdx], &tmp)
					}
					G_rounds(&tmp)
					Xor_block(&blocks[curIdx], &tmp)
				}
			}
		}
	}

	for i := range tmp.A {
		tmp.A[i] = 0
	}

	last := int(lane_size - 1)
	for lane := uint32(1); lane < config.Nb_lanes; lane++ {
		next := last + int(lane_size)
		Xor_block(&blocks[next], &blocks[last])
		last = next
	}

	var final_block [1024]byte
	Store64_le_buf(final_block[:], blocks[last].A[:], 128)

	// wipe work_area (blocks overlay)
	for i := range work_area[:need] {
		work_area[i] = 0
	}

	Extended_hash(hash, hash_size, final_block[:], 1024)
	for i := range final_block {
		final_block[i] = 0
	}
}
