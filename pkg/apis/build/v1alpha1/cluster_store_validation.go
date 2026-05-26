package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
)

func (s *ClusterStore) SetDefaults(context.Context) { _ = "STUB: not implemented"; return }

func (s *ClusterStore) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (s *ClusterStoreSpec) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

//noinspection GoNilness
