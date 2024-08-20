package services

import (
	"context"
	"github.com/patipolchat/timesheet-api-hexagonal/internal/core/entities"
	"github.com/patipolchat/timesheet-api-hexagonal/internal/core/ports"
	"github.com/patipolchat/timesheet-api-hexagonal/internal/models"
	"github.com/patipolchat/timesheet-api-hexagonal/pkg/types"
	"time"
)

type Timesheet struct {
	timesheetRepo ports.TimesheetRepository
}

func (t *Timesheet) Create(ctx context.Context, createTimesheet *models.CreateTimesheet) (*entities.Timesheet, error) {
	y, m, d := createTimesheet.StartTime.Date()
	date := types.PureDate{Year: y, Month: int(m), Day: d}
	startTime := types.PureTime{
		Hour:   createTimesheet.StartTime.Hour(),
		Minute: createTimesheet.StartTime.Minute(),
	}
	endTime := types.PureTime{
		Hour:   createTimesheet.EndTime.Hour(),
		Minute: createTimesheet.EndTime.Minute(),
	}
	workHour, overTime, leaveTime := t.calHour(ctx, startTime, endTime)

	timesheet := &entities.Timesheet{
		Date:         date,
		StartTime:    startTime,
		EndTime:      endTime,
		WorkHour:     workHour,
		OverTimeHour: overTime,
		Details:      createTimesheet.Details,
		LeaveTime:    leaveTime,
		Reason:       "Create Timesheet",
	}

	return t.timesheetRepo.Create(ctx, timesheet)
}

func (t *Timesheet) Update(ctx context.Context, updateTimesheet *models.UpdateTimesheet) (*entities.Timesheet, error) {
	timesheet, err := t.getOne(ctx, *updateTimesheet.StartTime)
	if err != nil {
		return nil, err
	}

	timesheet.StartTime = types.PureTime{
		Hour:   updateTimesheet.StartTime.Hour(),
		Minute: updateTimesheet.StartTime.Minute(),
	}

	timesheet.EndTime = types.PureTime{
		Hour:   updateTimesheet.EndTime.Hour(),
		Minute: updateTimesheet.EndTime.Minute(),
	}

	timesheet.WorkHour, timesheet.OverTimeHour, timesheet.LeaveTime = t.calHour(ctx, timesheet.StartTime, timesheet.EndTime)

	timesheet.Details = updateTimesheet.Details
	timesheet.Reason = updateTimesheet.Reason

	return t.timesheetRepo.Update(ctx, timesheet)
}

func (t *Timesheet) Delete(ctx context.Context, getTimesheet models.GetTimesheet) error {
	y, m, d := getTimesheet.Date.Date()
	pureDate := types.PureDate{Year: y, Month: int(m), Day: d}
	return t.timesheetRepo.DeleteByDate(ctx, pureDate)
}

func (t *Timesheet) Get(ctx context.Context, date time.Time) (*entities.Timesheet, error) {
	y, m, d := date.Date()
	pureDate := types.PureDate{Year: y, Month: int(m), Day: d}
	return t.timesheetRepo.FindByDate(ctx, pureDate)
}

func (t *Timesheet) GetAll(ctx context.Context, paginate models.Paginate) ([]*entities.Timesheet, error) {
	limit := paginate.PerPage
	offset := paginate.PerPage * (paginate.Page - 1)
	return t.timesheetRepo.GetAll(ctx, limit, offset)
}

func (t *Timesheet) getOne(ctx context.Context, date time.Time) (*entities.Timesheet, error) {
	y, m, d := date.Date()
	pureDate := types.PureDate{Year: y, Month: int(m), Day: d}
	return t.timesheetRepo.FindByDate(ctx, pureDate)
}

func (t *Timesheet) calHour(ctx context.Context, startTime types.PureTime, endTime types.PureTime) (workHour, overTime, leaveTime types.PureTime) {
	launchMin := 0
	if startTime.Hour <= 12 && endTime.Hour >= 13 {
		launchMin = 60
	}
	startMinutes := startTime.ToMinutes()
	endMinutes := endTime.ToMinutes()
	workMinutes := endMinutes - startMinutes - launchMin
	if workMinutes < 0 {
		workMinutes = 0
	}
	workHour = types.PureTime{}
	workHour.FromMinutes(workMinutes)
	if workHour.ToMinutes() > 8*60 {
		overMin := workHour.ToMinutes() - 8*60
		overTime.FromMinutes(overMin)
		workHour.FromMinutes(8 * 60)
	}

	if workHour.ToMinutes() < 8*60 {
		leave := 8*60 - workHour.ToMinutes()
		leaveTime.FromMinutes(leave)
	}
	return workHour, overTime, leaveTime
}

func NewTimesheetService(timesheetRepo ports.TimesheetRepository) *Timesheet {
	return &Timesheet{
		timesheetRepo: timesheetRepo,
	}
}
