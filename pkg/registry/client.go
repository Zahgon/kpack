package registry

import (
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

type Client struct {
}

func (t *Client) Fetch(keychain authn.Keychain, repoName string) (v1.Image, string, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), "", nil
}

func (t *Client) Save(keychain authn.Keychain, tag string, image v1.Image) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func timestampTag() string { _ = "STUB: not implemented"; return "" }

func getIdentifier(image v1.Image, ref name.Reference) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func previousDigest(keychain authn.Keychain, ref name.Reference) string {
	_ = "STUB: not implemented"
	return ""
}

func handleError(err error) error { _ = "STUB: not implemented"; return nil }
