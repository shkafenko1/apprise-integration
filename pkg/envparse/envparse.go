package envparse

import "os"

func GetEnv(key, fallback string) string {
	val := os.Getenv(key)
	if fallback == "required" && val == "" {
		panic("Required environment variable is missing: " + key)
	}
	if val != "" {
		return val
	}
	return fallback
}
