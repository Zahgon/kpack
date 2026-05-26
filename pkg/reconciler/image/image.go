package image

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	coreinformers "k8s.io/client-go/informers/core/v1"
	k8sclient "k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"knative.dev/pkg/controller"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	"github.com/pivotal/kpack/pkg/client/clientset/versioned"
	buildinformers "github.com/pivotal/kpack/pkg/client/informers/externalversions/build/v1alpha2"
	buildlisters "github.com/pivotal/kpack/pkg/client/listers/build/v1alpha2"
	"github.com/pivotal/kpack/pkg/duckbuilder"
	"github.com/pivotal/kpack/pkg/reconciler"
)

const (
	ReconcilerName = "Images"
	Kind           = "Image"
)

func NewController(
	ctx context.Context,
	opt reconciler.Options,
	k8sClient k8sclient.Interface,
	imageInformer buildinformers.ImageInformer,
	buildInformer buildinformers.BuildInformer,
	duckbuilderInformer *duckbuilder.DuckBuilderInformer,
	sourceResolverInformer buildinformers.SourceResolverInformer,
	pvcInformer coreinformers.PersistentVolumeClaimInformer,
	enablePriorityClasses bool,
) *controller.Impl {
	_ = "STUB: not implemented"
	return nil
}

type Reconciler struct {
	Client                versioned.Interface
	DuckBuilderLister     *duckbuilder.DuckBuilderLister
	ImageLister           buildlisters.ImageLister
	BuildLister           buildlisters.BuildLister
	SourceResolverLister  buildlisters.SourceResolverLister
	PvcLister             corelisters.PersistentVolumeClaimLister
	Tracker               reconciler.Tracker
	K8sClient             k8sclient.Interface
	EnablePriorityClasses bool
}

func (c *Reconciler) Reconcile(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Reconciler) reconcileImage(ctx context.Context, image *buildapi.Image) (*buildapi.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Reconciler) reconcileSourceResolver(ctx context.Context, image *buildapi.Image) (*buildapi.SourceResolver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Reconciler) reconcileBuildCache(ctx context.Context, image *buildapi.Image) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Reconciler) deleteOldBuilds(ctx context.Context, image *buildapi.Image) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Reconciler) fetchAllBuilds(image *buildapi.Image) (buildList, error) {
	_ = "STUB: not implemented"
	return *new(buildList), nil
}

func (c *Reconciler) fetchLastBuild(image *buildapi.Image) (*buildapi.Build, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Reconciler) updateStatus(ctx context.Context, desired *buildapi.Image) error {
	_ = "STUB: not implemented"
	return nil
}

func sourceResolversEqual(desiredSourceResolver *buildapi.SourceResolver, sourceResolver *buildapi.SourceResolver) bool {
	_ = "STUB: not implemented"
	return false
}

func buildCacheEqual(desiredBuildCache *corev1.PersistentVolumeClaim, buildCache *corev1.PersistentVolumeClaim) bool {
	_ = "STUB: not implemented"
	return false
}

func reconcilerKeyForBuilderKind(image *buildapi.Image) reconciler.Key {
	_ = "STUB: not implemented"
	return *new(reconciler.Key)
}
