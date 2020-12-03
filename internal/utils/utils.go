package utils

import (
	"fmt"
	"os"
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadEnvFile loads .env file.
func LoadEnvFile() {
	path, err := filepath.Abs(".env")
	if err != nil {
		log.Fatalf("Unable to resolve path: Error: %v", err)
	}

	fmt.Printf("dir [loadEnv] (%v)\n", path)

	err = godotenv.Load(path)
	if err != nil {
		log.Fatalf("Unable to load .env file: Error: %v", err)
	}

	fmt.Printf("Mongdb URI [loadEnv] (%v)\n", os.Getenv("MONGODB_URI"))
}

