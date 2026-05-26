package notary

import (
	"github.com/theupdateframework/notary/client"
	"github.com/theupdateframework/notary/storage"
	"github.com/theupdateframework/notary/tuf/data"
	"github.com/theupdateframework/notary/tuf/signed"
)

type RemoteRepositoryFactory struct {
}

func (r *RemoteRepositoryFactory) GetRepository(url string, gun data.GUN, remoteStore storage.RemoteStore, cryptoService signed.CryptoService) (Repository, error) {
	_ = "STUB: not implemented"
	return *new(Repository), nil
}

type RemoteRepository struct {
	repo client.Repository
}

func (r *RemoteRepository) PublishTarget(target *client.Target) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RemoteRepository) getRoles(target *client.Target) ([]data.RoleName, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
