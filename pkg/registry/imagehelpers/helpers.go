package imagehelpers

import (
	"time"

	v1 "github.com/google/go-containerregistry/pkg/v1"
)

func GetCreatedAt(image v1.Image) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func GetEnv(image v1.Image, key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func SetEnv(image v1.Image, key, value string) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

func HasLabel(image v1.Image, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func GetStringLabel(image v1.Image, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetLabel(image v1.Image, key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func SetStringLabel(image v1.Image, key, value string) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

func SetStringLabels(image v1.Image, labels map[string]string) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

func SetLabels(image v1.Image, labels map[string]interface{}) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

func SetWorkingDir(image v1.Image, dir string) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

func GetWorkingDir(image v1.Image) (string, error) { _ = "STUB: not implemented"; return "", nil }

func configFile(image v1.Image) (*v1.ConfigFile, error) { _ = "STUB: not implemented"; return nil, nil }
