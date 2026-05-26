package cnb

func serializeEnvVars(envVars []envVariable, platformDir string) error {
	_ = "STUB: not implemented"
	return nil
}

type envVariable struct {
	Name  string `json:"name" toml:"name"`
	Value string `json:"value" toml:"value"`
}

type buildEnvVariable struct {
	Env []envVariable `toml:"env"`
}

type envBuildVariable struct {
	Env []envVariable
}

func (a *envBuildVariable) UnmarshalTOML(f interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func buildEnv(v []map[string]interface{}) ([]envVariable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
