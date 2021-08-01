package model

import "github.com/google/uuid"

type Schedule struct {
	ScheduleID uuid.UUID
	CronString string
}
