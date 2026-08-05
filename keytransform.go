package secretstream

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

// Key transform names (WAL-G compatible).
const (
	KeyTransformBase64 = "base64"
	KeyTransformHex    = "hex"
	KeyTransformNone   = "none"
)

const minimalKeyLength = 25

type keyTransformRegEntry struct {
	typ string
	fun func(userInput string) ([]byte, error)
}

var keyTransformReg = []keyTransformRegEntry{
	{typ: KeyTransformBase64, fun: keyTransformBase64},
	{typ: KeyTransformHex, fun: keyTransformHex},
	{typ: KeyTransformNone, fun: keyTransformNone},
}

// KeyTransform decodes userInput according to transformType into exactly
// expectedLen bytes (use KeyBytes for secretstream).
func KeyTransform(userInput string, transformType string, expectedLen int) ([]byte, error) {
	for _, entry := range keyTransformReg {
		if entry.typ == transformType {
			decoded, err := entry.fun(userInput)
			if err != nil {
				return nil, err
			}
			if len(decoded) != expectedLen {
				return nil, fmt.Errorf("key must be exactly %d bytes (got %d bytes)", expectedLen, len(decoded))
			}
			return decoded, nil
		}
	}
	var builder strings.Builder
	for idx, entry := range keyTransformReg {
		if idx > 0 {
			if idx+1 == len(keyTransformReg) {
				builder.WriteString(" or ")
			} else {
				builder.WriteString(", ")
			}
		}
		builder.WriteString(entry.typ)
	}
	return nil, fmt.Errorf("unknown key transform '%s' (must be %s)", transformType, builder.String())
}

func keyTransformBase64(userInput string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(userInput)
	if err != nil {
		return nil, fmt.Errorf("while base64 decoding key: %w", err)
	}
	return decoded, nil
}

func keyTransformHex(userInput string) ([]byte, error) {
	decoded, err := hex.DecodeString(userInput)
	if err != nil {
		return nil, fmt.Errorf("while hex decoding key: %w", err)
	}
	return decoded, nil
}

// keyTransformNone mimics older WAL-G: pad/truncate ASCII to KeyBytes.
//
// WARNING (legacy — prefer hex/base64 with full 32-byte keys):
//   - inputs longer than KeyBytes are silently truncated (prefix collision risk);
//   - inputs in [minimalKeyLength, KeyBytes) are zero-padded (reduced key space).
func keyTransformNone(userInput string) ([]byte, error) {
	if len(userInput) < minimalKeyLength {
		return nil, &ErrShortKey{keyLength: len(userInput)}
	}
	if len(userInput) > KeyBytes {
		return []byte(userInput[:KeyBytes]), nil
	}
	if len(userInput) < KeyBytes {
		buf := make([]byte, KeyBytes)
		copy(buf, userInput)
		return buf, nil
	}
	return []byte(userInput), nil
}

// ErrShortKey is returned when KeyTransformNone receives too little material.
type ErrShortKey struct {
	keyLength int
}

func (e *ErrShortKey) Error() string {
	return fmt.Sprintf("key length must not be less than %v, got %v", minimalKeyLength, e.keyLength)
}
