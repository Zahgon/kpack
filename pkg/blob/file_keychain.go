package blob

import (
	"fmt"
	"io/fs"
)

var errMultipleAuths = fmt.Errorf("only one of username/password, bearer, authorization is allowed")

type fileCredential struct {
	domain     string
	secretName string

	username      string
	password      string
	bearer        string
	authorization string
}

type fileCreds struct {
	creds []fileCredential
}

func NewMountedSecretBlobKeychain(volumeName string, secrets []string) (*fileCreds, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readFile(dirFs fs.FS, filename string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *fileCreds) Resolve(blobUrl string) (string, map[string]string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}
