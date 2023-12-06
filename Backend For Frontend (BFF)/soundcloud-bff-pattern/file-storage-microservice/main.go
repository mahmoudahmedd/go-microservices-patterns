package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// Media struct to represent a track
type Media struct {
	TrackId string            `json:"track_id"`
	Files   map[string]string `json:"download_links"`
}

// Artwork struct to represent artwork information
type Artwork struct {
	URL    string `json:"url"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from the file storage microservice!"})
	})

	router.GET("/file-storage-api/v1/tracks/:trackId/files", getTrackFiles)
	router.GET("/file-storage-api/v1/tracks/:trackId/artwork", getTrackArtwork)

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}
}

func getTrackArtwork(c *gin.Context) {
	trackID := c.Param("trackId")
	width := c.Query("width")
	height := c.Query("height")

	artworkUrl := fetchTrackArtwork(trackID, width, height)

	c.JSON(http.StatusOK, Artwork{
		Height: height,
		Width:  width,
		URL:    artworkUrl,
	})
}

func getTrackFiles(c *gin.Context) {
	trackID := c.Param("trackId")

	fileStorage := fetchTrackFiles(trackID)

	c.JSON(http.StatusOK, Media{
		TrackId: trackID,
		Files:   fileStorage,
	})
}

func fetchTrackFiles(trackID string) map[string]string {
	files := map[string]string{
		"mp3":  "https://soundcloud.s3.aws.com/" + trackID + "/file.mp3",
		"flac": "https://soundcloud.s3.aws.com/" + trackID + "/file.flac",
	}

	return files
}

func fetchTrackArtwork(trackID, width, height string) string {
	return "https://i1.sndcdn.com/artworks-" + trackID + "-t" + width + "x" + height + ".jpg"
}
