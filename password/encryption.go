package password

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/pem"
	"errors"
)

// PublicKey stores all information about a user's public key.
type PublicKey struct {
	Block       *pem.Block
	Key         string
	Fingerprint string
}

// LoadPublicKey loads and decodes the user's public key.
func LoadPublicKey(publicKey string) (*PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("Public key could not be decoded.")
	}

	sum := sha256.Sum256(block.Bytes)

	return &PublicKey{
		Block:       block,
		Key:         publicKey,
		Fingerprint: hex.EncodeToString(sum[:]),
	}, nil
}
