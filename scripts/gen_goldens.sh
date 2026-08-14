#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
BIN="$ROOT/testdata/libsodium_interop/bin/driver_secretstream"
GOLD="$ROOT/testdata/libsodium_interop/golden"
mkdir -p "$GOLD" "$(dirname "$BIN")"

cc -O2 "$ROOT/testdata/libsodium_interop/driver_secretstream.c" \
  $(pkg-config --cflags --libs libsodium) \
  -o "$BIN"

MANIFEST="$GOLD/manifest.json"
echo '[' >"$MANIFEST"
first=1
sizes=(0 1 15 16 63 64 65 1000 8191 8192 8193 20000 65536 100000)
for sz in "${sizes[@]}"; do
  id=$(printf 'n%06d' "$sz")
  keyf="$GOLD/${id}.key"
  plainf="$GOLD/${id}.plain"
  wiref="$GOLD/${id}.wire"
  # fixed key for reproducibility
  python3 -c "open('$keyf','wb').write(bytes([0x11+(i%200) for i in range(32)]))"
  python3 -c "open('$plainf','wb').write(bytes([i%251 for i in range($sz)]))"
  "$BIN" push-stream "$keyf" "$plainf" "$wiref"
  # verify pull
  outf="$GOLD/${id}.pull.out"
  "$BIN" pull-stream "$keyf" "$wiref" "$outf"
  cmp -s "$plainf" "$outf"
  rm -f "$outf"
  sha_p=$(sha256sum "$plainf" | awk '{print $1}')
  sha_w=$(sha256sum "$wiref" | awk '{print $1}')
  if [[ $first -eq 0 ]]; then echo ',' >>"$MANIFEST"; fi
  first=0
  printf '  {"id":"%s","bytes":%s,"sha256_plain":"%s","sha256_wire":"%s"}' \
    "$id" "$sz" "$sha_p" "$sha_w" >>"$MANIFEST"
  echo "OK $id ($sz B)"
done
echo >>"$MANIFEST"
echo ']' >>"$MANIFEST"
echo "goldens -> $GOLD"
