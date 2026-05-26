// Package pkcs8 implements functions to encrypt, decrypt, parse and to convert
// EC private keys to PKCS#8 format. However this package is hard forked from
// https://github.com/youmark/pkcs8 and modified function signatures to match
// signatures of crypto/x509 and cloudflare/cfssl/helpers to simplify package
// swapping. License for original package is as follow:

// The MIT License (MIT)
//
// Copyright (c) 2014 youmark
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package pkcs8

import (
	"crypto"
	"encoding/asn1"
	"encoding/pem"
)

// Copy from crypto/x509
var (
	oidPublicKeyECDSA = asn1.ObjectIdentifier{1, 2, 840, 10045, 2, 1}
)

// Unencrypted PKCS#8
var (
	oidPKCS5PBKDF2 = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 5, 12}
	oidPBES2       = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 5, 13}
	oidAES256CBC   = asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 1, 42}
)

type ecPrivateKey struct {
	Version       int
	PrivateKey    []byte
	NamedCurveOID asn1.ObjectIdentifier `asn1:"optional,explicit,tag:0"`
	PublicKey     asn1.BitString        `asn1:"optional,explicit,tag:1"`
}

type privateKeyInfo struct {
	Version             int
	PrivateKeyAlgorithm []asn1.ObjectIdentifier
	PrivateKey          []byte
}

// Encrypted PKCS8
type pbkdf2Params struct {
	Salt           []byte
	IterationCount int
}

type pbkdf2Algorithms struct {
	IDPBKDF2     asn1.ObjectIdentifier
	PBKDF2Params pbkdf2Params
}

type pbkdf2Encs struct {
	EncryAlgo asn1.ObjectIdentifier
	IV        []byte
}

type pbes2Params struct {
	KeyDerivationFunc pbkdf2Algorithms
	EncryptionScheme  pbkdf2Encs
}

type pbes2Algorithms struct {
	IDPBES2     asn1.ObjectIdentifier
	PBES2Params pbes2Params
}

type encryptedPrivateKeyInfo struct {
	EncryptionAlgorithm pbes2Algorithms
	EncryptedData       []byte
}

// ParsePrivateKeyPEMWithPassword parses an encrypted or a decrypted PKCS#8 PEM to crypto.signer
func ParsePrivateKeyPEMWithPassword(pemBytes, password []byte) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

// IsEncryptedPEMBlock checks if a PKCS#8 PEM-block is encrypted or not
func IsEncryptedPEMBlock(block *pem.Block) bool { _ = "STUB: not implemented"; return false }

// DecryptPEMBlock requires PKCS#8 PEM Block and password to decrypt and return unencrypted der []byte
func DecryptPEMBlock(block *pem.Block, password []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Remove padding from key as it might be used to encode to memory as pem

func encryptPrivateKey(pkey, password []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// Calculate key from password based on PKCS5 algorithm
	// Use 8 byte salt, 16 byte IV, and 2048 iteration
	return nil, nil
}

// Use AES256-CBC mode, pad plaintext with PKCS5 padding scheme

// EncryptPEMBlock takes DER-format bytes and password to return an encrypted PKCS#8 PEM-block
func EncryptPEMBlock(data, password []byte) (*pem.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConvertECPrivateKeyPEM takes an EC Private Key as input and returns PKCS#8 version of it
func ConvertECPrivateKeyPEM(inPEM []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// remove curve oid from private bytes as it is already mentioned in algorithm

// ConvertToECPrivateKeyPEM takes an unencrypted PKCS#8 PEM and converts it to
// EC Private Key
func ConvertToECPrivateKeyPEM(inPEM []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
