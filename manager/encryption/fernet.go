package encryption

import (
	"github.com/moby/swarmkit/v2/api"

	"github.com/fernet/fernet-go"
)

// Fernet wraps the `fernet` library as an implementation of encrypter/decrypter.
type Fernet struct {
	key fernet.Key
}

// NewFernet returns a new Fernet encrypter/decrypter with the given key
func NewFernet(key []byte) Fernet { _ = "STUB: not implemented"; return *new(Fernet) }

// Algorithm returns the type of algorithm this is (Fernet, which uses AES128-CBC)
func (f Fernet) Algorithm() api.MaybeEncryptedRecord_Algorithm {
	_ = "STUB: not implemented"
	return *new(api.MaybeEncryptedRecord_Algorithm)
}

// Encrypt encrypts some bytes and returns an encrypted record
func (f Fernet) Encrypt(data []byte) (*api.MaybeEncryptedRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fernet generates its own IVs, so nonce is empty

// Decrypt decrypts a MaybeEncryptedRecord and returns some bytes
func (f Fernet) Decrypt(record api.MaybeEncryptedRecord) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// -1 skips the TTL check, since we don't care about message expiry

// VerifyandDecrypt returns a nil message if it can't be verified and decrypted
