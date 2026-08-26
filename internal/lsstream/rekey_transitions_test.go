// SPDX-License-Identifier: Apache-2.0

package lsstream

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

type scriptOp struct {
	kind string
	tag  byte
	n    int
	ctr  string
}

func driverBin(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("caller")
	}
	p := filepath.Join(filepath.Dir(file), "..", "..", "testdata", "libsodium_interop", "bin", "driver_secretstream")
	if _, err := os.Stat(p); err != nil {
		t.Fatalf("pilote C absent (%s) — lancer make interop-driver", p)
	}
	return p
}

func detMsg(n int) []byte {
	m := make([]byte, n)
	for i := range m {
		m[i] = byte((i*7 + n) & 0xff)
	}
	return m
}

func parseScript(s string) []scriptOp {
	var ops []scriptOp
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		f := strings.Fields(line)
		switch f[0] {
		case "force-counter":
			ops = append(ops, scriptOp{kind: "force-counter", ctr: f[1]})
		case "rekey":
			ops = append(ops, scriptOp{kind: "rekey"})
		case "msg":
			var tag byte
			switch f[1] {
			case "MESSAGE":
				tag = TagMessage
			case "PUSH":
				tag = TagPush
			case "REKEY":
				tag = TagRekey
			case "FINAL":
				tag = TagFinal
			default:
				panic("tag " + f[1])
			}
			var n int
			_, _ = fmt.Sscanf(f[2], "%d", &n)
			ops = append(ops, scriptOp{kind: "msg", tag: tag, n: n})
		default:
			panic("cmd " + f[0])
		}
	}
	return ops
}

func forceCounter(st *streamState, hexctr string) error {
	b, err := hex.DecodeString(hexctr)
	if err != nil || len(b) != 4 {
		return fmt.Errorf("force-counter %q", hexctr)
	}
	copy(st.nonce[0:4], b)
	return st.initCipher()
}

func goPush(key, header []byte, ops []scriptOp) ([]byte, error) {
	st, err := initFromHeader(key, header)
	if err != nil {
		return nil, err
	}
	var out bytes.Buffer
	out.Write(header)
	for _, op := range ops {
		switch op.kind {
		case "force-counter":
			if err := forceCounter(st, op.ctr); err != nil {
				return nil, err
			}
		case "rekey":
			if err := st.rekey(); err != nil {
				return nil, err
			}
		case "msg":
			chunk, err := st.push(detMsg(op.n), op.tag)
			if err != nil {
				return nil, err
			}
			out.Write(chunk)
		}
	}
	return out.Bytes(), nil
}

func goPull(key, wire []byte, ops []scriptOp) (lines []string, plain []byte, err error) {
	if len(wire) < HeaderBytes {
		return nil, nil, fmt.Errorf("wire trop court")
	}
	st, err := initPull(key, wire[:HeaderBytes])
	if err != nil {
		return nil, nil, err
	}
	off := HeaderBytes
	for _, op := range ops {
		switch op.kind {
		case "force-counter":
			if err := forceCounter(st, op.ctr); err != nil {
				return lines, plain, err
			}
		case "rekey":
			if err := st.rekey(); err != nil {
				return lines, plain, err
			}
		case "msg":
			clen := op.n + ABytes
			if off+clen > len(wire) {
				return lines, plain, fmt.Errorf("chunk tronqué off=%d clen=%d wire=%d", off, clen, len(wire))
			}
			m, tag, perr := st.pull(wire[off : off+clen])
			if perr != nil {
				return append(lines, "MAC fail"), plain, perr
			}
			lines = append(lines, fmt.Sprintf("tag=%d len=%d ok", tag, len(m)))
			plain = append(plain, m...)
			off += clen
		}
	}
	return lines, plain, nil
}

func expectPlain(ops []scriptOp) []byte {
	var p []byte
	for _, op := range ops {
		if op.kind == "msg" {
			p = append(p, detMsg(op.n)...)
		}
	}
	return p
}

func firstDiverging(a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return i
		}
	}
	if len(a) != len(b) {
		return n
	}
	return -1
}

func runTransition(t *testing.T, script string) {
	t.Helper()
	drv := driverBin(t)
	dir := t.TempDir()
	key := bytes.Repeat([]byte{0x5a}, KeyBytes)
	keyf := filepath.Join(dir, "key.bin")
	scriptf := filepath.Join(dir, "script.txt")
	wireCf := filepath.Join(dir, "wire_c.bin")
	wireGf := filepath.Join(dir, "wire_g.bin")
	plainCf := filepath.Join(dir, "plain_c.bin")
	plainGf := filepath.Join(dir, "plain_g.bin")
	if err := os.WriteFile(keyf, key, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(scriptf, []byte(script), 0o600); err != nil {
		t.Fatal(err)
	}
	if out, err := exec.Command(drv, "push-script", keyf, scriptf, wireCf).CombinedOutput(); err != nil {
		t.Fatalf("c push-script: %v\n%s", err, out)
	}
	wireC, err := os.ReadFile(wireCf)
	if err != nil {
		t.Fatal(err)
	}
	if len(wireC) < HeaderBytes {
		t.Fatalf("fil C trop court (%d)", len(wireC))
	}
	ops := parseScript(script)
	wireG, err := goPush(key, wireC[:HeaderBytes], ops)
	if err != nil {
		t.Fatalf("go push: %v", err)
	}
	if !bytes.Equal(wireG, wireC) {
		d := firstDiverging(wireG, wireC)
		t.Fatalf("cmp fil Go / fil C : diverge octet %d (go=%d c=%d)", d, len(wireG), len(wireC))
	}
	if err := os.WriteFile(wireGf, wireG, 0o600); err != nil {
		t.Fatal(err)
	}
	cPull := exec.Command(drv, "pull-script", keyf, wireGf, scriptf, plainGf)
	cOut, err := cPull.CombinedOutput()
	if err != nil {
		t.Fatalf("c pull-script sur fil Go: %v\n%s", err, cOut)
	}
	goLines, goPlain, err := goPull(key, wireC, ops)
	if err != nil {
		t.Fatalf("go pull sur fil C: %v\n%s", err, strings.Join(goLines, "\n"))
	}
	wantPlain := expectPlain(ops)
	if !bytes.Equal(goPlain, wantPlain) {
		t.Fatalf("go plaintext diverge (got=%d want=%d)", len(goPlain), len(wantPlain))
	}
	cPlain, err := os.ReadFile(plainGf)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(cPlain, wantPlain) {
		t.Fatalf("c plaintext diverge (got=%d want=%d)", len(cPlain), len(wantPlain))
	}
	cLines := strings.TrimSpace(string(cOut))
	gLines := strings.Join(goLines, "\n")
	if cLines != gLines {
		t.Fatalf("lignes pull C vs Go:\nC:\n%s\nGo:\n%s", cLines, gLines)
	}
	_ = os.WriteFile(plainCf, goPlain, 0o600)
	t.Logf("cmp identique len=%d\nC pull-script sur fil Go:\n%s\nGo pull sur fil C:\n%s", len(wireC), cLines, gLines)
}

func TestInterop_RekeyTransitions(t *testing.T) {
	cases := []struct {
		name   string
		script string
	}{
		{"T01", "msg MESSAGE 4096\nmsg FINAL 0\n"},
		{"T02", "msg PUSH 100\nmsg MESSAGE 100\nmsg FINAL 0\n"},
		{"T03", "msg REKEY 100\nmsg MESSAGE 100\nmsg FINAL 0\n"},
		{"T04", "msg FINAL 100\nmsg MESSAGE 100\n"},
		{"T05", "msg MESSAGE 100\nrekey\nmsg MESSAGE 100\nmsg FINAL 0\n"},
		{"T06", "force-counter FFFFFFFF\nmsg MESSAGE 100\nmsg MESSAGE 100\nmsg FINAL 0\n"},
		{"T07", "force-counter FFFFFFFF\nmsg REKEY 100\nmsg MESSAGE 100\nmsg FINAL 0\n"},
		{"T08", "force-counter FFFFFFFF\nmsg PUSH 100\nmsg MESSAGE 100\nmsg FINAL 0\n"},
		{"T09", "force-counter FFFFFFFF\nmsg FINAL 100\n"},
		{"T10", "msg MESSAGE 0\nmsg MESSAGE 100\nmsg FINAL 0\n"},
		{"T11", "msg MESSAGE 100\nmsg FINAL 0\n"},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			runTransition(t, tc.script)
			if tc.name == "T11" {
				runTransition(t, "msg FINAL 100\n")
			}
		})
	}
}
