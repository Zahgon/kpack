package slsa

import (
	"context"

	"github.com/google/go-containerregistry/pkg/authn"
	ggcrv1 "github.com/google/go-containerregistry/pkg/v1"
	intoto "github.com/in-toto/in-toto-golang/in_toto"
	"github.com/secure-systems-lab/go-securesystemslib/dsse"
)

type Signer = dsse.Signer

const (
	DssePayloadType   = "application/vnd.dsse.envelope.v1+json"
	IntotoPayloadType = "application/vnd.in-toto+json"
)

func (*Attester) Sign(ctx context.Context, stmt intoto.Statement, signers ...Signer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*Attester) Write(ctx context.Context, digestStr string, payload []byte, keychain authn.Keychain) (ggcrv1.Image, string, error) {
	_ = "STUB: not implemented"
	return *new(ggcrv1.Image), "", nil
}

// Overwrite any existing attestations with a new one. The only time this is
// relevant is when multiple builds result in bit-for-bit same images (since
// the digest would be the same in both builds).

// TODO: figure out how to determine if we should use the default docker media
// type. since all the secrets/signatures are combined into a single
// attestation image, it'll probably have to be on the Build resource
func scratchImage() ggcrv1.Image { _ = "STUB: not implemented"; return *new(ggcrv1.Image) }
