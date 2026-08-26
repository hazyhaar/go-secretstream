// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55

// Hand overrides — survive regen of monocypher_aead_sgoiter.go (sgoiter emit gaps).

type Slide_ctx struct {
	// Hand-fix widths: C monocypher uses i16 next_index, i8 next_digit, u8 next_check.
	Next_index int16
	Next_digit int8
	Next_check uint8
}

func Slide_init(ctx *Slide_ctx, scalar []byte) {
	i := 252
	for i > 0 && scalar_bit(scalar, i) == 0 {
		i--
	}
	ctx.Next_check = uint8(i + 1)
	ctx.Next_index = -1
	ctx.Next_digit = -1
}

// Slide_step — monocypher 4.0.2 (signed window digit).

func Slide_step(ctx *Slide_ctx, width int, i int, scalar []byte) int {
	if i == int(ctx.Next_check) {
		if scalar_bit(scalar, i) == scalar_bit(scalar, i-1) {
			ctx.Next_check--
		} else {
			w := width
			if i+1 < w {
				w = i + 1
			}
			v := -(scalar_bit(scalar, i) << (w - 1))
			for j := 0; j < w-1; j++ {
				v += scalar_bit(scalar, i-(w-1)+j) << j
			}
			v += scalar_bit(scalar, i-w)
			lsb := v & (^v + 1)
			s := 0
			if lsb&0xAA != 0 {
				s |= 1
			}
			if lsb&0xCC != 0 {
				s |= 2
			}
			if lsb&0xF0 != 0 {
				s |= 4
			}
			ctx.Next_index = int16(i - (w - 1) + s)
			ctx.Next_digit = int8(v >> s)
			ctx.Next_check -= uint8(w)
		}
	}
	if i == int(ctx.Next_index) {
		return int(ctx.Next_digit)
	}
	return 0
}
