package envparse

import "os"

func GetEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	if fallback == "" {
		panic("Required environment variable is missing: " + key)
	}
	return fallback
}
