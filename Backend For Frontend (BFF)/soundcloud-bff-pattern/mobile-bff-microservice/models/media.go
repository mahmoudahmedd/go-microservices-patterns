package models

// Media struct to represent a track
type Media struct {
	Files map[string]string `json:"download_links"`
}
