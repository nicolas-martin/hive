package model

import "github.com/google/uuid"

type Update struct {
	Usernames []string
}

type User struct {
	DisplayName string
	SlackUserID string
}

type UserUpdate struct {
	UserID       uuid.UUID
	UpdateID     uuid.UUID
	RecordingURL string
}
