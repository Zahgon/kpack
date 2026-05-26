package blob

import (
	"fmt"
	"io"
	"log"
	"os"
)

var errUnexpectedBlobType = fmt.Errorf("unexpected blob file type, must be one of .zip, .tar.gz, .tar, .jar")

type Fetcher struct {
	Logger   *log.Logger
	Keychain Keychain
}

func (f *Fetcher) Fetch(dir string, blobURL string, stripComponents int, metadataDir string) error {
	_ = "STUB: not implemented"
	return nil
}

func downloadBlob(blobURL string, headers map[string]string) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func classifyFile(reader io.ReadSeeker) (string, error) { _ = "STUB: not implemented"; return "", nil }

func sha256sum(reader io.ReadSeeker) (string, error) { _ = "STUB: not implemented"; return "", nil }

type Project struct {
	Source Source `toml:"source"`
}

type Source struct {
	Type     string   `toml:"type"`
	Metadata Metadata `toml:"metadata"`
	Version  Version  `toml:"version"`
}

type Metadata struct {
	Url string `toml:"url"`
}

type Version struct {
	SHA256 string `toml:"sha256sum"`
}
