package dockercreds

var registryDomains = []string{
	// Allow naked domains
	"%s",
	// Allow scheme-prefixed.
	"https://%s",
	"http://%s",
	// Allow scheme-prefixes with version in url path.
	"https://%s/v1/",
	"http://%s/v1/",
	"https://%s/v2/",
	"http://%s/v2/",
}

type RegistryMatcher struct {
	Registry string
}

func (m RegistryMatcher) Match(reg string) bool { _ = "STUB: not implemented"; return false }

func registryString(reg string) string { _ = "STUB: not implemented"; return "" }
