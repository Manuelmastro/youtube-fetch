package main

import (
	"assignment/config"
	"assignment/handlers"
	youtubeapi "assignment/yt_api"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables and initialize the database
	config.LoadEnv()

	// Start the background process for fetching YouTube data
	go youtubeapi.BackgroundProcess() // No need for an additional timer here

	// Create a Gin router
	r := gin.Default()

	// Define routes
	r.GET("/videos", handlers.GetVideos)           // GET API for fetching videos
	r.GET("/videos/search", handlers.SearchVideos) // Search API

	// Start the server
	log.Println("Server is running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
