package slsa

import (
	"context"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"io"
)

// NewPKCS8Signer can parse either a RSA, ECDSA, or ED25519 private key in PEM
// format and convert it into a dsse signer. It currently doesn't support
// encrypted keys.
//
// For RSA, this uses RSASSA-PKCS1-V1_5-SIGN with SHA256 as the hash function
// For ECDSA, this uses rand.Reader as the source for k
func NewPKCS8Signer(key []byte, id string) (Signer, error) {
	_ = "STUB: not implemented"
	return *new(Signer), nil
}

var _ Signer = (*rsaSigner)(nil)
var _ Signer = (*ecdsaSigner)(nil)
var _ Signer = (*ed25519Signer)(nil)

type rsaSigner struct {
	key   *rsa.PrivateKey
	keyid string
}

func (s *rsaSigner) KeyID() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s *rsaSigner) Sign(ctx context.Context, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ecdsaSigner struct {
	key    *ecdsa.PrivateKey
	keyid  string
	randFn io.Reader // this should be rand.Reader for anything other than tests
}

func (s *ecdsaSigner) KeyID() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s *ecdsaSigner) Sign(ctx context.Context, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ed25519Signer struct {
	key   ed25519.PrivateKey
	keyid string
}

func (s *ed25519Signer) KeyID() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s *ed25519Signer) Sign(ctx context.Context, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
