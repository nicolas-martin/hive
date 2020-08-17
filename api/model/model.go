package model

import "github.com/google/uuid"

type Update struct {
	UpdateID uuid.UUID
	Users    []*User
}

type User struct {
	UserID      uuid.UUID
	DisplayName string
	SlackUserID string
}

type UserUpdate struct {
	UserUpdateID uuid.UUID
	UserID       uuid.UUID
	UpdateID     uuid.UUID
	RecordingURL string
}
