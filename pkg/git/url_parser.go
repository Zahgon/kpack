package git

import (
	"regexp"
)

var shortScpRegex = regexp.MustCompile(`^(ssh://)?(.*)@([[:alnum:]\.-]+):(.*)$`)

// ParseURL converts a short scp-like SSH syntax to a proper SSH URL.
// Git's ssh protocol supports a URL like user@hostname:path syntax, which is
// not a valid ssh url but is inherited from scp. Because the library we
// use for git relies on the Golang SSH support, we need to convert it to a
// proper SSH URL.
// See https://git-scm.com/book/en/v2/Git-on-the-Server-The-Protocols
func parseURL(url string) string { _ = "STUB: not implemented"; return "" }
