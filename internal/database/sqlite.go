package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectSQLiteDB connects to SQLite database.
func ConnectSQLiteDB() *DB {
	connectionString := os.Getenv("DB_CONNECTION_STRING")

	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Panic("Cannot continue without a database.\n")
	}

	return &DB{connection: db}
}
