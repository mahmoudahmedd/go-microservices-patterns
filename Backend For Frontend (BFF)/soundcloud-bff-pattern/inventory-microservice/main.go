package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"repository-microservice/models"
	"time"
)

var playlists []models.Playlist
var comments []models.Comment
var tracks []models.Track

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	seed()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from the inventory microservice!"})
	})

	router.GET("/inventory-api/v1/tracks/:id", getTrackById)
	router.GET("/inventory-api/v1/tracks/:id/comments", getCommentsByTrackId)
	router.GET("/inventory-api/v1/playlists/users/:id", getPlaylistsByUserId)

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}
}

// Handler function to get a track by id
func getTrackById(c *gin.Context) {
	id := c.Param("id")

	for _, track := range tracks {
		if track.TrackId == id {
			c.JSON(http.StatusOK, track)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Track not found"})
}

// Handler function to get a comments by track id
func getCommentsByTrackId(c *gin.Context) {
	id := c.Param("id")

	var commentsRes []models.Comment
	for _, comment := range comments {
		if comment.TrackId == id {
			commentsRes = append(commentsRes, comment)
		}
	}

	c.JSON(http.StatusOK, commentsRes)
}

// Handler function to get a comments by track id
func getPlaylistsByUserId(c *gin.Context) {
	c.JSON(http.StatusOK, playlists)
}

func seed() {
	tracks = append(tracks, models.NewTrack("123", "Song 1", "Artist 1", "Album 1", "4:30", "320kbps", "Pop", "2022-01-01", 50, 1000, true))
	tracks = append(tracks, models.NewTrack("2", "Song 2", "Artist 2", "Album 2", "3:45", "256kbps", "Rock", "2022-02-15", 30, 800, true))
	tracks = append(tracks, models.NewTrack("3", "Song 3", "Artist 3", "Album 3", "5:15", "192kbps", "Hip Hop", "2022-03-20", 20, 500, true))

	playlists = append(playlists, models.NewPlaylist("101", "My Playlist 1", 1))
	playlists = append(playlists, models.NewPlaylist("102", "My Playlist 2", 2))
	playlists = append(playlists, models.NewPlaylist("103", "My Playlist 3", 3))

	comments = append(comments, models.NewComment("user1", "123", "Nice track!", time.Now()))
	comments = append(comments, models.NewComment("user2", "123", "Love this playlist!", time.Now()))
	comments = append(comments, models.NewComment("user3", "123", "Great song!", time.Now()))

}
