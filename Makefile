.PHONY: test interop-driver interop-goldens interop-test test-sgoiter ci

export GOWORK ?= off

test:
	go test -count=1 ./...
	go test -count=1 ./internal/lsstream/
	go test -count=1 ./internal/monocypher55/

interop-driver:
	mkdir -p testdata/libsodium_interop/bin
	cc -O2 testdata/libsodium_interop/driver_secretstream.c \
	  $$(pkg-config --cflags --libs libsodium) \
	  -o testdata/libsodium_interop/bin/driver_secretstream

interop-goldens: interop-driver
	bash scripts/gen_goldens.sh

interop-test: interop-driver
	go test -count=1 -run 'Interop|Libsodium' ./...

test-sgoiter:
	go test -count=1 -tags aead_c2simd ./...

ci: test interop-test
