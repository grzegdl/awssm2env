package env

import (
	"os"
	"strings"
)

// SetEnvs sets environments from string map
func SetEnvs(envs map[string]string) error {
	for k, v := range envs {
		err := os.Setenv(k, v)

		if err != nil {
			return err
		}
	}

	return nil
}

// Envs2Map reads string slice with format "key=value" as key/val map
func Envs2Map(e []string) map[string]string {
	envs := make(map[string]string)

	for _, env := range e {
		kv := strings.SplitN(env, "=", 2)
		if len(kv) > 1 {
			envs[kv[0]] = kv[1]
		}
	}

	return envs
}
