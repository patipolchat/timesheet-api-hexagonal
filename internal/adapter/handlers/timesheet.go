package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/patipolchat/timesheet-api-hexagonal/internal/core/ports"
	"github.com/patipolchat/timesheet-api-hexagonal/internal/models"
	"github.com/patipolchat/timesheet-api-hexagonal/pkg/customEcho"
)

type Timesheet struct {
	timesheetService ports.TimesheetService
}

func (t *Timesheet) HandleCreateRequest(c echo.Context) error {
	cc := c.(*customEcho.Context)
	ctx := cc.Request().Context()
	req := new(CreateTimesheetRequest)
	if err := cc.BindAndValidate(req); err != nil {
		return err
	}
	createTimesheet := models.CreateTimesheet{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Details:   req.Detail,
	}
	timesheet, err := t.timesheetService.Create(ctx, &createTimesheet)
	if err != nil {
		return err
	}
	return cc.JSON(201, timesheet)
}

func NewTimesheetHandler(timesheetService ports.TimesheetService) *Timesheet {
	return &Timesheet{
		timesheetService: timesheetService,
	}
}
