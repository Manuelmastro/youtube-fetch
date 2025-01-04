package config

import (
	"assignment/models"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres" // Import the PostgreSQL driver for GORM v2
	"gorm.io/gorm"
)

var (
	ApiKeys    []string
	CurrentKey int
	//Mu         sync.Mutex
	Db    *gorm.DB
	Query string
)

// loadEnv loads environment variables from a .env file
func LoadEnv() {
	// Load the environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the API keys from the environment and split them
	ApiKeys = strings.Split(os.Getenv("YOUTUBE_API_KEYS"), ",")
	Query = os.Getenv("QUERY")

	// Connect to the PostgreSQL database using GORM (with the correct dialect)
	dsn := os.Getenv("DB_CONNECTION") // Get the DB connection string from the environment
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Use postgres.Open(dsn) instead of a string
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Automatically migrate the schema (create the table if not already created)
	if err := Db.AutoMigrate(&models.Video{}); err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}

}
