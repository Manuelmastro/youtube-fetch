package youtubeapi

import (
	"log"
	"time"
)

func BackgroundProcess() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	// Use for range to receive ticks from the ticker channel
	for range ticker.C {
		videos, err := fetchYouTubeData()
		if err != nil {
			log.Println("Error fetching data:", err)
			continue
		}
		storeData(videos)
	}
}
