package v1alpha2

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"knative.dev/pkg/apis"

	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

type ImageContextKey string

const (
	HasDefaultStorageClass ImageContextKey = "hasDefaultStorageClass"
	IsExpandable           ImageContextKey = "isExpandable"
)

var (
	defaultFailedBuildHistoryLimit     int64 = 10
	defaultSuccessfulBuildHistoryLimit int64 = 10
	defaultCacheSize                   resource.Quantity
)

func init() {
	defaultCacheSize = resource.MustParse("2G")
}

func (i *Image) SetDefaults(ctx context.Context) { _ = "STUB: not implemented"; return }

func (i *Image) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (i *Image) ValidateMetadata(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (i *Image) validateName(imageName string) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (is *ImageSpec) ValidateSpec(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (is *ImageSpec) validateTag(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (is *ImageSpec) validateAdditionalTags(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (is *ImageSpec) validateSameRegistry() *apis.FieldError { _ = "STUB: not implemented"; return nil }

// We only care about the non-nil error cases here as we validate
// the tag validity in other methods which should display appropriate errors.

func (is *ImageSpec) validateVolumeCache(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (ib *ImageBuild) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

// check that the timestamp in CreationTime is in valid format

func validateBuilder(builder v1.ObjectReference) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (is *ImageSpec) validateBuildHistoryLimit() *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (c *ImageCacheConfig) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func validateNotary(ctx context.Context, config *corev1alpha1.NotaryConfig) *apis.FieldError {
	_ = "STUB: not implemented"
	//only allow the kpack controller to create resources with notary
	return nil
}

func validateCnbBindings(ctx context.Context, bindings corev1alpha1.CNBBindings) *apis.FieldError {
	_ = "STUB: not implemented"
	//only allow the kpack controller to create resources with cnb bindings
	return nil
}
