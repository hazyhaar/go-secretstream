// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55_test

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// TestNoHostAbsolutePaths interdit la présence de tout chemin absolu hôte
// dans les sources et tests de secretstream55.
func TestNoHostAbsolutePaths(t *testing.T) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("impossible de résoudre runtime.Caller")
	}
	pkgDir := filepath.Dir(currentFile)

	forbidden := []string{
		// Motifs = guillemet ouvrant + chemin : un LITTÉRAL de chaîne qui commence
		// par un chemin absolu hôte. Le chemin de module (`code.hazyhaar.fr/devhoros/…`
		// dans les imports) n'est pas un chemin hôte et ne doit pas déclencher.
		// Les octets sont épelés pour que ce fichier ne se signale pas lui-même
		// à la garde `c2simd-fyne-guard verify-portability`.
		string([]byte{'"', '/', 'd', 'e', 'v', 'h', 'o', 'r', 'o', 's'}),
		string([]byte{'"', '/', 'h', 'o', 'm', 'e', '/'}),
		string([]byte{'"', '/', 't', 'm', 'p', '/', 'c', 'l', 'a', 'u', 'd', 'e'}),
	}

	err := filepath.Walk(pkgDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		if filepath.Base(path) == "portability_test.go" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		src := string(data)
		for _, f := range forbidden {
			if strings.Contains(src, f) {
				t.Errorf("VIOLATION DE PORTABILITÉ : %s contient la chaîne absolue interdite %s", path, f)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("erreur lors du parcours : %v", err)
	}
}
