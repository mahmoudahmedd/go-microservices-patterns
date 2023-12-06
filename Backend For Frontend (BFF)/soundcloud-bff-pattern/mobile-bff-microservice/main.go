package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
	"os"
	"repository-microservice/models"
	"strconv"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from the mobile bff microservice!"})
	})

	router.GET("/mobile-bff-api/v1/tracks/:trackId", getMobileTrack)

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}
}

// Handler function to get a track by id
func getMobileTrack(c *gin.Context) {
	trackId := c.Param("trackId")
	var trackResponse = map[string]interface{}{
		"track":   fetchTrackById(trackId),
		"artwork": fetchTrackArtwork(trackId, 200, 200),
		"media":   fetchTrackMedia(trackId),
	}

	c.JSON(http.StatusOK, trackResponse)
}

func fetchTrackById(trackId string) models.Track {
	resp, _ := resty.New().R().Get("http://inventory:8080/inventory-api/v1/tracks/" + trackId)
	var track models.Track
	_ = json.Unmarshal(resp.Body(), &track)
	return track
}

func fetchTrackArtwork(trackId string, width, height int) models.Artwork {
	resp, _ := resty.New().R().Get("http://file-storage:8080/file-storage-api/v1/tracks/" + trackId + "/artwork?width=" + strconv.Itoa(width) + "&height=" + strconv.Itoa(height))
	var artwork models.Artwork
	_ = json.Unmarshal(resp.Body(), &artwork)
	return artwork
}

func fetchTrackMedia(trackId string) models.Media {
	resp, _ := resty.New().R().Get("http://file-storage:8080/file-storage-api/v1/tracks/" + trackId + "/files")
	var media models.Media
	_ = json.Unmarshal(resp.Body(), &media)
	return media
}
