package youtubeapi

import (
	"assignment/config"
	"assignment/models"
	"log"
	"time"

	"gorm.io/gorm/clause"
)

// Store data in the database
func storeData(videos []models.Video) {
	for _, video := range videos {
		// Log the video insertion attempt
		log.Printf("%s Info: Inserting video with VideoID: %s into the database", time.Now().Format(time.RFC3339), video.VideoID)

		// Set search_vector to an empty string, so it won't be inserted manually
		video.SearchVector = ""

		// Use GORM's Create or Save method with OnConflict to handle "upsert"
		if err := config.Db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "video_id"}}, // Specify the unique constraint column(s)
			DoNothing: true,                                // Do nothing if the record already exists
		}).Create(&video).Error; err != nil {
			log.Printf("%s Error: Failed to insert video with VideoID: %s, error: %v", time.Now().Format(time.RFC3339), video.VideoID, err)
		} else {
			log.Printf("%s Info: Video inserted into database with VideoID: %s", time.Now().Format(time.RFC3339), video.VideoID)
		}
	}
}
