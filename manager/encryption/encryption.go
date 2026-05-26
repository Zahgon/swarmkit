package encryption

import (
	"github.com/moby/swarmkit/v2/api"
)

// This package defines the interfaces and encryption package

const humanReadablePrefix = "SWMKEY-1-"

// ErrCannotDecrypt is the type of error returned when some data cannot be decryptd as plaintext
type ErrCannotDecrypt struct {
	msg string
}

func (e ErrCannotDecrypt) Error() string {
	_ = "STUB: not implemented"

	// A Decrypter can decrypt an encrypted record
	return ""
}

type Decrypter interface {
	Decrypt(api.MaybeEncryptedRecord) ([]byte, error)
}

// A Encrypter can encrypt some bytes into an encrypted record
type Encrypter interface {
	Encrypt(data []byte) (*api.MaybeEncryptedRecord, error)
}

type noopCrypter struct{}

func (n noopCrypter) Decrypt(e api.MaybeEncryptedRecord) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n noopCrypter) Encrypt(data []byte) (*api.MaybeEncryptedRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n noopCrypter) Algorithm() api.MaybeEncryptedRecord_Algorithm {
	_ = "STUB: not implemented"
	return *new(api.MaybeEncryptedRecord_Algorithm)
}

// NoopCrypter is just a pass-through crypter - it does not actually encrypt or
// decrypt any data
var NoopCrypter = noopCrypter{}

// specificDecryptor represents a specific type of Decrypter, like NaclSecretbox or Fernet.
// It does not apply to a more general decrypter like MultiDecrypter.
type specificDecrypter interface {
	Decrypter
	Algorithm() api.MaybeEncryptedRecord_Algorithm
}

// MultiDecrypter is a decrypter that will attempt to decrypt with multiple decrypters.  It
// references them by algorithm, so that only the relevant decrypters are checked instead of
// every single one. The reason for multiple decrypters per algorithm is to support hitless
// encryption key rotation.
//
// For raft encryption for instance, during an encryption key rotation, it's possible to have
// some raft logs encrypted with the old key and some encrypted with the new key, so we need a
// decrypter that can decrypt both.
type MultiDecrypter struct {
	decrypters map[api.MaybeEncryptedRecord_Algorithm][]Decrypter
}

// Decrypt tries to decrypt using any decrypters that match the given algorithm.
func (m MultiDecrypter) Decrypt(r api.MaybeEncryptedRecord) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewMultiDecrypter returns a new MultiDecrypter given multiple Decrypters.  If any of
// the Decrypters are also MultiDecrypters, they are flattened into a single map, but
// it does not deduplicate any decrypters.
// Note that if something is neither a MultiDecrypter nor a specificDecrypter, it is
// ignored.
func NewMultiDecrypter(decrypters ...Decrypter) MultiDecrypter {
	_ = "STUB: not implemented"
	return *new(MultiDecrypter)
}

// Decrypt turns a slice of bytes serialized as an MaybeEncryptedRecord into a slice of plaintext bytes
func Decrypt(encryptd []byte, decrypter Decrypter) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nope, this wasn't marshalled as a MaybeEncryptedRecord

// Encrypt turns a slice of bytes into a serialized MaybeEncryptedRecord slice of bytes
func Encrypt(plaintext []byte, encrypter Encrypter) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Defaults returns a default encrypter and decrypter.  If the FIPS parameter is set to
// true, the only algorithm supported on both the encrypter and decrypter will be fernet.
func Defaults(key []byte, fips bool) (Encrypter, Decrypter) {
	_ = "STUB: not implemented"
	return *new(Encrypter), *new(Decrypter)
}

// GenerateSecretKey generates a secret key that can be used for encrypting data
// using this package
func GenerateSecretKey() []byte { _ = "STUB: not implemented"; return nil }

// panic if we can't read random data

// HumanReadableKey displays a secret key in a human readable way
func HumanReadableKey(key []byte) string {
	_ = "STUB: not implemented"
	// base64-encode the key
	return ""
}

// ParseHumanReadableKey returns a key as bytes from recognized serializations of
// said keys
func ParseHumanReadableKey(key string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
