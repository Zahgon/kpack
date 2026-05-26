package v1alpha2

import (
	"context"
	"regexp"

	authv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/apis"
)

const kpackControllerServiceAccountUsername = "system:serviceaccount:kpack:controller"

func (b *Build) SetDefaults(ctx context.Context) { _ = "STUB: not implemented"; return }

func (b *Build) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (bs *BuildSpec) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func resourceCreatedByKpackController(info *authv1.UserInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (bs *BuildSpec) validateImmutableFields(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (bs *BuildSpec) validateNodeSelector(_ context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func validateBuildEnvSecretKeyRefs(env []corev1.EnvVar) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (lb *LastBuild) Validate(context context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (c *BuildCacheConfig) Validate(context context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

var serviceNameRE = regexp.MustCompile(`^[a-z0-9\-\.]{1,253}$`)

func (ss Services) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

// check name uniqueness
