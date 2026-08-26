// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/hazyhaar/go-secretstream/internal/formatgen"
)

func TestFormatCUEDrift(t *testing.T) {
	tmp := t.TempDir()
	if err := formatgen.Generate("spec", tmp); err != nil {
		t.Fatalf("régénération: %v", err)
	}
	products := []string{
		"format_v2_gen.go",
		"FORMAT_V2.md",
		filepath.Join("testdata", "v2", "vectors.json"),
	}
	for _, rel := range products {
		got, err := os.ReadFile(filepath.Join(tmp, rel))
		if err != nil {
			t.Fatalf("produit régénéré %s: %v", rel, err)
		}
		want, err := os.ReadFile(rel)
		if err != nil {
			t.Fatalf("produit suivi %s: %v", rel, err)
		}
		if !bytes.Equal(got, want) {
			t.Errorf("dérive %s: régénéré %d octets, suivi %d octets", rel, len(got), len(want))
		}
	}

	tracked, err := os.ReadFile("format_v2_gen.go")
	if err != nil {
		t.Fatal(err)
	}
	regen, err := os.ReadFile(filepath.Join(tmp, "format_v2_gen.go"))
	if err != nil {
		t.Fatal(err)
	}
	mutated := bytes.Replace(regen, []byte("HeaderSize = 36"), []byte("HeaderSize = 37"), 1)
	if bytes.Equal(mutated, regen) {
		t.Fatal("mutation de HeaderSize dans t.TempDir() n'a pas pris")
	}
	if bytes.Equal(mutated, tracked) {
		t.Fatal("copie mutée encore égale au fichier suivi : la garde de dérive serait insensible")
	}
}
