package dockercreds

func ParseDockerConfigSecret(dir string) (DockerCreds, error) {
	_ = "STUB: not implemented"
	return *new(DockerCreds), nil
}

func ParseBasicAuthSecrets(volumeName string, secrets []string) (DockerCreds, error) {
	_ = "STUB: not implemented"
	return *new(DockerCreds), nil
}

func parseDockerCfg(path string) (DockerCreds, error) {
	_ = "STUB: not implemented"
	return *new(DockerCreds), nil
}

func parseDockerConfigJson(path string) (DockerCreds, error) {
	_ = "STUB: not implemented"
	return *new(DockerCreds), nil
}

func fileExists(file string) (bool, error) { _ = "STUB: not implemented"; return false, nil }
