package registry

import (
	"log"
	"os"

	"github.com/google/go-containerregistry/pkg/authn"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

type contentType string

const (
	ContentTypeLabelKey string = "source.contenttype.kpack.io"

	zip   contentType = "zip"
	jar   contentType = "jar"
	war   contentType = "war"
	tar   contentType = "tar"
	targz contentType = "tar.gz"
)

type ImageClient interface {
	Fetch(keychain authn.Keychain, repoName string) (v1.Image, string, error)
}

type Fetcher struct {
	Logger   *log.Logger
	Client   ImageClient
	Keychain authn.Keychain
}

func (f *Fetcher) Fetch(dir, registryImage string, metadataDir string) error {
	_ = "STUB: not implemented"
	return nil
}

func getContentType(img v1.Image) (contentType, error) {
	_ = "STUB: not implemented"
	return *new(contentType), nil
}

func handleSource(img v1.Image, dir string) error { _ = "STUB: not implemented"; return nil }

func handleZip(img v1.Image, dir string) error { _ = "STUB: not implemented"; return nil }

func handleTar(img v1.Image, dir string) error { _ = "STUB: not implemented"; return nil }

func handleTarGZ(img v1.Image, dir string) error { _ = "STUB: not implemented"; return nil }

func getSourceFile(img v1.Image, dir string) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fetchLayer(layer v1.Layer, dir string) error { _ = "STUB: not implemented"; return nil }

type Project struct {
	Source Source `toml:"source"`
}

type Source struct {
	Type     string   `toml:"type"`
	Metadata Metadata `toml:"metadata"`
	Version  Version  `toml:"version"`
}

type Metadata struct {
	Image string `toml:"image"`
}

type Version struct {
	Digest string `toml:"digest"`
}
