package logs

import (
	"context"
	"io"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sclient "k8s.io/client-go/kubernetes"
)

type BuildLogsClient struct {
	k8sClient k8sclient.Interface
	processed map[readyContainer]interface{}
}

func NewBuildLogsClient(k8sClient k8sclient.Interface) *BuildLogsClient {
	_ = "STUB: not implemented"
	return nil
}

func (c *BuildLogsClient) Tail(ctx context.Context, writer io.Writer, image, build, namespace string, timestamp bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *BuildLogsClient) TailImage(ctx context.Context, writer io.Writer, image, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *BuildLogsClient) GetImageLogs(ctx context.Context, writer io.Writer, image, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *BuildLogsClient) TailBuildName(ctx context.Context, writer io.Writer, namespace string, buildName string, timestamp bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *BuildLogsClient) tailPods(ctx context.Context, writer io.Writer, namespace string, listOptions metav1.ListOptions, exitPodComplete bool, follow, timestamp bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *BuildLogsClient) getPodLogs(ctx context.Context, writer io.Writer, namespace string, listOptions metav1.ListOptions, follow, timestamp bool) error {
	_ = "STUB: not implemented"
	return nil
}

type readyContainer struct {
	podName       string
	containerName string
	namespace     string
}

func (c *BuildLogsClient) watchReadyContainers(ctx context.Context, readyContainers chan<- readyContainer, namespace string, listOptions metav1.ListOptions, exitPodComplete bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *BuildLogsClient) getContainers(ctx context.Context, namespace string, listOptions metav1.ListOptions) ([]readyContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func finished(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func (c *BuildLogsClient) streamLogsForContainer(ctx context.Context, writer io.Writer, readyContainer readyContainer, follow, timestamp bool) error {
	_ = "STUB: not implemented"
	return nil
}

func cyan(s string) string { _ = "STUB: not implemented"; return "" }
