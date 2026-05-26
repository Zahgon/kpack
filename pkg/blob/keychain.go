package blob

type Keychain interface {
	Resolve(url string) (authHeader string, headers map[string]string, err error)
}

var DefaultKeychain = NewMultiKeychain(
	azKeychain{},
	gcpKeychain{},
)

type multiKeychain struct {
	keychains []Keychain
}

func NewMultiKeychain(creds ...Keychain) Keychain { _ = "STUB: not implemented"; return *new(Keychain) }

func (m *multiKeychain) Resolve(url string) (string, map[string]string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}
