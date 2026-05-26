package logs

import (
	"context"
	"io"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
	"github.com/pivotal/kpack/pkg/client/clientset/versioned"

	"k8s.io/apimachinery/pkg/watch"
	watchTools "k8s.io/client-go/tools/watch"
)

type imageWaiter struct {
	KpackClient versioned.Interface
	logTailer   ImageLogTailer
}

type ImageLogTailer interface {
	TailBuildName(ctx context.Context, writer io.Writer, buildName, namespace string, timestamp bool) error
}

func NewImageWaiter(kpackClient versioned.Interface, logTailer ImageLogTailer) *imageWaiter {
	_ = "STUB: not implemented"
	return nil
}

func (w *imageWaiter) Wait(ctx context.Context, writer io.Writer, image *v1alpha1.Image) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func imageUpdateHasResolved(generation int64) func(event watch.Event) (bool, error) {
	_ = "STUB: not implemented"
	return nil
}

// image is reconciled

// image is resolved

// Build scheduled

// still waiting on build to be scheduled

// update skipped

// still waiting on update

func (w *imageWaiter) resultOfImageWait(ctx context.Context, writer io.Writer, generation int64, image *v1alpha1.Image) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func imageFailure(name, statusMessage string) error { _ = "STUB: not implemented"; return nil }

func (w *imageWaiter) waitBuild(ctx context.Context, writer io.Writer, namespace, buildName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// tail logs

func buildHasResolved(event watch.Event) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func buildFailure(statusMessage string) error { _ = "STUB: not implemented"; return nil }

func (w *imageWaiter) buildWatchUntil(ctx context.Context, namespace, buildName string, condition watchTools.ConditionFunc) (*v1alpha1.Build, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// event is nil if precondition is true

func filterErrors(condition watchTools.ConditionFunc) watchTools.ConditionFunc {
	_ = "STUB: not implemented"
	return *new(watchTools.ConditionFunc)
}

func (w *imageWaiter) imageBuildStarted(ctx context.Context, namespace, buildName string) error {
	_ = "STUB: not implemented"
	return nil
}
