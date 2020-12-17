package utils

import (
	"fmt"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadEnvFile loads .env file.
func LoadEnvFile() {
	path, err := filepath.Abs(".env")
	if err != nil {
		return
	}

	fmt.Printf("dir [loadEnv] (%v)\n", path)

	err = godotenv.Load(path)
	if err != nil {
		return
	}
}
