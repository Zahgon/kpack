package slsa

import (
	"context"

	"github.com/secure-systems-lab/go-securesystemslib/dsse"
	"github.com/sigstore/sigstore/pkg/signature"
)

var _ dsse.Signer = (*cosignSigner)(nil)

type cosignSigner struct {
	signer signature.Signer
	keyid  string
}

// NewCosignSigner loads a cosign private key into a dsse signer. The main difference between this signer and the one
// provided by sigstore's dsse.WrappedSigner is that this signer doesn't compute the PAE when signing
func NewCosignSigner(key, pass []byte, id string) (*cosignSigner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// KeyID implements dsse.Signer.
func (s *cosignSigner) KeyID() (string, error) {
	_ = "STUB: not implemented"
	return "",

		// Sign implements dsse.Signer.
		nil
}

func (s *cosignSigner) Sign(ctx context.Context, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
