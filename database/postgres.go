package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Start - starts a new database connection with a PostgreSQL database from given envs
func Start(ctx context.Context) *gorm.DB {
	var err error

	if DB != nil {
		return DB
	}

	// read required envs
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	sslmode := os.Getenv("DB_SSLMODE")

	// read optional envs
	mode := os.Getenv("ENVIRONMENT_MODE")
	debugmode := strings.EqualFold(mode, "DEBUG")

	// envs validation
	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		log.Fatal("missing required envs")
	}

	// database connection credentials
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	if debugmode {
		log.Printf("Connecting to postgres database %s at %s:%s", dbname, host, port)
	}

	config := &gorm.Config{}
	if debugmode {
		config.Logger = logger.Default.LogMode(logger.Info)
	}

	// Create a new PostgreSQL database connection
	postgres, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	if debugmode {
		log.Printf("Connected to postgres database %s at %s:%s", dbname, host, port)
	}

	DB = postgres.WithContext(ctx)

	return DB
}
