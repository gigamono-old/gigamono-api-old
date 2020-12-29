package helpers

import (
	"log"
)

// SetInitLogFile sets the file where initialization logs go.
func SetInitLogFile() {
	file, err := OpenFile("logs/init.log", true)
	if err != nil {
		log.Printf("Cannot open or create 'logs/init.log' file: %v\nFalling back to stdout/stderr\n", err)
	}

	log.SetOutput(file)
}
