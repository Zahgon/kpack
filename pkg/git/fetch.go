package git

import (
	"log"

	"github.com/go-git/go-git/v5/plumbing/protocol/packp/capability"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

type Fetcher struct {
	Logger               *log.Logger
	Keychain             GitKeychain
	InitializeSubmodules bool
}

func init() {
	//remove multi_ack and multi_ack_detailed from unsupported capabilities to enable Azure DevOps git support
	transport.UnsupportedCapabilities = []capability.Capability{
		capability.ThinPack,
	}
}

func (f Fetcher) Fetch(dir, gitURL, gitRevision, metadataDir string) error {
	_ = "STUB: not implemented"
	return nil
}

//resolvedSourceConfig.Git.Revision is the hash of the commit

type Project struct {
	Source Source `toml:"source"`
}

type Source struct {
	Type     string   `toml:"type"`
	Metadata Metadata `toml:"metadata"`
	Version  Version  `toml:"version"`
}

type Metadata struct {
	Repository string `toml:"repository"`
	Revision   string `toml:"revision"`
}

type Version struct {
	Commit string `toml:"commit"`
}
