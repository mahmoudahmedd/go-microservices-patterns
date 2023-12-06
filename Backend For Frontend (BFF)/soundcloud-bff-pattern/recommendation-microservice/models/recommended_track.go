package models

// RecommendedTrack represents a repository
type RecommendedTrack struct {
	TrackId string `json:"track_id"`
	Title   string `json:"title"`
	Artist  string `json:"artist"`
}

// NewRecommendedTrack initializes a new Customer
func NewRecommendedTrack(trackId, title, artist string) RecommendedTrack {
	return RecommendedTrack{
		TrackId: trackId,
		Title:   title,
		Artist:  artist,
	}
}
