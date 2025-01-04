package handlers

import (
	"assignment/config"
	"assignment/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type VideoResponse struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	PublishDatetime string `json:"publish_datetime"`
	ThumbnailURL    string `json:"thumbnail_url"`
	VideoID         string `json:"video_id"`
}

// GetVideos fetches paginated and sorted videos with selected fields only.
func GetVideos(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	var videos []models.Video
	var total int64

	// Fetch videos and count total
	config.Db.Model(&models.Video{}).Count(&total)
	config.Db.Order("publish_datetime DESC").Limit(limit).Offset(offset).Find(&videos)

	// Map to custom struct
	var videoResponses []VideoResponse
	for _, video := range videos {
		videoResponses = append(videoResponses, VideoResponse{
			Title:           video.Title,
			Description:     video.Description,
			PublishDatetime: video.PublishTime.Format(time.RFC3339), // Format datetime
			ThumbnailURL:    video.ThumbnailURL,
			VideoID:         video.VideoID,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"total":       total,
		"page":        page,
		"limit":       limit,
		"total_pages": (total + int64(limit) - 1) / int64(limit),
		"videos":      videoResponses,
	})
}

// SearchVideos fetches videos by title or description with selected fields only.
func SearchVideos(c *gin.Context) {
	titleQuery := c.Query("title")
	descriptionQuery := c.Query("description")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	var videos []models.Video
	var total int64

	query := config.Db.Model(&models.Video{})
	if titleQuery != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+titleQuery+"%")
	}
	if descriptionQuery != "" {
		query = query.Where("LOWER(description) LIKE ?", "%"+descriptionQuery+"%")
	}

	query.Count(&total).Order("publish_datetime DESC").Limit(limit).Offset(offset).Find(&videos)

	// Map to custom struct
	var videoResponses []VideoResponse
	for _, video := range videos {
		videoResponses = append(videoResponses, VideoResponse{
			Title:           video.Title,
			Description:     video.Description,
			PublishDatetime: video.PublishTime.Format(time.RFC3339),
			ThumbnailURL:    video.ThumbnailURL,
			VideoID:         video.VideoID,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"total":       total,
		"page":        page,
		"limit":       limit,
		"total_pages": (total + int64(limit) - 1) / int64(limit),
		"videos":      videoResponses,
	})
}
