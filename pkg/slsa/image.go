package slsa

import (
	"github.com/google/go-containerregistry/pkg/authn"
	ggcrv1 "github.com/google/go-containerregistry/pkg/v1"
	slsacommon "github.com/in-toto/in-toto-golang/in_toto/slsa_provenance/common"
)

const (
	ProjectMetadataLabel = "io.buildpacks.project.metadata"
)

type ImageFetcher interface {
	Fetch(keychain authn.Keychain, repoName string) (ggcrv1.Image, string, error)
}

type reader struct {
	fetcher ImageFetcher
}

func NewImageReader(fetcher ImageFetcher) *reader { _ = "STUB: not implemented"; return nil }

func (r *reader) Read(keychain authn.Keychain, repoName string) (string, string, map[string]string, error) {
	_ = "STUB: not implemented"
	return "", "", nil, nil
}

func extractSourceFromLabel(labels map[string]string) (string, slsacommon.DigestSet, error) {
	_ = "STUB: not implemented"
	return "", *new(slsacommon.DigestSet), nil
}

// while sha256 support is available, go-git still defaults to sha1 for now
// https://github.com/go-git/go-git/issues/706

type project struct {
	Source source `json:"source"`
}

type source struct {
	Type     string   `json:"type"`
	Metadata metadata `json:"metadata"`
	Version  version  `json:"version"`
}

type metadata struct {
	Repository string `json:"repository"`
	Revision   string `json:"revision"`
	Image      string `json:"image"`
	Url        string `json:"url"`
}

type version struct {
	Commit string `json:"commit"`
	Digest string `json:"digest"`
	SHA256 string `json:"sha256sum"`
}
