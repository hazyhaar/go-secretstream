// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"
)

func v2FixedKey(t *testing.T) []byte {
	t.Helper()
	key := make([]byte, 32)
	for i := range key {
		key[i] = byte(0xA5 ^ i*17)
	}
	return key
}

func v2Seal(t *testing.T, key []byte, frags [][]byte, ads [][]byte) []byte {
	t.Helper()
	var buf bytes.Buffer
	enc, err := NewEncryptor(&buf, key)
	if err != nil {
		t.Fatalf("NewEncryptor: %v", err)
	}
	for i, f := range frags {
		var ad []byte
		if ads != nil {
			ad = ads[i]
		}
		if _, err := enc.WriteWithAD(f, ad); err != nil {
			t.Fatalf("WriteWithAD %d: %v", i, err)
		}
	}
	if err := enc.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	return buf.Bytes()
}

func splitV2(t *testing.T, wire []byte) (hdr []byte, frames [][]byte) {
	t.Helper()
	if len(wire) < HeaderSizeV2 {
		t.Fatalf("fil v2 trop court: %d", len(wire))
	}
	hdr = append([]byte(nil), wire[:HeaderSizeV2]...)
	rest := wire[HeaderSizeV2:]
	for len(rest) > 0 {
		if len(rest) < 4 {
			t.Fatalf("préfixe de longueur tronqué, reste %d", len(rest))
		}
		n := int(binary.BigEndian.Uint32(rest[:4]))
		total := 4 + n
		if n < 0 || len(rest) < total {
			t.Fatalf("trame tronquée: besoin %d, reste %d", total, len(rest))
		}
		frames = append(frames, append([]byte(nil), rest[:total]...))
		rest = rest[total:]
	}
	return hdr, frames
}

func joinV2(hdr []byte, frames ...[]byte) []byte {
	var out bytes.Buffer
	out.Write(hdr)
	for _, f := range frames {
		out.Write(f)
	}
	return out.Bytes()
}

func assertHardFail(t *testing.T, dec *Decryptor, p, ad []byte) {
	t.Helper()
	motif := byte(0xA5)
	for i := range p {
		p[i] = motif
	}
	n, err := dec.ReadWithAD(p, ad)
	if err == nil || err == io.EOF {
		t.Fatalf("échec dur attendu, obtenu err=%v n=%d", err, n)
	}
	if n != 0 {
		t.Fatalf("échec dur a déclaré %d octets", n)
	}
	n2, err2 := dec.ReadWithAD(p, ad)
	if err2 == nil || n2 != 0 {
		t.Fatalf("erreur collante attendue: n2=%d err2=%v", n2, err2)
	}
	if err2.Error() != err.Error() {
		t.Fatalf("erreur collante différente: première %v, seconde %v", err, err2)
	}
}

func TestV2Attack_TruncationAfterCompleteFrameNoFinal(t *testing.T) {
	key := v2FixedKey(t)
	plain := bytes.Repeat([]byte{0x11}, 200)
	wire := v2Seal(t, key, [][]byte{plain}, nil)
	hdr, frames := splitV2(t, wire)
	if len(frames) < 2 {
		t.Fatalf("attendu message+final, obtenu %d", len(frames))
	}
	truncated := joinV2(hdr, frames[0])
	dec, err := NewDecryptor(bytes.NewReader(truncated), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	got := make([]byte, len(plain))
	n, err := dec.Read(got)
	if err != nil || n != len(plain) || !bytes.Equal(got, plain) {
		t.Fatalf("première trame: n=%d err=%v", n, err)
	}
	assertHardFail(t, dec, make([]byte, 64), nil)
}

func TestV2Attack_TruncationMidFrame(t *testing.T) {
	key := v2FixedKey(t)
	plain := bytes.Repeat([]byte{0x22}, 400)
	wire := v2Seal(t, key, [][]byte{plain}, nil)
	if len(wire) < HeaderSizeV2+20 {
		t.Fatalf("fil trop court: %d", len(wire))
	}
	truncated := wire[:HeaderSizeV2+4+20]
	dec, err := NewDecryptor(bytes.NewReader(truncated), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	assertHardFail(t, dec, make([]byte, len(plain)), nil)
}

func TestV2Attack_DeleteIntermediateFrame(t *testing.T) {
	key := v2FixedKey(t)
	a := bytes.Repeat([]byte{0x01}, 80)
	b := bytes.Repeat([]byte{0x02}, 80)
	c := bytes.Repeat([]byte{0x03}, 80)
	wire := v2Seal(t, key, [][]byte{a, b, c}, nil)
	hdr, frames := splitV2(t, wire)
	if len(frames) != 4 {
		t.Fatalf("attendu 3 messages + final, obtenu %d", len(frames))
	}
	deleted := joinV2(hdr, frames[0], frames[2], frames[3])
	dec, err := NewDecryptor(bytes.NewReader(deleted), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	got := make([]byte, len(a))
	if n, err := dec.Read(got); err != nil || n != len(a) {
		t.Fatalf("première trame: n=%d err=%v", n, err)
	}
	assertHardFail(t, dec, make([]byte, len(c)), nil)
}

func TestV2Attack_PermuteTwoFrames(t *testing.T) {
	key := v2FixedKey(t)
	a := bytes.Repeat([]byte{0x11}, 80)
	b := bytes.Repeat([]byte{0x22}, 80)
	wire := v2Seal(t, key, [][]byte{a, b}, nil)
	hdr, frames := splitV2(t, wire)
	swapped := joinV2(hdr, frames[1], frames[0], frames[2])
	dec, err := NewDecryptor(bytes.NewReader(swapped), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	assertHardFail(t, dec, make([]byte, len(b)), nil)
}

func TestV2Attack_ReplayFrame(t *testing.T) {
	key := v2FixedKey(t)
	a := bytes.Repeat([]byte{0x31}, 80)
	b := bytes.Repeat([]byte{0x32}, 80)
	wire := v2Seal(t, key, [][]byte{a, b}, nil)
	hdr, frames := splitV2(t, wire)
	replayed := joinV2(hdr, frames[0], frames[0], frames[2])
	dec, err := NewDecryptor(bytes.NewReader(replayed), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	got := make([]byte, len(a))
	if n, err := dec.Read(got); err != nil || n != len(a) {
		t.Fatalf("première trame: n=%d err=%v", n, err)
	}
	assertHardFail(t, dec, make([]byte, len(a)), nil)
}

func TestV2Attack_TagMessageToFinalOnWire(t *testing.T) {
	key := v2FixedKey(t)
	plain := bytes.Repeat([]byte{0x41}, 80)
	wire := v2Seal(t, key, [][]byte{plain}, nil)
	hdr, frames := splitV2(t, wire)
	frames[0][4] = TagFinal
	tampered := joinV2(hdr, frames...)
	dec, err := NewDecryptor(bytes.NewReader(tampered), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	assertHardFail(t, dec, make([]byte, len(plain)), nil)
}

func TestV2Attack_TagFinalToMessageOnWire(t *testing.T) {
	key := v2FixedKey(t)
	plain := bytes.Repeat([]byte{0x42}, 80)
	wire := v2Seal(t, key, [][]byte{plain}, nil)
	hdr, frames := splitV2(t, wire)
	final := frames[len(frames)-1]
	final[4] = TagMessage
	tampered := joinV2(hdr, frames...)
	dec, err := NewDecryptor(bytes.NewReader(tampered), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	got := make([]byte, len(plain))
	if n, err := dec.Read(got); err != nil || n != len(plain) {
		t.Fatalf("message: n=%d err=%v", n, err)
	}
	assertHardFail(t, dec, make([]byte, 16), nil)
}

func TestV2Attack_ADInjectivity_ab_c_vs_a_bc(t *testing.T) {
	key := v2FixedKey(t)
	p1 := []byte("fragment-un")
	p2 := []byte("fragment-deux")
	wireAB_C := v2Seal(t, key, [][]byte{p1, p2}, [][]byte{[]byte("ab"), []byte("c")})

	dec, err := NewDecryptor(bytes.NewReader(wireAB_C), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	got := make([]byte, len(p1))
	if n, err := dec.ReadWithAD(got, []byte("a")); err == nil {
		t.Fatalf("découpage (a, bc) accepté sur un flux (ab, c): n=%d", n)
	}
	if n, err := dec.ReadWithAD(got, []byte("a")); err == nil || n != 0 {
		t.Fatalf("collant attendu après mauvais découpage: n=%d err=%v", n, err)
	}

	decOK, err := NewDecryptor(bytes.NewReader(wireAB_C), key)
	if err != nil {
		t.Fatalf("NewDecryptor ok: %v", err)
	}
	if n, err := decOK.ReadWithAD(got, []byte("ab")); err != nil || n != len(p1) {
		t.Fatalf("découpage d'origine: n=%d err=%v", n, err)
	}
}

func TestV2Attack_AlteredFrameLength(t *testing.T) {
	key := v2FixedKey(t)
	plain := bytes.Repeat([]byte{0x51}, 80)
	wire := v2Seal(t, key, [][]byte{plain}, nil)
	hdr, frames := splitV2(t, wire)
	binary.BigEndian.PutUint32(frames[0][:4], uint32(len(frames[0])-4+8))
	tampered := joinV2(hdr, frames...)
	dec, err := NewDecryptor(bytes.NewReader(tampered), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	assertHardFail(t, dec, make([]byte, len(plain)+16), nil)
}

func TestV2Attack_AlteredMAC(t *testing.T) {
	key := v2FixedKey(t)
	plain := bytes.Repeat([]byte{0x61}, 80)
	wire := v2Seal(t, key, [][]byte{plain}, nil)
	hdr, frames := splitV2(t, wire)
	frames[0][len(frames[0])-1] ^= 0xFF
	wire = joinV2(hdr, frames...)
	dec, err := NewDecryptor(bytes.NewReader(wire), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	assertHardFail(t, dec, make([]byte, len(plain)), nil)
}

func TestV2Attack_AlteredCiphertext(t *testing.T) {
	key := v2FixedKey(t)
	plain := bytes.Repeat([]byte{0x71}, 80)
	wire := v2Seal(t, key, [][]byte{plain}, nil)
	hdr, frames := splitV2(t, wire)
	frames[0][6] ^= 0xFF
	tampered := joinV2(hdr, frames...)
	dec, err := NewDecryptor(bytes.NewReader(tampered), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	p := bytes.Repeat([]byte{0xA5}, len(plain))
	n, err := dec.Read(p)
	if err == nil || err == io.EOF {
		t.Fatalf("chiffré altéré accepté: n=%d err=%v", n, err)
	}
	if n != 0 {
		t.Fatalf("échec a déclaré %d octets", n)
	}
	if bytes.Equal(p, plain) {
		t.Fatal("clair rendu malgré MAC invalide")
	}
	assertHardFail(t, dec, make([]byte, len(plain)), nil)
}

func TestV2Attack_DataAfterTagFinal(t *testing.T) {
	key := v2FixedKey(t)
	plain := bytes.Repeat([]byte{0x81}, 80)
	wire := v2Seal(t, key, [][]byte{plain}, nil)
	hdr, frames := splitV2(t, wire)
	withExtra := joinV2(hdr, frames...)
	withExtra = append(withExtra, 0x00)
	dec, err := NewDecryptor(bytes.NewReader(withExtra), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	got := make([]byte, len(plain))
	if n, err := dec.Read(got); err != nil || n != len(plain) {
		t.Fatalf("message: n=%d err=%v", n, err)
	}
	assertHardFail(t, dec, make([]byte, 8), nil)

	secondFinal := joinV2(hdr, frames[0], frames[len(frames)-1], frames[len(frames)-1])
	dec2, err := NewDecryptor(bytes.NewReader(secondFinal), key)
	if err != nil {
		t.Fatalf("NewDecryptor second final: %v", err)
	}
	if n, err := dec2.Read(got); err != nil || n != len(plain) {
		t.Fatalf("message avant second final: n=%d err=%v", n, err)
	}
	assertHardFail(t, dec2, make([]byte, 8), nil)
}

func TestV2Attack_HeaderVersionFlagsMagic(t *testing.T) {
	key := v2FixedKey(t)
	plain := bytes.Repeat([]byte{0x91}, 40)
	wire := v2Seal(t, key, [][]byte{plain}, nil)

	t.Run("version3", func(t *testing.T) {
		w := append([]byte(nil), wire...)
		binary.BigEndian.PutUint16(w[8:10], 3)
		_, err := NewDecryptor(bytes.NewReader(w), key)
		if err == nil {
			t.Fatal("version 3 acceptée")
		}
	})
	t.Run("flags_nonzero", func(t *testing.T) {
		w := append([]byte(nil), wire...)
		binary.BigEndian.PutUint16(w[10:12], 1)
		_, err := NewDecryptor(bytes.NewReader(w), key)
		if err == nil {
			t.Fatal("flags ≠ 0 acceptés")
		}
	})
	t.Run("magic_altered", func(t *testing.T) {
		w := append([]byte(nil), wire...)
		w[0] ^= 0x01
		dec, err := NewDecryptor(bytes.NewReader(w), key)
		if err != nil {
			t.Fatalf("magic altéré doit prendre le chemin v1, pas échouer à l'en-tête: %v", err)
		}
		assertHardFail(t, dec, make([]byte, len(plain)), nil)
	})
}

func TestV2_CloseIdempotent(t *testing.T) {
	key := v2FixedKey(t)
	var buf bytes.Buffer
	enc, err := NewEncryptor(&buf, key)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := enc.Write([]byte("x")); err != nil {
		t.Fatal(err)
	}
	if err := enc.Close(); err != nil {
		t.Fatal(err)
	}
	afterFirst := append([]byte(nil), buf.Bytes()...)
	if err := enc.Close(); err != nil {
		t.Fatalf("second Close: %v", err)
	}
	if !bytes.Equal(buf.Bytes(), afterFirst) {
		t.Fatal("second Close a réémis un bloc terminal")
	}
	_, frames := splitV2(t, afterFirst)
	final := frames[len(frames)-1]
	if final[4] != TagFinal {
		t.Fatalf("dernier tag = 0x%02x, attendu TagFinal", final[4])
	}
	n := int(binary.BigEndian.Uint32(final[:4]))
	if n != MinFramePayloadV2 {
		t.Fatalf("longueur TagFinal = %d, attendu %d", n, MinFramePayloadV2)
	}
}

func TestV2_ReadAfterEOFStable(t *testing.T) {
	key := v2FixedKey(t)
	plain := []byte("fin-stable")
	wire := v2Seal(t, key, [][]byte{plain}, nil)
	dec, err := NewDecryptor(bytes.NewReader(wire), key)
	if err != nil {
		t.Fatal(err)
	}
	got := make([]byte, len(plain))
	if n, err := dec.Read(got); err != nil || n != len(plain) {
		t.Fatalf("message: n=%d err=%v", n, err)
	}
	for i := 0; i < 3; i++ {
		n, err := dec.Read(got)
		if n != 0 || err != io.EOF {
			t.Fatalf("lecture %d après EOF: n=%d err=%v", i, n, err)
		}
	}
}

func TestV2_HotPathAllocsUnchanged(t *testing.T) {
	key := v2FixedKey(t)
	chunk := bytes.Repeat([]byte{0x5A}, 64)
	var buf bytes.Buffer
	enc, err := NewEncryptor(&buf, key)
	if err != nil {
		t.Fatal(err)
	}
	const warmup = 32
	for i := 0; i < warmup; i++ {
		if _, err := enc.Write(chunk); err != nil {
			t.Fatal(err)
		}
	}
	writeAllocs := testing.AllocsPerRun(50, func() {
		if _, err := enc.Write(chunk); err != nil {
			t.Fatalf("Write: %v", err)
		}
	})
	if err := enc.Close(); err != nil {
		t.Fatal(err)
	}

	dec, err := NewDecryptor(bytes.NewReader(buf.Bytes()), key)
	if err != nil {
		t.Fatal(err)
	}
	p := make([]byte, 64)
	for i := 0; i < warmup; i++ {
		if n, err := dec.Read(p); err != nil || n != 64 {
			t.Fatalf("warmup Read: n=%d err=%v", n, err)
		}
	}
	readAllocs := testing.AllocsPerRun(50, func() {
		n, err := dec.Read(p)
		if err != nil || n != 64 {
			t.Fatalf("Read: n=%d err=%v", n, err)
		}
	})
	t.Logf("Write %v allocs/op, Read %v allocs/op", writeAllocs, readAllocs)
	if writeAllocs > 2 {
		t.Fatalf("Write chemin chaud: %v allocs/op, plafond historique 2", writeAllocs)
	}
	if readAllocs > 2 {
		t.Fatalf("Read chemin chaud: %v allocs/op, plafond historique 2", readAllocs)
	}
}
