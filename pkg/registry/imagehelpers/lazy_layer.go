package imagehelpers

import (
	"io"
	"sync"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type LazyMountableLayerArgs struct {
	Digest, DiffId, Image string
	Size                  int64
	Keychain              authn.Keychain
}

func NewLazyMountableLayer(args LazyMountableLayerArgs) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

type lazyMountableLayer struct {
	sync.Once
	keychain            authn.Keychain
	layer               v1.Layer
	fullyQualifiedLayer name.Digest
	digest              string
	diffId              string
	size                int64
}

func (m *lazyMountableLayer) Digest() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

func (m *lazyMountableLayer) DiffID() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

func (m *lazyMountableLayer) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (m lazyMountableLayer) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

func (m *lazyMountableLayer) Compressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (m *lazyMountableLayer) Uncompressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (m *lazyMountableLayer) fetchRemoteLayer() error { _ = "STUB: not implemented"; return nil }
