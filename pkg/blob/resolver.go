package blob

import (
	"context"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

type Resolver struct {
}

func (*Resolver) Resolve(ctx context.Context, sourceResolver *buildapi.SourceResolver) (corev1alpha1.ResolvedSourceConfig, error) {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.ResolvedSourceConfig), nil
}

func (*Resolver) CanResolve(sourceResolver *buildapi.SourceResolver) bool {
	_ = "STUB: not implemented"
	return false
}
