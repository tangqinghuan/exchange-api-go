package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash"
)

// Signer signs provided payloads.
type Signer interface {
	// Sign signs provided payload and returns encoded string sum.
	Sign(payload []byte) string
}

// HmacSigner uses HMAC for sign.
type HmacSigner struct {
	h   func() hash.Hash
	e   func(src []byte) string
	key []byte
}

// NewHmacSigner ...
func NewHmacSigner(h func() hash.Hash, e func(src []byte) string, key []byte) Signer {
	return &HmacSigner{h, e, key}
}

// NewHmacSha224HexSigner The signature is the hex digest of an HMAC-SHA224 hash where the message is your payload,
// and the secret key is your API secret.
func NewHmacSha224HexSigner(key []byte) Signer {
	return NewHmacSigner(sha256.New224, hex.EncodeToString, key)
}

// NewHmacSha256HexSigner The signature is the hex digest of an HMAC-SHA256 hash where the message is your payload,
// and the secret key is your API secret.
func NewHmacSha256HexSigner(key []byte) Signer {
	return NewHmacSigner(sha256.New, hex.EncodeToString, key)
}

// NewHmacSha384HexSigner The signature is the hex digest of an HMAC-SHA256 hash where the message is your payload,
// and the secret key is your API secret.
func NewHmacSha384HexSigner(key []byte) Signer {
	return NewHmacSigner(sha512.New384, hex.EncodeToString, key)
}

// NewHmacSha512HexSigner ...
func NewHmacSha512HexSigner(key []byte) Signer {
	return NewHmacSigner(sha512.New, hex.EncodeToString, key)
}

// NewHmacSha256Base64Signer The signature is the base64 of an HMAC-SHA256 hash where the message is your payload,
// and the secret key is your API secret.
func NewHmacSha256Base64Signer(key []byte) Signer {
	return NewHmacSigner(sha256.New, base64.StdEncoding.EncodeToString, key)
}

// Sign signs provided payload and returns encoded string sum.
func (s *HmacSigner) Sign(payload []byte) string {
	mac := hmac.New(s.h, s.key)
	mac.Write(payload)
	return s.e(mac.Sum(nil))
}
