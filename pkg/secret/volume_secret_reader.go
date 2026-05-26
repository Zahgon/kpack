package secret

func ReadBasicAuthSecret(secretVolume, secretName string) (BasicAuth, error) {
	_ = "STUB: not implemented"
	return *new(BasicAuth), nil
}

func ReadSshSecret(secretVolume, secretName string) (SSH, error) {
	_ = "STUB: not implemented"
	return *new(SSH), nil
}

func volumeName(VolumePath, secretName string) string { _ = "STUB: not implemented"; return "" }
