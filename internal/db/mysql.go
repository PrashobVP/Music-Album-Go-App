package db

import (
	"fmt"
	"os"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect initializes the MySQL database connection using GORM.
func Connect() (*gorm.DB, error) {
	// Get database configuration from environment variables with optional defaults.
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Log a fatal error if any of the critical environment variables are missing.
	if dbUser == "" || dbPass == "" || dbHost == "" || dbName == "" {
		log.Fatal("Missing one or more required database environment variables.")
	}

	// Format the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		dbUser, dbPass, dbHost, dbName)

	// Connect to the MySQL database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

