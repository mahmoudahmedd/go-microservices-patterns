package models

// Track represents a repository
type Track struct {
	TrackId      string `json:"track_id"`
	Title        string `json:"title"`
	Artist       string `json:"artist"`
	Album        string `json:"album"`
	Duration     string `json:"duration"`
	Bitrate      string `json:"bitrate"`
	Genre        string `json:"genre"`
	ReleaseDate  string `json:"release_date"`
	Likes        int    `json:"likes"`
	Plays        int    `json:"plays"`
	Downloadable bool   `json:"downloadable"`
}

// NewTrack initializes a new Customer
func NewTrack(trackId, title, artist, album, duration, bitrate, genre, releaseDate string, likes, plays int, downloadable bool) Track {
	return Track{
		TrackId:      trackId,
		Title:        title,
		Artist:       artist,
		Album:        album,
		Duration:     duration,
		Bitrate:      bitrate,
		Genre:        genre,
		ReleaseDate:  releaseDate,
		Likes:        likes,
		Plays:        plays,
		Downloadable: downloadable,
	}
}
