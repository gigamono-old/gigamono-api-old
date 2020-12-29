package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectMySQLDB connects to MySQL database.
func ConnectMySQLDB() *DB {
	connectionString := os.Getenv("DB_CONNECTION_STRING")

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Panic("Cannot continue without a database.\n")
	}

	return &DB{connection: db}
}
