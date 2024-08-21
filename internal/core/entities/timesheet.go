package entities

import (
	"github.com/google/uuid"
	"github.com/patipolchat/timesheet-api-hexagonal/pkg/types"
)

type Timesheet struct {
	ID           uuid.UUID
	Date         types.PureDate
	StartTime    types.PureTime
	EndTime      types.PureTime
	WorkHour     types.PureTime
	OverTimeHour types.PureTime
	LeaveTime    types.PureTime
	Details      string
	Reason       string
}

func (t Timesheet) TableName() string {
	return "timesheets"
}
