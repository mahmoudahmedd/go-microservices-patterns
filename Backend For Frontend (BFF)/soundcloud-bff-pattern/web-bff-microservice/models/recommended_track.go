package models

// RecommendedTrack represents a repository
type RecommendedTrack struct {
	TrackId string `json:"track_id"`
	Title   string `json:"title"`
	Artist  string `json:"artist"`
}
