package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnvFile loads .env file.
func LoadEnvFile() {
	err := godotenv.Load() // Loads .env file
	if err != nil {
		log.Printf("Unable load environment variables: %v\n", err)
	}
}
