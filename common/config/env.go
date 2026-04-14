package config

import (
	"fmt"
	"os"
)

// GetEnv retrieves the value of the environment variable named by the key.
// It returns the value, which will be the fallback value if the variable is not present.
func GetEnv(key string) (string, error) {
	if value, ok := os.LookupEnv(key); ok {
		return value, nil
	}
	return "", fmt.Errorf("Terjadi kesalahan: File .env tidak ditemukan")
}
