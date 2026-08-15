module github.com/hazyhaar/go-secretstream

go 1.26

require (
	code.hazyhaar.fr/devhoros/c2simd v0.0.0
	golang.org/x/crypto v0.54.0
)

replace code.hazyhaar.fr/devhoros/c2simd => github.com/hazyhaar/c2simd v0.0.0-20260815212259-38aa4a2a8809

require golang.org/x/sys v0.47.0 // indirect
