package handlers

import "time"

type CreateTimesheetRequest struct {
	StartTime *time.Time `json:"start_time" validate:"required"`
	EndTime   *time.Time `json:"end_time" validate:"required"`
	Detail    string     `json:"detail" validate:"required"`
}
