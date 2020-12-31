package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectSQLiteDB connects to SQLite database.
func ConnectSQLiteDB(connectionString string, newLogger *logger.Interface) *gorm.DB {
	db, err := gorm.Open(
		sqlite.Open(connectionString),
		&gorm.Config{
			Logger: *newLogger,
		},
	)

	if err != nil {
		log.Panic("Cannot continue without a database.\n")
	}

	return db
}
