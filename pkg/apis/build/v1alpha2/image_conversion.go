package v1alpha2

import (
	"context"

	"knative.dev/pkg/apis"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
)

const (
	servicesConversionAnnotation              = "kpack.io/services"
	tolerationsConversionAnnotation           = "kpack.io/tolerations"
	nodeSelectorConversionAnnotation          = "kpack.io/nodeSelector"
	affinityConversionAnnotation              = "kpack.io/affinity"
	runtimeClassNameConversionAnnotation      = "kpack.io/runtimeClassName"
	schedulerNameConversionAnnotation         = "kpack.io/schedulerName"
	buildTimeoutConversionAnnotation          = "kpack.io/buildTimeout"
	storageClassNameConversionAnnotation      = "kpack.io/cache.volume.storageClassName"
	registryTagConversionAnnotation           = "kpack.io/cache.registry.tag"
	projectDescriptorPathConversionAnnotation = "kpack.io/projectDescriptorPath"
	cosignAnnotationConversionAnnotation      = "kpack.io/cosignAnnotation"
	defaultProcessConversionAnnotation        = "kpack.io/defaultProcess"
)

func (i *Image) ConvertTo(_ context.Context, to apis.Convertible) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *Image) ConvertFrom(_ context.Context, from apis.Convertible) error {
	_ = "STUB: not implemented"
	return nil
}

func (is *ImageSpec) convertTo(to *v1alpha1.ImageSpec) { _ = "STUB: not implemented"; return }

func (i *Image) convertFromAnnotations(fromAnnotations *map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (is *ImageSpec) convertToAnnotations(toAnnotations map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (is *ImageSpec) convertFrom(from *v1alpha1.ImageSpec) { _ = "STUB: not implemented"; return }

func (is *ImageStatus) convertFrom(from *v1alpha1.ImageStatus) { _ = "STUB: not implemented"; return }

func (is *ImageStatus) convertTo(to *v1alpha1.ImageStatus) { _ = "STUB: not implemented"; return }
