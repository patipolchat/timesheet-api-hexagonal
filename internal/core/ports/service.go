package ports

import (
	"context"
	"github.com/patipolchat/timesheet-api-hexagonal/internal/core/entities"
	"github.com/patipolchat/timesheet-api-hexagonal/internal/models"
	"time"
)

type TimesheetService interface {
	Create(ctx context.Context, createTimesheet *models.CreateTimesheet) (*entities.Timesheet, error)
	Update(ctx context.Context, updateTimesheet *models.UpdateTimesheet) (*entities.Timesheet, error)
	Delete(ctx context.Context, getTimesheet models.GetTimesheet) error
	Get(ctx context.Context, date time.Time) (*entities.Timesheet, error)
	GetAll(ctx context.Context, paginate models.Paginate) ([]*entities.Timesheet, error)
}
