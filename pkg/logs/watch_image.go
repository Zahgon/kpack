package logs

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
	"github.com/pivotal/kpack/pkg/client/clientset/versioned"
)

type watchOneImage struct {
	kpackClient versioned.Interface
	image       *v1alpha1.Image
	ctx         context.Context
}

func (w watchOneImage) Watch(options v1.ListOptions) (watch.Interface, error) {
	_ = "STUB: not implemented"
	return *new(watch.Interface), nil
}
