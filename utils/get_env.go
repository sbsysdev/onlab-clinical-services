package utils

import "os"

func GetEnv(key, defaultValue string) string {
	result, ok := os.LookupEnv(key)

	if !ok {
		return defaultValue
	}

	return result
}
