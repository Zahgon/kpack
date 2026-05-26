package flaghelpers

type CredentialsFlags []string

func (i *CredentialsFlags) String() string { _ = "STUB: not implemented"; return "" }

func (i *CredentialsFlags) Set(value string) error { _ = "STUB: not implemented"; return nil }

func GetEnvBool(key string, defaultValue bool) bool { _ = "STUB: not implemented"; return false }

func GetEnvInt(key string, defaultValue int) int { _ = "STUB: not implemented"; return 0 }
