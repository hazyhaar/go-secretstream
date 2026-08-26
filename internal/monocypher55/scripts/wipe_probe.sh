#!/usr/bin/env bash
# Désassemble un symbole d'un binaire de test (oracle de la garde wipe).
# Fail-closed : binaire absent, vide, ou symbole sans ligne TEXT → code ≠ 0.
# Usage : wipe_probe.sh <binaire> <regex-symbole>
set -euo pipefail
export GOTOOLCHAIN="${GOTOOLCHAIN:-go1.27.0}"
if [[ $# -ne 2 ]]; then
  echo "usage: wipe_probe.sh <binaire> <regex-symbole>" >&2
  exit 2
fi
BIN=$1
RE=$2
if [[ ! -f "$BIN" ]] || [[ ! -s "$BIN" ]]; then
  echo "wipe_probe: binaire introuvable ou vide : $BIN" >&2
  exit 2
fi
out=$(go tool objdump -s "$RE" "$BIN")
if ! grep -q '^TEXT ' <<<"$out"; then
  echo "wipe_probe: symbole introuvable : $RE" >&2
  exit 3
fi
printf '%s\n' "$out"
