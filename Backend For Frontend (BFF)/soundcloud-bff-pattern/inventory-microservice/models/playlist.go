package models

// Playlist represents a repository
type Playlist struct {
	PlaylistId   string `json:"playlist_id"`
	PlaylistName string `json:"playlist_name"`
	Position     int    `json:"position"`
}

// NewPlaylist creates and returns a new Playlist with specified values
func NewPlaylist(playlistId, playlistName string, position int) Playlist {
	return Playlist{
		PlaylistId:   playlistId,
		PlaylistName: playlistName,
		Position:     position,
	}
}
