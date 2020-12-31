package helpers

import (
	"log"
	"os"
)

// SetStatusLogFile sets the file where status logs go.
func SetStatusLogFile() {
	file, err := OpenLogFile("status.log")
	if err != nil {
		log.Printf("Cannot open or create 'logs/status.log' file: %v\nFalling back to stdout/stderr\n", err)
	}
	log.SetPrefix("\n")
	log.SetOutput(file)
}

// OpenLogFile opens/creates log file.
func OpenLogFile(dest string) (*os.File, error) {
	return OpenFile("logs/"+dest, true)
}
