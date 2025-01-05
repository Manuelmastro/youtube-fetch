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
		log.Println("Error loading .env file")
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
	setupTSVector(Db)
}

// setupTSVector ensures the TSVector column is set up and indexed with triggers
func setupTSVector(db *gorm.DB) {
	// Add the search_vector column
	db.Exec(`
		ALTER TABLE videos
		ADD COLUMN IF NOT EXISTS search_vector tsvector
		GENERATED ALWAYS AS (
			to_tsvector('english', COALESCE(title, '') || ' ' || COALESCE(description, ''))
		) STORED;
	`)

	// Create a GIN index for the search_vector column
	db.Exec(`
		CREATE INDEX IF NOT EXISTS idx_videos_search_vector
		ON videos
		USING GIN (search_vector);
	`)

	// Add a trigger to update search_vector on title or description changes
	db.Exec(`
		CREATE OR REPLACE FUNCTION update_search_vector()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.search_vector := to_tsvector('english', COALESCE(NEW.title, '') || ' ' || COALESCE(NEW.description, ''));
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;
	`)

	db.Exec(`
		CREATE TRIGGER trigger_update_search_vector
		BEFORE INSERT OR UPDATE ON videos
		FOR EACH ROW EXECUTE FUNCTION update_search_vector();
	`)
}
