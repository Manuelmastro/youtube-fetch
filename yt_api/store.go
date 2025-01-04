package youtubeapi

import (
	"assignment/config"
	"assignment/models"
	"log"

	"gorm.io/gorm/clause"
)

// Store data in the database
func storeData(videos []models.Video) {
	for _, video := range videos {
		// Use GORM's Create or Save method with OnConflict to handle "upsert"
		if err := config.Db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "video_id"}}, // Specify the unique constraint column(s)
			DoNothing: true,                                // Do nothing if the record already exists
		}).Create(&video).Error; err != nil {
			log.Println("Failed to insert video:", err)
		}
	}
}
