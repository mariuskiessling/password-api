package password

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"log"
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

// Encrypt takes a password and encrypts it with the user's public key. The
// password is encrypted using RSA-OAEP with SHA256. After encrypting the
// password is encoded as a base64 string.
func (pk *PublicKey) Encrypt(pw string) (encryptedPw string, err error) {
	key, err := x509.ParsePKIXPublicKey(pk.Block.Bytes)
	if err != nil {
		log.Println(err)
		return "", errors.New("Could not parse PKCS1 public key to internal RSA key.")
	}
	cipher, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, key.(*rsa.PublicKey), []byte(pw), []byte(""))
	if err != nil {
		log.Println(err)
		return "", errors.New("Could not encrypt password.")
	}

	return base64.StdEncoding.EncodeToString(cipher), nil
}
