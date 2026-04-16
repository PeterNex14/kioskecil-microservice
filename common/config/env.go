package config

import "os"

// GetEnv retrieves the value of the environment variable named by the key.
// It returns the value, which will be the fallback value if the variable is not present.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
