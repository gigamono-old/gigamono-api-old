package database

import (
	"errors"
	"log"
	"strings"

	"gorm.io/gorm"
)

// DBType represents the type of the database.
type DBType int

// DBType values.
const (
	POSTGRES DBType = iota
	MYSQL
	SQLITE
)

func (dbType DBType) String() string {
	var res string

	switch dbType {
	case POSTGRES:
		res = "POSTGRES"
	case MYSQL:
		res = "MYSQL"
	case SQLITE:
		res = "SQLITE"
	}

	return res
}

// ToDbType converts string representation to
func ToDbType(ty string) (DBType, error) {
	ty = strings.ToLower(ty)

	switch ty {
	case "postgres", "postgresql", "psql":
		return POSTGRES, nil
	case "mysql", "mysqldb":
		return MYSQL, nil
	case "sqlite":
		return SQLITE, nil
	default:
		return 0, errors.New("Unsupported database type")
	}
}

// DB contains a db connection.
type DB struct {
	connection *gorm.DB
	dbType     DBType
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
