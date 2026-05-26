package v1alpha2

import (
	"context"

	"knative.dev/pkg/apis"
)

func (s *ClusterBuildpack) SetDefaults(context.Context) { _ = "STUB: not implemented"; return }

func (s *ClusterBuildpack) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (s *ClusterBuildpackSpec) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}
