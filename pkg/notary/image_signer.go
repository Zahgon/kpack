package notary

import (
	"log"

	"github.com/buildpacks/lifecycle/platform/files"
	"github.com/google/go-containerregistry/pkg/authn"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/theupdateframework/notary/client"
	"github.com/theupdateframework/notary/cryptoservice"
	"github.com/theupdateframework/notary/storage"
	"github.com/theupdateframework/notary/tuf/data"
	"github.com/theupdateframework/notary/tuf/signed"
)

type ImageFetcher interface {
	Fetch(keychain authn.Keychain, repoName string) (v1.Image, string, error)
}

type RepositoryFactory interface {
	GetRepository(url string, gun data.GUN, remoteStore storage.RemoteStore, cryptoService signed.CryptoService) (Repository, error)
}

type Repository interface {
	PublishTarget(target *client.Target) error
}

type ImageSigner struct {
	Logger  *log.Logger
	Client  ImageFetcher
	Factory RepositoryFactory
}

func (s *ImageSigner) Sign(url, notarySecretDir string, report files.Report, keychain authn.Keychain) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ImageSigner) makeGUNAndTargets(report files.Report, keychain authn.Keychain) (data.GUN, []*client.Target, error) {
	_ = "STUB: not implemented"
	return *new(data.GUN), nil, nil
}

func (s *ImageSigner) makeCryptoService(notarySecretDir string) (*cryptoservice.CryptoService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func k8sSecretPassRetriever(notarySecretDir string) func(_, _ string, _ bool, _ int) (passphrase string, giveup bool, err error) {
	_ = "STUB: not implemented"
	return nil
}
