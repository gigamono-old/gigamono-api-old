package database

import (
	"log"
	"os"
	"time"

	"github.com/sageflow/sageflow-api/internal/database/models"
	"github.com/sageflow/sageflow-api/internal/helpers"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Setup connects to and prepares the database.
func Setup() *DB {
	db := Connect()             // Connects to database
	db.AddMissingTableColumns() // Add expected tables and columns. This may be bad.
	return db
}

// Connect connects to the appropriate database.
func Connect() *DB {
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	newLogger := createStatusLogger()

	dbType, err := ToDbType(os.Getenv("DB_TYPE"))
	if err != nil {
		log.Panicf("%v\n", err)
	}

	var db *gorm.DB

	switch dbType {
	case POSTGRES:
		db = ConnectPostgresDB(connectionString, newLogger)
	case MYSQL:
		db = ConnectMySQLDB(connectionString, newLogger)
	case SQLITE:
		db = ConnectSQLiteDB(connectionString, newLogger)
	default:
		log.Panic("Unsupported database type\n")
	}

	log.Printf("Database successfully connected: %v\n", dbType.String())

	return &DB{
		connection: db,
		dbType:     dbType,
	}
}

// AddMissingTableColumns adds missing tables and columns.
func (db *DB) AddMissingTableColumns() {
	db.connection.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.SocialLogin{},
		&models.Workspace{},
		&models.Group{},
		&models.AccessControl{},
		&models.Role{},
		&models.Folder{},
		&models.Workflow{},
		&models.Account{},
		&models.AuthInfo{},
		&models.App{},
		&models.Theme{},
	)

	log.Println("Initial migrations successful")
}

func createStatusLogger() *logger.Interface {
	file, err := helpers.OpenLogFile("status.log")
	if err != nil {
		log.Printf("Cannot open or create 'logs/status.log' file: %v\nFalling back to stdout/stderr\n", err)
	}

	newLogger := logger.New(
		log.New(file, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Warn, // Log level
			Colorful:      false,       // Disable color
		},
	)

	return &newLogger
}
