package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
)

func (s *ClusterStack) SetDefaults(context.Context) { _ = "STUB: not implemented"; return }

func (s *ClusterStack) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (ss *ClusterStackSpec) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (ssi *ClusterStackSpecImage) Validate(context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}
