package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectPostgresDB connects to Postgres database.
func ConnectPostgresDB() *DB {
	connectionString := os.Getenv("DB_CONNECTION_STRING")

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Panic("Cannot continue without a database.\n")
	}

	return &DB{connection: db}
}
