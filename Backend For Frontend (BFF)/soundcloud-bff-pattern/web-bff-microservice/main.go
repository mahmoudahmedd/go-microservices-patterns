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
		c.JSON(http.StatusOK, gin.H{"message": "Hello from the web bff microservice!"})
	})

	router.GET("/web-bff-api/v1/tracks/:trackId", getWebTrack)

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}
}

// Handler function to get a track by id
func getWebTrack(c *gin.Context) {
	trackId := c.Param("trackId")
	var trackResponse = map[string]interface{}{
		"track":          fetchTrackById(trackId),
		"artwork":        fetchTrackArtwork(trackId, 500, 500),
		"related_tracks": fetchRelatedTracks(trackId),
		"playlists":      fetchPlaylists(trackId),
		"comments":       fetchComments(trackId),
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

func fetchRelatedTracks(trackId string) []models.RecommendedTrack {
	resp, _ := resty.New().R().Get("http://recommendation:8080/recommendation-api/v1/tracks/" + trackId + "/related")
	var recommendedTrack []models.RecommendedTrack
	_ = json.Unmarshal(resp.Body(), &recommendedTrack)
	return recommendedTrack
}

func fetchPlaylists(userId string) []models.Playlist {
	resp, _ := resty.New().R().Get("http://inventory:8080/inventory-api/v1/playlists/users/" + userId)
	var playlists []models.Playlist
	_ = json.Unmarshal(resp.Body(), &playlists)
	return playlists
}

func fetchComments(trackId string) []models.Comment {
	resp, _ := resty.New().R().Get("http://inventory:8080/inventory-api/v1/tracks/" + trackId + "/comments")
	var comments []models.Comment
	_ = json.Unmarshal(resp.Body(), &comments)
	return comments
}
