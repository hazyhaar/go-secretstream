// High-level crypter (WAL-G style API), pure Go, no CGO.

package secretstream

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

// Crypter is a reusable encrypt/decrypt factory (key loaded once).
type Crypter struct {
	key []byte

	KeyInline string

	KeyPath string

	KeyTransform string

	mutex sync.RWMutex
}

// Name returns the crypter label.
func (c *Crypter) Name() string {
	return "Libsodium"
}

// CrypterFromKey builds a Crypter from inline key material + transform.
func CrypterFromKey(key string, keyTransform string) *Crypter {
	return &Crypter{KeyInline: key, KeyTransform: keyTransform}
}

// CrypterFromKeyPath builds a Crypter that loads the key from path on first use.
func CrypterFromKeyPath(path string, keyTransform string) *Crypter {
	return &Crypter{KeyPath: path, KeyTransform: keyTransform}
}

func (c *Crypter) setup() error {
	c.mutex.RLock()
	if c.key != nil {
		c.mutex.RUnlock()
		return nil
	}
	c.mutex.RUnlock()
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.key != nil {
		return nil
	}
	if c.KeyInline == "" && c.KeyPath == "" {
		return fmt.Errorf("secretstream Crypter: must have a key or key path")
	}
	keyString := c.KeyInline
	if keyString == "" {
		keyFileContents, err := os.ReadFile(c.KeyPath)
		if err != nil {
			return fmt.Errorf("secretstream Crypter: unable to read key from file: %v", err)
		}
		keyString = strings.TrimSpace(string(keyFileContents))
	}
	key, err := KeyTransform(keyString, c.KeyTransform, KeyBytes)
	if err != nil {
		return fmt.Errorf("secretstream Crypter: during key transform: %v", err)
	}
	c.key = key
	return nil
}

// Encrypt returns a WriteCloser that encrypts to writer (Close emits TAG_FINAL).
func (c *Crypter) Encrypt(writer io.Writer) (io.WriteCloser, error) {
	if err := c.setup(); err != nil {
		return nil, err
	}
	return NewWriter(writer, c.key), nil
}

// Decrypt returns a Reader that decrypts from reader.
func (c *Crypter) Decrypt(reader io.Reader) (io.Reader, error) {
	if err := c.setup(); err != nil {
		return nil, err
	}
	return NewReader(reader, c.key), nil
}
