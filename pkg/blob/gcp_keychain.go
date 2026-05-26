package blob

const (
	gcpScope = "https://www.googleapis.com/auth/devstorage.read_only"
)

type gcpKeychain struct{}

func (g gcpKeychain) Resolve(url string) (string, map[string]string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}
