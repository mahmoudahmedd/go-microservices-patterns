package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"repository-microservice/models"
	"time"
)

var recommendedTracks []models.RecommendedTrack

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	seed()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from the recommendation microservice!"})
	})

	router.GET("/recommendation-api/v1/tracks/:trackId/related", getRecommendedTracks)

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}
}

func getRecommendedTracks(c *gin.Context) {
	_ = c.Param("trackId")

	// logic to get recommended tracks
	time.Sleep(time.Second * 2)

	c.JSON(http.StatusOK, recommendedTracks)
}

func seed() {
	recommendedTracks = append(recommendedTracks, models.NewRecommendedTrack("777", "Song 777", "Artist 1"))
	recommendedTracks = append(recommendedTracks, models.NewRecommendedTrack("888", "Song 888", "Artist 2"))
	recommendedTracks = append(recommendedTracks, models.NewRecommendedTrack("999", "Song 999", "Artist 3"))
}
