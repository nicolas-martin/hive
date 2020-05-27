package model

import "github.com/google/uuid"

type Update struct {
	ID        uuid.UUID
	Usernames []string
}

type UserUpdate struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	SlackUserID  string
	UpdateID     uuid.UUID
	RecordingURL string
}
