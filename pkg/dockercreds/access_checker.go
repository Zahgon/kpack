package dockercreds

import (
	"github.com/google/go-containerregistry/pkg/authn"
)

func VerifyWriteAccess(keychain authn.Keychain, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

func VerifyReadAccess(keychain authn.Keychain, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

func diagnoseIfTransportError(err error) error { _ = "STUB: not implemented"; return nil }

// transport.Error implements error to support the following error specification:
// https://github.com/docker/distribution/blob/master/docs/spec/api.md#errors

// handle artifactory. refer test case
