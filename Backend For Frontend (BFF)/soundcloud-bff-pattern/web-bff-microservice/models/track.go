package models

// Track represents a repository
type Track struct {
	TrackId     string `json:"track_id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	Duration    string `json:"duration"`
	Bitrate     string `json:"bitrate"`
	Genre       string `json:"genre"`
	ReleaseDate string `json:"release_date"`
	Likes       int    `json:"likes"`
	Plays       int    `json:"plays"`
}
