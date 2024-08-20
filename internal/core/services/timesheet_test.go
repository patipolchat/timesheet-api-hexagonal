package services

import (
	"context"
	"github.com/patipolchat/timesheet-api-hexagonal/internal/core/entities"
	"github.com/patipolchat/timesheet-api-hexagonal/internal/models"
	mockPorts "github.com/patipolchat/timesheet-api-hexagonal/mocks/internal_/core/ports"
	"github.com/patipolchat/timesheet-api-hexagonal/pkg/types"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TimesheetSuite struct {
	suite.Suite
	service           *Timesheet
	mockTimesheetRepo *mockPorts.TimesheetRepository
	StartTime         time.Time
	EndTme            time.Time
	OverTime          time.Time
	LeaveTime         time.Time
}

func (s *TimesheetSuite) SetupSuite() {
	s.StartTime, _ = time.Parse(time.RFC3339, "2024-08-20T08:30:00Z")
	s.EndTme, _ = time.Parse(time.RFC3339, "2024-08-20T17:30:00Z")
	s.OverTime, _ = time.Parse(time.RFC3339, "2024-08-20T18:30:00Z")
	s.LeaveTime, _ = time.Parse(time.RFC3339, "2024-08-20T13:30:00Z")
}

func (s *TimesheetSuite) SetupTest() {
	s.mockTimesheetRepo = mockPorts.NewTimesheetRepository(s.T())
	s.service = &Timesheet{
		timesheetRepo: s.mockTimesheetRepo,
	}
}

func (s *TimesheetSuite) TestCreate() {
	ctx := context.Background()
	expect := &entities.Timesheet{
		ID: "",
		Date: types.PureDate{
			Year:  2024,
			Month: 8,
			Day:   20,
		},
		StartTime:    types.PureTime{Hour: 8, Minute: 30},
		EndTime:      types.PureTime{Hour: 17, Minute: 30},
		WorkHour:     types.PureTime{Hour: 8, Minute: 0},
		OverTimeHour: types.PureTime{Hour: 0, Minute: 0},
		LeaveTime:    types.PureTime{Hour: 0, Minute: 0},
		Details:      "Create Timesheet",
		Reason:       "Create Timesheet",
	}
	s.mockTimesheetRepo.EXPECT().Create(ctx, expect).Return(expect, nil)

	createTimesheet := &models.CreateTimesheet{
		StartTime: &s.StartTime,
		EndTime:   &s.EndTme,
		Details:   "Create Timesheet",
	}
	got, err := s.service.Create(ctx, createTimesheet)
	s.Require().NoError(err)
	s.Equal(expect, got)
}

func (s *TimesheetSuite) TestCreateOverTimeHour() {
	ctx := context.Background()
	expect := &entities.Timesheet{
		ID: "",
		Date: types.PureDate{
			Year:  2024,
			Month: 8,
			Day:   20,
		},
		StartTime:    types.PureTime{Hour: 8, Minute: 30},
		EndTime:      types.PureTime{Hour: 18, Minute: 30},
		WorkHour:     types.PureTime{Hour: 8, Minute: 0},
		OverTimeHour: types.PureTime{Hour: 1, Minute: 0},
		LeaveTime:    types.PureTime{Hour: 0, Minute: 0},
		Details:      "Overtime",
		Reason:       "Create Timesheet",
	}
	s.mockTimesheetRepo.EXPECT().Create(ctx, expect).Return(expect, nil)

	createTimesheet := &models.CreateTimesheet{
		StartTime: &s.StartTime,
		EndTime:   &s.OverTime,
		Details:   "Overtime",
	}
	got, err := s.service.Create(ctx, createTimesheet)
	s.Require().NoError(err)
	s.Equal(expect, got)
}

func (s *TimesheetSuite) TestCreateLeaveHour() {
	ctx := context.Background()
	expect := &entities.Timesheet{
		ID: "",
		Date: types.PureDate{
			Year:  2024,
			Month: 8,
			Day:   20,
		},
		StartTime:    types.PureTime{Hour: 8, Minute: 30},
		EndTime:      types.PureTime{Hour: 13, Minute: 30},
		WorkHour:     types.PureTime{Hour: 4, Minute: 0},
		OverTimeHour: types.PureTime{Hour: 0, Minute: 0},
		LeaveTime:    types.PureTime{Hour: 4, Minute: 0},
		Details:      "Overtime",
		Reason:       "Create Timesheet",
	}
	s.mockTimesheetRepo.EXPECT().Create(ctx, expect).Return(expect, nil)

	createTimesheet := &models.CreateTimesheet{
		StartTime: &s.StartTime,
		EndTime:   &s.LeaveTime,
		Details:   "Overtime",
	}
	got, err := s.service.Create(ctx, createTimesheet)
	s.Require().NoError(err)
	s.Equal(expect, got)
}

func TestTimesheetSuite(t *testing.T) {
	suite.Run(t, new(TimesheetSuite))
}
