package blob

import (
	"regexp"
)

var (
	azScope = "https://storage.azure.com/.default"
	azRegex = regexp.MustCompile(`.*[a-z0-9]+\.([a-z]+)\.core\.windows\.net\/.*`)
)

type azKeychain struct{}

func (a azKeychain) Resolve(url string) (string, map[string]string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// https://learn.microsoft.com/en-us/rest/api/storageservices/get-file
