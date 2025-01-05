package youtubeapi

import (
	"log"
	"time"
)

// function to make fetchYouTubeData() function call at 1 minute rime interval
func BackgroundProcess() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		videos, err := fetchYouTubeData()
		if err != nil {
			log.Println("Error fetching data:", err)
			continue
		}
		storeData(videos)
	}
}
