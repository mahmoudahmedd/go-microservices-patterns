package models

import "time"

// Comment represents a repository
type Comment struct {
	UserId      string    `json:"user_id"`
	TrackId     string    `json:"track_id"`
	CommentText string    `json:"comment_text"`
	Timestamp   time.Time `json:"timestamp"`
}
