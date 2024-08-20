package entities

import (
	"github.com/patipolchat/timesheet-api-hexagonal/pkg/types"
)

type Timesheet struct {
	ID           string
	Date         types.PureDate
	StartTime    types.PureTime
	EndTime      types.PureTime
	WorkHour     types.PureTime
	OverTimeHour types.PureTime
	LeaveTime    types.PureTime
	Details      string
	Reason       string
}
