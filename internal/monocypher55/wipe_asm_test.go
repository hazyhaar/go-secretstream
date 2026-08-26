// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

var (
	reStore   = regexp.MustCompile(`(?i)^((?:V)?MOV\S*)\s+(.+),\s*(.+)$`)
	reDispReg = regexp.MustCompile(`^((?:0x)?[0-9A-Fa-f]+)?\(([A-Z0-9]+)\)$`)
	reLEA     = regexp.MustCompile(`(?i)^LEA[QL]?\s+((?:0x)?[0-9A-Fa-f]+)?\(SP\),\s*([A-Z0-9]+)$`)
)

var keepWipeSymbols = []any{
	Crypto_wipe,
	Crypto_eddsa_to_x25519,
	Crypto_x25519_to_eddsa,
	Crypto_chacha20_x,
}

type wipeCase struct {
	name string
	re   string
	want int
}

func TestWipeSurvivesInAsm(t *testing.T) {
	requireAmd64(t)
	if keepWipeSymbols == nil {
		t.Fatal("keepWipeSymbols élidé")
	}
	bin := compileTestBinary(t)

	cases := []wipeCase{
		{"Crypto_wipe", `\.Crypto_wipe$`, 1},
		{"Crypto_eddsa_reduce", `\.Crypto_eddsa_reduce$`, int(unsafe.Sizeof([16]uint32{}))},
		{"Crypto_eddsa_mul_add", `\.Crypto_eddsa_mul_add$`, int(unsafe.Sizeof([16]uint32{}))},
		{"Crypto_eddsa_key_pair", `\.Crypto_eddsa_key_pair$`, int(unsafe.Sizeof([64]uint8{}))},
		{"Crypto_eddsa_sign", `\.Crypto_eddsa_sign$`, int(unsafe.Sizeof([64]uint8{}))},
		{"Crypto_eddsa_to_x25519", `\.Crypto_eddsa_to_x25519$`, int(unsafe.Sizeof([10]int32{}))},
		{"Crypto_eddsa_scalarbase", `\.Crypto_eddsa_scalarbase$`, int(unsafe.Sizeof(Ge{}))},
		{"Crypto_x25519", `\.Crypto_x25519$`, int(unsafe.Sizeof([32]uint8{}))},
		{"Crypto_x25519_to_eddsa", `\.Crypto_x25519_to_eddsa$`, int(unsafe.Sizeof([10]int32{}))},
		{"Crypto_x25519_dirty_fast", `\.Crypto_x25519_dirty_fast$`, int(unsafe.Sizeof([32]uint8{}))},
		{"redc", `\.redc$`, int(unsafe.Sizeof([16]uint32{}))},
		{"Crypto_x25519_inverse", `\.Crypto_x25519_inverse$`, int(unsafe.Sizeof([16]uint32{}))},
		{"fe_tobytes", `\.fe_tobytes$`, int(unsafe.Sizeof([10]int{}))},
		{"fe_isodd", `\.fe_isodd$`, int(unsafe.Sizeof([32]uint8{}))},
		{"fe_isequal", `\.fe_isequal$`, int(unsafe.Sizeof([32]uint8{}))},
		{"fe_invert", `\.fe_invert$`, int(unsafe.Sizeof([10]int32{}))},
		{"ge_tobytes", `\.ge_tobytes$`, int(unsafe.Sizeof([10]int32{}))},
		{"select_lop", `\.select_lop$`, int(unsafe.Sizeof([10]int32{}))},
		{"Crypto_elligator_map", `\.Crypto_elligator_map$`, int(unsafe.Sizeof([10]int32{}))},
		{"Crypto_elligator_rev", `\.Crypto_elligator_rev$`, int(unsafe.Sizeof([10]int32{}))},
		{"Crypto_chacha20_x", `\.Crypto_chacha20_x$`, int(unsafe.Sizeof([32]uint8{}))},
		{"blake_update_32", `\.blake_update_32$`, int(unsafe.Sizeof([4]uint8{}))},
		{"Crypto_poly1305_final", `\.Crypto_poly1305_final$`, 1},
		{"Crypto_blake2b_final", `\.Crypto_blake2b_final$`, 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			listing := objdumpSymbol(t, bin, tc.re)
			if err := attestEmittedWipe(t, bin, listing, tc.want); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestWipeGuard_NegativeControl(t *testing.T) {
	requireAmd64(t)
	want := int(unsafe.Sizeof([32]uint8{}))
	bin := compileTestBinary(t)
	wipeWitnessNoKeepAlive()

	t.Run("temoin_sans_keepalive", func(t *testing.T) {
		listing := objdumpSymbol(t, bin, `wipeWitnessNoKeepAlive$`)
		if err := attestStackWipeAfterLive(listing, want); err == nil {
			t.Fatal("le témoin conserve un effacement de pile malgré l'absence de KeepAlive ; le négatif de pile est absent")
		}
		if hasCall(listing, "Crypto_wipe") || hasMemclr(listing) {
			t.Fatal("le témoin appelle encore Crypto_wipe ou memclr ; le négatif est absent")
		}
	})
}

//go:noinline
func wipeWitnessNoKeepAlive() {
	var buf [32]uint8
	buf[0] = 0xaa
	buf[15] = 0xbb
	buf[31] = 0xcc
	buf = [32]uint8{}
}

func attestEmittedWipe(t *testing.T, bin, listing string, minBytes int) error {
	t.Helper()
	if hasCall(listing, "Crypto_wipe") {
		body := objdumpSymbol(t, bin, `monocypher55\.Crypto_wipe$`)
		if !wipeBodyHasZeros(body) && !hasMemclr(body) {
			return fmt.Errorf("CALL Crypto_wipe présent mais le corps n'a ni stockage nul ni memclr")
		}
		return nil
	}
	if hasMemclr(listing) {
		return nil
	}
	if minBytes <= 1 && wipeBodyHasZeros(listing) {
		return nil
	}
	return attestStackWipeAfterLive(listing, minBytes)
}

func wipeBodyHasZeros(listing string) bool {
	lines := asmLines(listing)
	for _, reg := range []string{"AX", "BX", "CX", "DX", "SI", "DI", "BP", "SP", "R8", "R9", "R10", "R11", "R12", "R13", "R14", "R15"} {
		if maxRun(zeroStores(lines, reg)) > 0 {
			return true
		}
	}
	return false
}

func requireAmd64(t *testing.T) {
	t.Helper()
	if runtime.GOARCH != "amd64" {
		t.Fatalf("garde de désassemblage limitée à amd64, GOARCH=%s", runtime.GOARCH)
	}
}

func compileTestBinary(t *testing.T) string {
	t.Helper()
	out := filepath.Join(t.TempDir(), "x.test")
	cmd := exec.Command("go", "test", "-c", "-o", out, ".")
	cmd.Env = probeEnv()
	if b, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("go test -c : %v\n%s", err, b)
	}
	st, err := os.Stat(out)
	if err != nil || st.Size() == 0 {
		t.Fatalf("binaire de test introuvable ou vide : %s", out)
	}
	return out
}

func objdumpSymbol(t *testing.T, bin, re string) string {
	t.Helper()
	script := wipeProbeScript(t)
	cmd := exec.Command(script, bin, re)
	cmd.Env = probeEnv()
	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("symbole %q introuvable : %v\n%s", re, err, b)
	}
	s := string(b)
	if !strings.Contains(s, "TEXT ") {
		t.Fatalf("symbole %q : dump sans TEXT\n%s", re, s)
	}
	return s
}

func wipeProbeScript(t *testing.T) string {
	t.Helper()
	for _, p := range []string{
		filepath.Join("scripts", "wipe_probe.sh"),
		filepath.Join("..", "..", "c2simd", "scripts", "wipe_probe.sh"),
	} {
		if st, err := os.Stat(p); err == nil && !st.IsDir() {
			abs, err := filepath.Abs(p)
			if err != nil {
				t.Fatal(err)
			}
			return abs
		}
	}
	t.Fatal("scripts/wipe_probe.sh introuvable")
	return ""
}

func probeEnv() []string {
	out := make([]string, 0, 16)
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, "GOFLAGS=") {
			v := strings.ReplaceAll(strings.TrimPrefix(e, "GOFLAGS="), "-race", "")
			out = append(out, "GOFLAGS="+strings.TrimSpace(v))
			continue
		}
		if strings.HasPrefix(e, "GOTOOLCHAIN=") || strings.HasPrefix(e, "GOEXPERIMENT=") {
			continue
		}
		// GOWORK est hérité tel quel : la recette du poste le pose, un conteneur
		// vierge n'en a pas (le paquet se résout seul). Aucun chemin hôte ici.
		out = append(out, e)
	}
	return append(out,
		"GOTOOLCHAIN=go1.27.0",
		"GOEXPERIMENT=simd",
	)
}

func hasCall(listing, needle string) bool {
	for _, line := range strings.Split(listing, "\n") {
		instr := asmInstruction(line)
		if strings.Contains(instr, "CALL") && strings.Contains(instr, needle) {
			return true
		}
	}
	return false
}

func hasMemclr(listing string) bool {
	return hasCall(listing, "memclrNoHeapPointers") || hasCall(listing, "memclrHasPointers") || hasCall(listing, "runtime.memclr")
}

func attestStackWipeAfterLive(listing string, minBytes int) error {
	lines := happyPath(asmLines(listing))
	alias := map[string]int{"SP": 0}
	lastLive := -1
	for i, instr := range lines {
		noteStackAlias(alias, instr)
		_, src, dest, ok := parseStore(instr)
		if !ok || !strings.Contains(dest, "(") {
			continue
		}
		if destIsStack(dest, alias) {
			continue
		}
		if !isZeroSrc(src) {
			lastLive = i
		}
	}
	var rest []string
	if lastLive >= 0 {
		rest = lines[lastLive+1:]
	} else {
		rest = lines
	}
	stores := zeroStoresStackFrom(rest, aliasAt(lines, lastLive))
	if maxRun(stores) < minBytes {
		return fmt.Errorf("aucune plage nulle de %d octets vers la pile après le dernier stockage hors pile (stores=%v)", minBytes, stores)
	}
	return nil
}

func happyPath(lines []string) []string {
	for i, instr := range lines {
		u := strings.ToUpper(strings.TrimSpace(instr))
		if u == "RET" || strings.HasPrefix(u, "RET ") {
			return lines[:i+1]
		}
	}
	return lines
}

func aliasAt(lines []string, lastLive int) map[string]int {
	alias := map[string]int{"SP": 0}
	for i, instr := range lines {
		noteStackAlias(alias, instr)
		if lastLive >= 0 && i >= lastLive {
			break
		}
	}
	return alias
}

func zeroStoresStackFrom(lines []string, alias map[string]int) []span {
	if alias == nil {
		alias = map[string]int{"SP": 0}
	}
	var out []span
	for _, instr := range lines {
		noteStackAlias(alias, instr)
		op, src, dest, ok := parseStore(instr)
		if !ok || !isZeroSrc(src) {
			continue
		}
		m := reDispReg.FindStringSubmatch(dest)
		if m == nil {
			continue
		}
		base, ok := alias[m[2]]
		if !ok {
			continue
		}
		off, err := parseDisp(m[1])
		if err != nil {
			continue
		}
		w := storeWidth(op, src)
		if w <= 0 {
			continue
		}
		out = append(out, span{base + off, base + off + w})
	}
	return mergeSpans(out)
}

func noteStackAlias(alias map[string]int, instr string) {
	m := reLEA.FindStringSubmatch(instr)
	if m == nil {
		return
	}
	off, err := parseDisp(m[1])
	if err != nil {
		return
	}
	alias[m[2]] = off
}

func destIsStack(dest string, alias map[string]int) bool {
	m := reDispReg.FindStringSubmatch(dest)
	if m == nil {
		return false
	}
	_, ok := alias[m[2]]
	return ok
}

func zeroStoresStack(lines []string) []span {
	alias := map[string]int{"SP": 0}
	var out []span
	for _, instr := range lines {
		noteStackAlias(alias, instr)
		op, src, dest, ok := parseStore(instr)
		if !ok || !isZeroSrc(src) {
			continue
		}
		m := reDispReg.FindStringSubmatch(dest)
		if m == nil {
			continue
		}
		base, ok := alias[m[2]]
		if !ok {
			continue
		}
		off, err := parseDisp(m[1])
		if err != nil {
			continue
		}
		w := storeWidth(op, src)
		if w <= 0 {
			continue
		}
		out = append(out, span{base + off, base + off + w})
	}
	return mergeSpans(out)
}

func asmLines(listing string) []string {
	var out []string
	for _, line := range strings.Split(listing, "\n") {
		if instr := asmInstruction(line); instr != "" {
			out = append(out, instr)
		}
	}
	return out
}

func asmInstruction(line string) string {
	var last string
	for _, p := range strings.Split(line, "\t") {
		p = strings.TrimSpace(p)
		if p != "" {
			last = p
		}
	}
	if last == "" || strings.HasPrefix(last, "TEXT ") {
		return ""
	}
	if isHexBlob(last) {
		return ""
	}
	return last
}

func isHexBlob(s string) bool {
	if len(s) < 4 || len(s)%2 != 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

func parseStore(instr string) (op, src, dest string, ok bool) {
	m := reStore.FindStringSubmatch(instr)
	if m == nil {
		return "", "", "", false
	}
	return m[1], strings.TrimSpace(m[2]), strings.TrimSpace(m[3]), true
}

type span struct{ lo, hi int }

func zeroStores(lines []string, reg string) []span {
	var out []span
	for _, instr := range lines {
		op, src, dest, ok := parseStore(instr)
		if !ok || !isZeroSrc(src) {
			continue
		}
		m := reDispReg.FindStringSubmatch(dest)
		if m == nil || m[2] != reg {
			continue
		}
		off, err := parseDisp(m[1])
		if err != nil {
			continue
		}
		w := storeWidth(op, src)
		if w <= 0 {
			continue
		}
		out = append(out, span{off, off + w})
	}
	return mergeSpans(out)
}

func isZeroSrc(src string) bool {
	s := strings.TrimSpace(strings.ToUpper(src))
	switch s {
	case "X15", "Y15", "Z15":
		return true
	}
	if strings.HasPrefix(s, "$") {
		n, err := strconv.ParseInt(s[1:], 0, 64)
		return err == nil && n == 0
	}
	return false
}

func storeWidth(op, src string) int {
	u := strings.ToUpper(op)
	srcU := strings.ToUpper(src)
	switch {
	case strings.Contains(u, "MOVUPS") || strings.Contains(u, "MOVAPS") || strings.Contains(u, "MOVDQU") || strings.Contains(u, "MOVDQA"):
		if strings.HasPrefix(srcU, "Z") {
			return 64
		}
		if strings.HasPrefix(srcU, "Y") {
			return 32
		}
		return 16
	case strings.HasSuffix(u, "Q"):
		return 8
	case strings.HasSuffix(u, "L"):
		return 4
	case strings.HasSuffix(u, "W"):
		return 2
	case strings.HasSuffix(u, "B"):
		return 1
	default:
		return 0
	}
}

func parseDisp(s string) (int, error) {
	if s == "" {
		return 0, nil
	}
	n, err := strconv.ParseInt(s, 0, 64)
	return int(n), err
}

func mergeSpans(in []span) []span {
	if len(in) == 0 {
		return nil
	}
	s := append([]span(nil), in...)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j].lo < s[i].lo {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	out := []span{s[0]}
	for _, x := range s[1:] {
		last := &out[len(out)-1]
		if x.lo <= last.hi {
			if x.hi > last.hi {
				last.hi = x.hi
			}
			continue
		}
		out = append(out, x)
	}
	return out
}

func maxRun(s []span) int {
	best := 0
	for _, x := range s {
		if d := x.hi - x.lo; d > best {
			best = d
		}
	}
	return best
}
