.PHONY: test test-fallback interop-driver interop-goldens interop-test ci

export GOWORK ?= off

# Recette nominale : chemin SIMD exigé (Go 1.27, GOEXPERIMENT=simd, amd64
# AVX2). SECRETSTREAM_REQUIRE_SIMD=1 fait échouer la suite si le repli
# scalaire est sélectionné par accident (TestSIMDSelectionNotSilent).
test:
	GOEXPERIMENT=simd SECRETSTREAM_REQUIRE_SIMD=1 go test -count=1 ./...
	GOEXPERIMENT=simd SECRETSTREAM_REQUIRE_SIMD=1 go test -count=1 ./internal/lsstream/
	GOEXPERIMENT=simd go test -count=1 ../monocypher55/

# Repli scalaire, explicitement : mêmes tests sans SIMD (le bench gate se
# déclare non applicable au lieu d'échouer).
test-fallback:
	go test -count=1 ./...
	go test -count=1 ../monocypher55/

interop-driver:
	mkdir -p testdata/libsodium_interop/bin
	cc -O2 testdata/libsodium_interop/driver_secretstream.c \
	  $$(pkg-config --cflags --libs libsodium) \
	  -o testdata/libsodium_interop/bin/driver_secretstream

interop-goldens: interop-driver
	bash scripts/gen_goldens.sh

interop-test: interop-driver
	go test -count=1 -run 'Interop|Libsodium' ./...

ci: test interop-test
