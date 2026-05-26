package v1alpha1

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"knative.dev/pkg/apis"
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

func (is *ImageSpec) validateCacheSize(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (ib *ImageBuild) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func validateBuilder(builder v1.ObjectReference) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}
