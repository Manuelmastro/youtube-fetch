package youtubeapi

import (
	"assignment/config"
	"assignment/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func rotateAPIKey() string {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	config.CurrentKey = (config.CurrentKey + 1) % len(config.ApiKeys)
	return config.ApiKeys[config.CurrentKey]
}

func fetchYouTubeData() ([]models.Video, error) {
	var allVideos []models.Video // To store all videos fetched
	apiKey := config.ApiKeys[config.CurrentKey]

	// Build the API URL with order=date to sort by most recent
	url := fmt.Sprintf(
		"https://www.googleapis.com/youtube/v3/search?part=snippet&q=%s&type=video&maxResults=50&order=date&key=%s",
		config.Query, apiKey)

	log.Printf("%s Info: Fetching data from YouTube API with query: %s", time.Now().Format(time.RFC3339), config.Query)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("API call failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 403 {
		apiKey = rotateAPIKey()
		log.Println("Quota exhausted, switched to next API key:", apiKey)
		return nil, fmt.Errorf("quota exhausted")
	}

	var data struct {
		Items []struct {
			ID struct {
				VideoID string `json:"videoId"`
			} `json:"id"`
			Snippet struct {
				Title       string `json:"title"`
				Description string `json:"description"`
				PublishedAt string `json:"publishedAt"`
				Thumbnails  struct {
					Default struct {
						URL string `json:"url"`
					} `json:"default"`
				} `json:"thumbnails"`
			} `json:"snippet"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	// Process the items and append them to the allVideos slice
	for _, item := range data.Items {
		publishTime, _ := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		allVideos = append(allVideos, models.Video{
			Title:        item.Snippet.Title,
			Description:  item.Snippet.Description,
			PublishTime:  publishTime,
			ThumbnailURL: item.Snippet.Thumbnails.Default.URL,
			VideoID:      item.ID.VideoID,
		})
	}

	log.Printf("%s Info: Fetched %d videos from YouTube API", time.Now().Format(time.RFC3339), len(data.Items))

	return allVideos, nil
}
