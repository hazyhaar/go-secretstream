// SPDX-License-Identifier: Apache-2.0 OR MIT

package formatgen

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hazyhaar/go-secretstream/internal/engine"
)

type cueRoot struct {
	FormatV2 formatV2 `json:"format_v2"`
}

type formatV2 struct {
	Endianness string     `json:"endianness"`
	Header     headerSpec `json:"header"`
	Frame      frameSpec  `json:"frame"`
	Tags       tagsSpec   `json:"tags"`
	AD         adSpec     `json:"ad"`
	ChunkSize  int        `json:"chunk_size"`
	KeySize    int        `json:"key_size"`
	FrameNonce frameNonce `json:"frame_nonce"`
	Hybrid     hybridSpec `json:"hybrid"`
	Vectors    []vectorIn `json:"vectors"`
}

type headerSpec struct {
	Size   int           `json:"size"`
	SizeV1 int           `json:"size_v1"`
	Fields []headerField `json:"fields"`
}

type headerField struct {
	Name     string `json:"name"`
	Offset   int    `json:"offset"`
	Size     int    `json:"size"`
	Kind     string `json:"kind"`
	Value    *int   `json:"value"`
	ValueHex string `json:"value_hex"`
}

type frameSpec struct {
	LengthSize int `json:"length_size"`
	TagSize    int `json:"tag_size"`
	MacSize    int `json:"mac_size"`
	MinPayload int `json:"min_payload"`
	MaxPayload int `json:"max_payload"`
}

type tagsSpec struct {
	Message  int   `json:"message"`
	Push     int   `json:"push"`
	Rekey    int   `json:"rekey"`
	Final    int   `json:"final"`
	Admitted []int `json:"admitted"`
}

type adSpec struct {
	DomainSize int `json:"domain_size"`
	SeqSize    int `json:"seq_size"`
	TagSize    int `json:"tag_size"`
	LenSize    int `json:"len_size"`
	PrefixLen  int `json:"prefix_len"`
}

type frameNonce struct {
	PrefixFromNonceOffset int `json:"prefix_from_nonce_offset"`
	PrefixSize            int `json:"prefix_size"`
	SeqSize               int `json:"seq_size"`
	IETFSize              int `json:"ietf_size"`
}

type hybridSpec struct {
	ProbeSize            int      `json:"probe_size"`
	CollisionProbability string   `json:"collision_probability"`
	Rules                []string `json:"rules"`
}

type vectorIn struct {
	Name         string   `json:"name"`
	KeyHex       string   `json:"key_hex"`
	NonceHex     string   `json:"nonce_hex"`
	FragmentsHex []string `json:"fragments_hex"`
	AdsHex       []string `json:"ads_hex"`
}

type vectorsFile struct {
	Engine  string      `json:"engine"`
	Vectors []vectorOut `json:"vectors"`
}

type vectorOut struct {
	Name         string   `json:"name"`
	KeyHex       string   `json:"key_hex"`
	NonceHex     string   `json:"nonce_hex"`
	FragmentsHex []string `json:"fragments_hex"`
	AdsHex       []string `json:"ads_hex"`
	HeaderHex    string   `json:"header_hex"`
	FramesHex    []string `json:"frames_hex"`
}

func Generate(specDir, outDir string) error {
	specAbs, err := filepath.Abs(specDir)
	if err != nil {
		return fmt.Errorf("spec: %w", err)
	}
	outAbs, err := filepath.Abs(outDir)
	if err != nil {
		return fmt.Errorf("out: %w", err)
	}
	cmd := exec.Command("cue", "export", "--out", "json", specAbs)
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("cue export: %w: %s", err, bytes.TrimSpace(ee.Stderr))
		}
		return fmt.Errorf("cue export: %w", err)
	}
	var root cueRoot
	if err := json.Unmarshal(out, &root); err != nil {
		return fmt.Errorf("cue json: %w", err)
	}
	spec := root.FormatV2
	goSrc, err := emitGo(spec)
	if err != nil {
		return err
	}
	md, err := emitMD(spec)
	if err != nil {
		return err
	}
	vec, err := emitVectors(spec)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(outAbs, "format_v2_gen.go"), goSrc, 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(outAbs, "FORMAT_V2.md"), md, 0o644); err != nil {
		return err
	}
	vecDir := filepath.Join(outAbs, "testdata", "v2")
	if err := os.MkdirAll(vecDir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(vecDir, "vectors.json"), vec, 0o644)
}

func fieldByName(fields []headerField, name string) (headerField, error) {
	for _, f := range fields {
		if f.Name == name {
			return f, nil
		}
	}
	return headerField{}, fmt.Errorf("champ d'en-tête %q absent du descripteur", name)
}

func emitGo(spec formatV2) ([]byte, error) {
	magicF, err := fieldByName(spec.Header.Fields, "magic")
	if err != nil {
		return nil, err
	}
	verF, err := fieldByName(spec.Header.Fields, "version")
	if err != nil {
		return nil, err
	}
	flagsF, err := fieldByName(spec.Header.Fields, "flags")
	if err != nil {
		return nil, err
	}
	if verF.Value == nil || flagsF.Value == nil {
		return nil, fmt.Errorf("version ou flags sans value")
	}
	magic, err := hex.DecodeString(magicF.ValueHex)
	if err != nil {
		return nil, fmt.Errorf("magic: %w", err)
	}
	if len(magic) != magicF.Size {
		return nil, fmt.Errorf("magic: %d octets, attendu %d", len(magic), magicF.Size)
	}
	var magBuf strings.Builder
	for i, b := range magic {
		if i > 0 {
			magBuf.WriteString(", ")
		}
		fmt.Fprintf(&magBuf, "0x%02x", b)
	}
	var buf bytes.Buffer
	buf.WriteString("// Code generated by formatgen. DO NOT EDIT.\n\n")
	buf.WriteString("package secretstream55\n\n")
	buf.WriteString("const (\n")
	fmt.Fprintf(&buf, "\t// HeaderSizeV1 est la taille de l'en-tête maison v1 (nonce seul) et\n")
	fmt.Fprintf(&buf, "\t// de l'en-tête libsodium.\n")
	fmt.Fprintf(&buf, "\tHeaderSizeV1 = %d\n", spec.Header.SizeV1)
	fmt.Fprintf(&buf, "\t// HeaderSize est la taille de l'en-tête maison v2 émis par NewEncryptor.\n")
	fmt.Fprintf(&buf, "\tHeaderSize = %d\n", spec.Header.Size)
	fmt.Fprintf(&buf, "\t// HeaderSizeV2 est un alias de HeaderSize (en-tête v2).\n")
	fmt.Fprintf(&buf, "\tHeaderSizeV2 = HeaderSize\n")
	fmt.Fprintf(&buf, "\tTagSize = %d\n", spec.Frame.MacSize)
	fmt.Fprintf(&buf, "\tChunkSize = %d\n", spec.ChunkSize)
	fmt.Fprintf(&buf, "\tTagMessage byte = 0x%02x\n", spec.Tags.Message)
	fmt.Fprintf(&buf, "\tTagPush byte = 0x%02x\n", spec.Tags.Push)
	fmt.Fprintf(&buf, "\tTagRekey byte = 0x%02x\n", spec.Tags.Rekey)
	fmt.Fprintf(&buf, "\tTagFinal byte = 0x%02x\n", spec.Tags.Final)
	fmt.Fprintf(&buf, "\tVersionV2 uint16 = %d\n", *verF.Value)
	fmt.Fprintf(&buf, "\tFlagsV2 uint16 = %d\n", *flagsF.Value)
	fmt.Fprintf(&buf, "\tMinFramePayloadV2 = %d\n", spec.Frame.MinPayload)
	fmt.Fprintf(&buf, "\tMaxFramePayloadV2 = %d\n", spec.Frame.MaxPayload)
	fmt.Fprintf(&buf, "\tadPrefixV2Len = %d\n", spec.AD.PrefixLen)
	buf.WriteString(")\n\n")
	fmt.Fprintf(&buf, "var magicV2 = [%d]byte{%s}\n", len(magic), magBuf.String())
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("gofmt: %w\n%s", err, buf.Bytes())
	}
	return formatted, nil
}

func emitMD(spec formatV2) ([]byte, error) {
	magicF, err := fieldByName(spec.Header.Fields, "magic")
	if err != nil {
		return nil, err
	}
	verF, err := fieldByName(spec.Header.Fields, "version")
	if err != nil {
		return nil, err
	}
	flagsF, err := fieldByName(spec.Header.Fields, "flags")
	if err != nil {
		return nil, err
	}
	nonceF, err := fieldByName(spec.Header.Fields, "nonce")
	if err != nil {
		return nil, err
	}
	if verF.Value == nil || flagsF.Value == nil {
		return nil, fmt.Errorf("version ou flags sans value")
	}
	var b strings.Builder
	b.WriteString("# Format de flux maison v2\n\n")
	b.WriteString("Le descripteur déclaratif `spec/format_v2.cue` gouverne ce format. ")
	b.WriteString("Les constantes Go, la présente spécification et les vecteurs de test en sont des produits générés. ")
	b.WriteString("Le code manuscrit ne redéfinit pas ces valeurs.\n\n")
	b.WriteString("L'ordre des octets du fil est grand-boutiste.\n\n")
	b.WriteString("## En-tête\n\n")
	fmt.Fprintf(&b, "L'en-tête v2 occupe %d octets. L'en-tête v1 (nonce seul, également taille libsodium) occupe %d octets.\n\n", spec.Header.Size, spec.Header.SizeV1)
	fmt.Fprintf(&b, "Le magique tient sur %d octets à l'offset %d (hexadécimal %s, ASCII « SS55-v2 » suivi d'un octet nul). ", magicF.Size, magicF.Offset, magicF.ValueHex)
	fmt.Fprintf(&b, "Le numéro de version occupe %d octets à l'offset %d et vaut %d. ", verF.Size, verF.Offset, *verF.Value)
	fmt.Fprintf(&b, "Le champ de drapeaux occupe %d octets à l'offset %d et doit valoir %d. ", flagsF.Size, flagsF.Offset, *flagsF.Value)
	fmt.Fprintf(&b, "Le nonce aléatoire occupe %d octets à l'offset %d.\n\n", nonceF.Size, nonceF.Offset)
	b.WriteString("La sous-clé se dérive par HChaCha20 à partir de la clé et des seize premiers octets du nonce.\n\n")
	b.WriteString("## Trame\n\n")
	fmt.Fprintf(&b, "Chaque trame commence par une longueur de %d octets grand-boutiste, qui compte le tag, le chiffré et le MAC. ", spec.Frame.LengthSize)
	fmt.Fprintf(&b, "Le tag tient sur %d octet. Le MAC tient sur %d octets. ", spec.Frame.TagSize, spec.Frame.MacSize)
	fmt.Fprintf(&b, "La charge utile minimale vaut %d octets (tag et MAC, chiffré vide). ", spec.Frame.MinPayload)
	fmt.Fprintf(&b, "La charge utile maximale vaut %d octets (tag, fragment d'au plus %d octets, MAC).\n\n", spec.Frame.MaxPayload, spec.ChunkSize)
	fmt.Fprintf(&b, "Le tag de message vaut 0x%02x. Le tag terminal vaut 0x%02x. ", spec.Tags.Message, spec.Tags.Final)
	fmt.Fprintf(&b, "Les valeurs 0x%02x (push) et 0x%02x (rekey) sont refusées. ", spec.Tags.Push, spec.Tags.Rekey)
	b.WriteString("Le tag entre dans la donnée associée authentifiée : le modifier sur le fil fait échouer le MAC.\n\n")
	fmt.Fprintf(&b, "Le nonce IETF de chaque trame (%d octets) est formé des %d octets du nonce d'en-tête à l'offset %d, suivis du compteur de séquence grand-boutiste de %d octets, en clair et sans masque XOR. ", spec.FrameNonce.IETFSize, spec.FrameNonce.PrefixSize, spec.FrameNonce.PrefixFromNonceOffset, spec.FrameNonce.SeqSize)
	b.WriteString("Le compteur part de zéro et s'incrémente à chaque trame. Un débordement est une erreur collante.\n\n")
	b.WriteString("## Donnée associée\n\n")
	fmt.Fprintf(&b, "La donnée associée authentifiée est le magique (%d octets), le compteur de séquence (%d octets), le tag (%d octet), la longueur de l'ad d'appelant (%d octets grand-boutiste), puis l'ad d'appelant. ", spec.AD.DomainSize, spec.AD.SeqSize, spec.AD.TagSize, spec.AD.LenSize)
	fmt.Fprintf(&b, "Le préfixe sans ad d'appelant tient sur %d octets. ", spec.AD.PrefixLen)
	b.WriteString("Le préfixe de longueur sépare les concaténations qui se confondaient en v1. L'ad d'appelant n'est pas transmise sur le fil : le lecteur la fournit.\n\n")
	b.WriteString("## Clôture et lecture\n\n")
	fmt.Fprintf(&b, "Close écrit une trame terminale à chiffré vide (longueur %d), puis efface les secrets. ", spec.Frame.MinPayload)
	b.WriteString("Un second Close n'émet rien. Write après Close reste une erreur. ")
	b.WriteString("Le lecteur ne rend io.EOF qu'après avoir authentifié cette trame terminale. ")
	b.WriteString("Une fin de flux avant ce terminal est une erreur collante de troncature, jamais un io.EOF. ")
	b.WriteString("Toute donnée après le terminal, une seconde trame terminale, ou un terminal à chiffré non vide, est une erreur.\n\n")
	b.WriteString("## Décodeur hybride\n\n")
	fmt.Fprintf(&b, "Le décodeur lit d'abord %d octets. ", spec.Hybrid.ProbeSize)
	for _, r := range spec.Hybrid.Rules {
		b.WriteString(r)
		b.WriteString(" ")
	}
	fmt.Fprintf(&b, "La probabilité qu'un nonce v1 aléatoire commence par le magique est %s.\n", spec.Hybrid.CollisionProbability)
	return []byte(b.String()), nil
}

func emitVectors(spec formatV2) ([]byte, error) {
	eng := engine.Default()
	magicF, err := fieldByName(spec.Header.Fields, "magic")
	if err != nil {
		return nil, err
	}
	verF, err := fieldByName(spec.Header.Fields, "version")
	if err != nil {
		return nil, err
	}
	flagsF, err := fieldByName(spec.Header.Fields, "flags")
	if err != nil {
		return nil, err
	}
	if verF.Value == nil || flagsF.Value == nil {
		return nil, fmt.Errorf("version ou flags sans value")
	}
	magic, err := hex.DecodeString(magicF.ValueHex)
	if err != nil {
		return nil, err
	}
	out := vectorsFile{Engine: "default"}
	for _, in := range spec.Vectors {
		key, err := hex.DecodeString(in.KeyHex)
		if err != nil {
			return nil, fmt.Errorf("%s key: %w", in.Name, err)
		}
		nonce, err := hex.DecodeString(in.NonceHex)
		if err != nil {
			return nil, fmt.Errorf("%s nonce: %w", in.Name, err)
		}
		if len(key) != spec.KeySize || len(nonce) != spec.Header.Size-magicF.Size-verF.Size-flagsF.Size {
			return nil, fmt.Errorf("%s: taille de clé ou de nonce hors descripteur", in.Name)
		}
		frags := make([][]byte, len(in.FragmentsHex))
		ads := make([][]byte, len(in.AdsHex))
		for i, h := range in.FragmentsHex {
			frags[i], err = hex.DecodeString(h)
			if err != nil {
				return nil, fmt.Errorf("%s fragment %d: %w", in.Name, i, err)
			}
		}
		for i, h := range in.AdsHex {
			ads[i], err = hex.DecodeString(h)
			if err != nil {
				return nil, fmt.Errorf("%s ad %d: %w", in.Name, i, err)
			}
		}
		hdr, frames, err := sealV2(eng, spec, magic, uint16(*verF.Value), uint16(*flagsF.Value), key, nonce, frags, ads)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", in.Name, err)
		}
		vo := vectorOut{
			Name:         in.Name,
			KeyHex:       in.KeyHex,
			NonceHex:     in.NonceHex,
			FragmentsHex: in.FragmentsHex,
			AdsHex:       in.AdsHex,
			HeaderHex:    hex.EncodeToString(hdr),
			FramesHex:    make([]string, len(frames)),
		}
		if vo.FragmentsHex == nil {
			vo.FragmentsHex = []string{}
		}
		if vo.AdsHex == nil {
			vo.AdsHex = []string{}
		}
		for i, f := range frames {
			vo.FramesHex[i] = hex.EncodeToString(f)
		}
		out.Vectors = append(out.Vectors, vo)
	}
	raw, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return nil, err
	}
	raw = append(raw, '\n')
	return raw, nil
}

func sealV2(eng engine.AEAD, spec formatV2, magic []byte, version, flags uint16, key, nonce []byte, frags, ads [][]byte) (header []byte, frames [][]byte, err error) {
	header = make([]byte, spec.Header.Size)
	copy(header[0:8], magic)
	binary.BigEndian.PutUint16(header[8:10], version)
	binary.BigEndian.PutUint16(header[10:12], flags)
	copy(header[12:36], nonce)
	var subkey [32]byte
	eng.HChaCha20(subkey[:], key, nonce[0:16])
	seq := uint64(0)
	seal := func(payload, ad []byte, tag byte) error {
		n12 := make([]byte, spec.FrameNonce.IETFSize)
		copy(n12[0:spec.FrameNonce.PrefixSize], nonce[spec.FrameNonce.PrefixFromNonceOffset:spec.FrameNonce.PrefixFromNonceOffset+spec.FrameNonce.PrefixSize])
		binary.BigEndian.PutUint64(n12[spec.FrameNonce.PrefixSize:], seq)
		prefix := spec.AD.PrefixLen
		adBuf := make([]byte, prefix+len(ad))
		copy(adBuf[:len(magic)], magic)
		binary.BigEndian.PutUint64(adBuf[spec.AD.DomainSize:spec.AD.DomainSize+spec.AD.SeqSize], seq)
		adBuf[spec.AD.DomainSize+spec.AD.SeqSize] = tag
		lenOff := spec.AD.DomainSize + spec.AD.SeqSize + spec.AD.TagSize
		binary.BigEndian.PutUint32(adBuf[lenOff:lenOff+spec.AD.LenSize], uint32(len(ad)))
		copy(adBuf[prefix:], ad)
		bound := adBuf[:prefix]
		if len(ad) > 0 {
			bound = adBuf
		}
		cipher := make([]byte, len(payload))
		var mac [16]byte
		if err := eng.LockSubkeyDst(cipher, &mac, subkey[:], n12, bound, payload); err != nil {
			return err
		}
		total := spec.Frame.TagSize + len(payload) + spec.Frame.MacSize
		frame := make([]byte, spec.Frame.LengthSize+total)
		binary.BigEndian.PutUint32(frame[0:spec.Frame.LengthSize], uint32(total))
		frame[spec.Frame.LengthSize] = tag
		copy(frame[spec.Frame.LengthSize+spec.Frame.TagSize:], cipher)
		copy(frame[spec.Frame.LengthSize+spec.Frame.TagSize+len(payload):], mac[:])
		frames = append(frames, frame)
		seq++
		return nil
	}
	for i, f := range frags {
		var ad []byte
		if i < len(ads) {
			ad = ads[i]
		}
		remain := f
		for len(remain) > 0 {
			n := len(remain)
			if n > spec.ChunkSize {
				n = spec.ChunkSize
			}
			if err := seal(remain[:n], ad, byte(spec.Tags.Message)); err != nil {
				return nil, nil, err
			}
			remain = remain[n:]
		}
	}
	if err := seal(nil, nil, byte(spec.Tags.Final)); err != nil {
		return nil, nil, err
	}
	return header, frames, nil
}
