package registryfakes

import (
	"github.com/google/go-containerregistry/pkg/authn"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

func NewFakeClient() *FakeClient { _ = "STUB: not implemented"; return nil }

type FakeClient struct {
	images        map[string]v1.Image
	readKeychains map[string]authn.Keychain

	savedImages    map[string]v1.Image
	writeKeychains map[string]authn.Keychain
	fetchError     error
}

func (f *FakeClient) Fetch(keychain authn.Keychain, repoName string) (v1.Image, string, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), "", nil
}

func (f *FakeClient) Save(keychain authn.Keychain, tag string, image v1.Image) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *FakeClient) AddImage(repoName string, image v1.Image, keychain authn.Keychain) {
	_ = "STUB: not implemented"
	return
}

func (f *FakeClient) AddSaveKeychain(tag string, keychain authn.Keychain) {
	_ = "STUB: not implemented"
	return
}

func (f *FakeClient) SavedImages() map[string]v1.Image { _ = "STUB: not implemented"; return nil }

func (f *FakeClient) SetFetchError(err error) { _ = "STUB: not implemented"; return }

func tryParsingTag(tag string) string { _ = "STUB: not implemented"; return "" }
