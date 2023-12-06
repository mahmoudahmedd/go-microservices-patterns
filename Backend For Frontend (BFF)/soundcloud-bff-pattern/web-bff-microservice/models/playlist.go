package models

// Playlist represents a repository
type Playlist struct {
	PlaylistId   string `json:"playlist_id"`
	PlaylistName string `json:"playlist_name"`
	Position     int    `json:"position"`
}
