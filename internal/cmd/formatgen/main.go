// SPDX-License-Identifier: Apache-2.0 OR MIT

package main

import (
	"fmt"
	"os"

	"github.com/hazyhaar/go-secretstream/internal/formatgen"
)

func main() {
	spec := "spec"
	out := "."
	if len(os.Args) > 1 {
		spec = os.Args[1]
	}
	if len(os.Args) > 2 {
		out = os.Args[2]
	}
	if err := formatgen.Generate(spec, out); err != nil {
		fmt.Fprintf(os.Stderr, "formatgen: %v\n", err)
		os.Exit(1)
	}
}
