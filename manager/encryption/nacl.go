package encryption

import (
	"github.com/moby/swarmkit/v2/api"
)

const naclSecretboxKeySize = 32
const naclSecretboxNonceSize = 24

// This provides the default implementation of an encrypter and decrypter, as well
// as the default KDF function.

// NACLSecretbox is an implementation of an encrypter/decrypter.  Encrypting
// generates random Nonces.
type NACLSecretbox struct {
	key [naclSecretboxKeySize]byte
}

// NewNACLSecretbox returns a new NACL secretbox encrypter/decrypter with the given key
func NewNACLSecretbox(key []byte) NACLSecretbox {
	_ = "STUB: not implemented"
	return *new(NACLSecretbox)
}

// Algorithm returns the type of algorithm this is (NACL Secretbox using XSalsa20 and Poly1305)
func (n NACLSecretbox) Algorithm() api.MaybeEncryptedRecord_Algorithm {
	_ = "STUB: not implemented"
	return *new(api.MaybeEncryptedRecord_Algorithm)
}

// Encrypt encrypts some bytes and returns an encrypted record
func (n NACLSecretbox) Encrypt(data []byte) (*api.MaybeEncryptedRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Seal's first argument is an "out", the data that the new encrypted message should be
// appended to.  Since we don't want to append anything, we pass nil.

// Decrypt decrypts a MaybeEncryptedRecord and returns some bytes
func (n NACLSecretbox) Decrypt(record api.MaybeEncryptedRecord) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Open's first argument is an "out", the data that the decrypted message should be
// appended to.  Since we don't want to append anything, we pass nil.
