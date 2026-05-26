package buildchange

import (
	corev1 "k8s.io/api/core/v1"
	k8sclient "k8s.io/client-go/kubernetes"
)

const (
	ansiRegex         = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
	MaxLogMessageSize = 800
)

type ProgressLogger struct {
	K8sClient k8sclient.Interface
}

// GetTerminationMessage creates a termination message for a pod.
// Message consists out of container name that has terminated unsuccessfully with the
// last line of the container log truncated to 800 characters and the command to get more info from that container.
func (p *ProgressLogger) GetTerminationMessage(pod *corev1.Pod, s *corev1.ContainerStatus) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *ProgressLogger) getContainerLogs(pod *corev1.Pod, s *corev1.ContainerStatus) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func createMoreInfoCommand(namespace, podName, containerName string) string {
	_ = "STUB: not implemented"
	return ""
}

func stripAnsi(str string) string { _ = "STUB: not implemented"; return "" }

func truncate(text string) string { _ = "STUB: not implemented"; return "" }
