package logs

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"

	"github.com/pivotal/kpack/pkg/client/clientset/versioned"
)

type watchOneBuild struct {
	buildName   string
	kpackClient versioned.Interface
	namespace   string
	context     context.Context
}

func (l *watchOneBuild) Watch(options v1.ListOptions) (watch.Interface, error) {
	_ = "STUB: not implemented"
	return *new(watch.Interface), nil
}

func (l *watchOneBuild) List(options v1.ListOptions) (runtime.Object, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Object), nil
}
