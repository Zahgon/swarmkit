package ca

import (
	"encoding/pem"
	"sync"

	"github.com/moby/swarmkit/v2/ca/keyutils"
)

const (
	// keyPerms are the permissions used to write the TLS keys
	keyPerms = 0o600
	// certPerms are the permissions used to write TLS certificates
	certPerms = 0o644
	// versionHeader is the TLS PEM key header that contains the KEK version
	versionHeader = "kek-version"
)

// PEMKeyHeaders is an interface for something that needs to know about PEM headers
// when reading or writing TLS keys in order to keep them updated with the latest
// KEK.
type PEMKeyHeaders interface {

	// UnmarshalHeaders loads the headers map given the current KEK
	UnmarshalHeaders(map[string]string, KEKData) (PEMKeyHeaders, error)

	// MarshalHeaders returns a header map given the current KEK
	MarshalHeaders(KEKData) (map[string]string, error)

	// UpdateKEK gets called whenever KeyReadWriter gets a KEK update.  This allows the
	// PEMKeyHeaders to optionally update any internal state.  It should return an
	// updated (if needed) versino of itself.
	UpdateKEK(KEKData, KEKData) PEMKeyHeaders
}

// KeyReader reads a TLS cert and key from disk
type KeyReader interface {

	// Read reads and returns the certificate and the key PEM bytes that are on disk
	Read() ([]byte, []byte, error)

	// Target returns a string representation of where the cert data is being read from
	Target() string
}

// KeyWriter writes a TLS key and cert to disk
type KeyWriter interface {

	// Write accepts a certificate and key in PEM format, as well as an optional KEKData object.
	// If there is a current KEK, the key is encrypted using the current KEK.  If the KEKData object
	// is provided (not nil), the key will be encrypted using the new KEK instead, and the current
	// KEK in memory will be replaced by the provided KEKData.  The reason to allow changing key
	// material and KEK in a single step, as opposed to two steps, is to prevent the key material
	// from being written unencrypted or with an old KEK in the first place (when a node gets a
	// certificate from the CA, it will also request the current KEK so it won't have to immediately
	// do a KEK rotation after getting the key).
	Write([]byte, []byte, *KEKData) error

	// ViewAndUpdateHeaders is a function that reads and updates the headers of the key in a single
	// transaction (e.g. within a lock).  It accepts a callback function which will be passed the
	// current header management object, and which must return a new, updated, or same header
	// management object.  KeyReadWriter then performs the following actions:
	//    - uses the old header management object and the current KEK to deserialize/decrypt
	//      the existing PEM headers
	//    - uses the new header management object and the current KEK to to reserialize/encrypt
	//      the PEM headers
	//    - writes the new PEM headers, as well as the key material, unchanged, to disk
	ViewAndUpdateHeaders(func(PEMKeyHeaders) (PEMKeyHeaders, error)) error

	// ViewAndRotateKEK is a function that just re-encrypts the TLS key and headers in a single
	// transaction (e.g. within a lock).  It accepts a callback unction which will be passed the
	// current KEK and the current headers management object, and which should return a new
	// KEK and header management object.  KeyReadWriter then performs the following actions:
	//    - uses the old KEK and header management object to deserialize/decrypt the
	//      TLS key and PEM headers
	//    - uses the new KEK and header management object to serialize/encrypt the TLS key
	//      and PEM headers
	//    - writes the new PEM headers and newly encrypted TLS key to disk
	ViewAndRotateKEK(func(KEKData, PEMKeyHeaders) (KEKData, PEMKeyHeaders, error)) error

	// GetCurrentState returns the current header management object and the current KEK.
	GetCurrentState() (PEMKeyHeaders, KEKData)

	// Target returns a string representation of where the cert data is being read from
	Target() string
}

// KEKData provides an optional update to the kek when writing.  The structure
// is needed so that we can tell the difference between "do not encrypt anymore"
// and there is "no update".
type KEKData struct {
	KEK     []byte
	Version uint64
}

// ErrInvalidKEK means that we cannot decrypt the TLS key for some reason
type ErrInvalidKEK struct {
	Wrapped error
}

func (e ErrInvalidKEK) Error() string { _ = "STUB: not implemented"; return "" }

// KeyReadWriter is an object that knows how to read and write TLS keys and certs to disk,
// optionally encrypted and optionally updating PEM headers.  It should be the only object which
// can write the TLS key, to ensure that writes are serialized and that the TLS key, the
// KEK (key encrypting key), and any headers which need to be written are never out of sync.
// It accepts a PEMKeyHeaders object, which is used to serialize/encrypt and deserialize/decrypt
// the PEM headers when given the current headers and the current KEK.
type KeyReadWriter struct {

	// This lock is held whenever a key is read from or written to disk, or whenever the internal
	// state of the KeyReadWriter (such as the KEK, the key formatter, or the PEM header management
	// object changes.)
	mu sync.Mutex

	kekData      KEKData
	paths        CertPaths
	headersObj   PEMKeyHeaders
	keyFormatter keyutils.Formatter
}

// NewKeyReadWriter creates a new KeyReadWriter
func NewKeyReadWriter(paths CertPaths, kek []byte, headersObj PEMKeyHeaders) *KeyReadWriter {
	_ = "STUB: not implemented"
	return nil
}

// SetKeyFormatter sets the keyformatter with which to encrypt and decrypt keys
func (k *KeyReadWriter) SetKeyFormatter(kf keyutils.Formatter) { _ = "STUB: not implemented"; return }

// Migrate checks to see if a temporary key file exists.  Older versions of
// swarmkit wrote temporary keys instead of temporary certificates, so
// migrate that temporary key if it exists.  We want to write temporary certificates,
// instead of temporary keys, because we may need to periodically re-encrypt the
// keys and modify the headers, and it's easier to have a single canonical key
// location than two possible key locations.
func (k *KeyReadWriter) Migrate() error { _ = "STUB: not implemented"; return nil }

// no key?  no migration

// it does exist - no need to decrypt, because previous versions of swarmkit
// which supported this temporary key did not support encrypting TLS keys

// no cert?  no migration

// nope, this does not match the cert

// Read will read a TLS cert and key from the given paths
func (k *KeyReadWriter) Read() ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// The cert is written to a temporary file first, then the key, and then
// the cert gets renamed - so, if interrupted, it's possible to end up with
// a cert that only exists in the temporary location.

//continue to try temp location

// either the cert doesn't exist, or it doesn't match the key - try the temp file, if it exists

// return the original error

// nope, it doesn't match either - remove and return the original error

// try to move the temp cert back to the regular location

// ViewAndRotateKEK re-encrypts the key with a new KEK
func (k *KeyReadWriter) ViewAndRotateKEK(cb func(KEKData, PEMKeyHeaders) (KEKData, PEMKeyHeaders, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// ViewAndUpdateHeaders updates the header manager, and updates any headers on the existing key
func (k *KeyReadWriter) ViewAndUpdateHeaders(cb func(PEMKeyHeaders) (PEMKeyHeaders, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// we WANT any original encryption headers

// GetCurrentState returns the current KEK data, including version
func (k *KeyReadWriter) GetCurrentState() (PEMKeyHeaders, KEKData) {
	_ = "STUB: not implemented"
	return *new(PEMKeyHeaders), *new(KEKData)
}

// Write attempts write a cert and key to text.  This can also optionally update
// the KEK while writing, if an updated KEK is provided.  If the pointer to the
// update KEK is nil, then we don't update. If the updated KEK itself is nil,
// then we update the KEK to be nil (data should be unencrypted).
func (k *KeyReadWriter) Write(certBytes, plaintextKeyBytes []byte, kekData *KEKData) error {
	_ = "STUB: not implemented"
	return nil
}

// current assumption is that the cert and key will be in the same directory

// Ensure that we will have a keypair on disk at all times by writing the cert to a
// temp path first.  This is because we want to have only a single copy of the key
// for rotation and header modification.

func (k *KeyReadWriter) genTempPaths() CertPaths { _ = "STUB: not implemented"; return *new(CertPaths) }

// Target returns a string representation of this KeyReadWriter, namely where
// it is writing to
func (k *KeyReadWriter) Target() string { _ = "STUB: not implemented"; return "" }

func (k *KeyReadWriter) readKeyblock() (*pem.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decode the PEM private key

// readKey returns the decrypted key pem bytes, and enforces the KEK if applicable
// (writes it back with the correct encryption if it is not correctly encrypted)
func (k *KeyReadWriter) readKey() (*pem.Block, error) { _ = "STUB: not implemented"; return nil, nil }

// If it's encrypted, we can't read without a passphrase (we're assuming
// empty passphrases are invalid)

// change header only if its pkcs8

// remove encryption PEM headers

// the key type doesn't change

// writeKey takes an unencrypted keyblock and, if the kek is not nil, encrypts it before
// writing it to disk.  If the kek is nil, writes it to disk unencrypted.
func (k *KeyReadWriter) writeKey(keyBlock *pem.Block, kekData KEKData, pkh PEMKeyHeaders) error {
	_ = "STUB: not implemented"
	return nil
}

// DowngradeKey converts the PKCS#8 key to PKCS#1 format and save it
func (k *KeyReadWriter) DowngradeKey() error { _ = "STUB: not implemented"; return nil }

// stop if the key is already downgraded to pkcs1

// add kek-version header back to the new key

// do not use krw.Write as it will convert the key to pkcs8

// merges one set of PEM headers onto another, excepting for key encryption value
// "proc-type" and "dek-info"
func mergePEMHeaders(original, newSet map[string]string) { _ = "STUB: not implemented"; return }
