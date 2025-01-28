package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    // Get the database URL from environment variable
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        log.Fatal("DATABASE_URL environment variable is not set")
    }

    var err error
    DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
    if err != nil {
        log.Fatal("Error connecting to database: ", err)
    }
    log.Println("Connection established successfully")
}