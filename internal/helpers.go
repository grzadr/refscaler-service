package internal

import (
	"os"
)

func DefaultEnv(key, value string) string {
	env, found := os.LookupEnv(key)

	if !found {
		return value
	}

	return env
}
