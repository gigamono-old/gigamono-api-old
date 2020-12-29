package database

import (
	"log"
	"os"
	"strings"

	"gorm.io/gorm"
)

// Connect to the appropriate database.
func Connect() *DB {
	dbType := strings.ToLower(os.Getenv("DB_TYPE"))

	var db *DB

	switch dbType {
	case "postgres", "postgresql", "psql":
		db = ConnectPostgresDB()
	case "mysql", "mysqldb":
		db = ConnectMySQLDB()
	case "sqlite":
		db = ConnectSQLiteDB()
	default:
		log.Panic("Unsupported database type\n")
	}

	return db
}

// DB contains a db connection.
type DB struct {
	connection *gorm.DB
}

// Close closes database connection.
func (db *DB) Close() {
	specificDB, err := db.connection.DB()
	if err != nil {
		log.Printf("Cannot fetch database for closure: %v\n", err)
	}

	log.Println("Closing database")

	specificDB.Close()

	log.Println("Closed database")
}
