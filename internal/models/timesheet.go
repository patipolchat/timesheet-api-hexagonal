package models

import (
	"time"
)

type CreateTimesheet struct {
	StartTime *time.Time `json:"start_time" binding:"required" validate:"required"`
	EndTime   *time.Time `json:"end_time" binding:"required" validate:"required"`
	Details   string     `json:"details" binding:"required" validate:"required"`
}

type UpdateTimesheet struct {
	StartTime *time.Time `json:"start_time" binding:"required" validate:"required"`
	EndTime   *time.Time `json:"end_time" binding:"required" validate:"required"`
	Details   string     `json:"details" binding:"required" validate:"required"`
	Reason    string     `json:"reason" binding:"required" validate:"required"`
}

type GetTimesheet struct {
	Date time.Time `json:"date" binding:"required" validate:"required"`
}
