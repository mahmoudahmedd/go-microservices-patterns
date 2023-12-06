package models

import "time"

// Comment represents a repository
type Comment struct {
	UserId      string    `json:"user_id"`
	TrackId     string    `json:"track_id"`
	CommentText string    `json:"comment_text"`
	Timestamp   time.Time `json:"timestamp"`
}

// NewComment creates and returns a new Comment with specified values.
func NewComment(userId, trackId, commentText string, timestamp time.Time) Comment {
	return Comment{
		UserId:      userId,
		TrackId:     trackId,
		CommentText: commentText,
		Timestamp:   timestamp,
	}
}
