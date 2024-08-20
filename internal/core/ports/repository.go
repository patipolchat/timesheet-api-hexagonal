package ports

import (
	"context"
	"github.com/patipolchat/timesheet-api-hexagonal/internal/core/entities"
	"github.com/patipolchat/timesheet-api-hexagonal/pkg/types"
)

type TimesheetRepository interface {
	Create(ctx context.Context, timesheet *entities.Timesheet) (*entities.Timesheet, error)
	Update(ctx context.Context, timesheet *entities.Timesheet) (*entities.Timesheet, error)
	FindByDate(ctx context.Context, date types.PureDate) (*entities.Timesheet, error)
	DeleteByDate(ctx context.Context, date types.PureDate) error
	GetAll(ctx context.Context, limit int, offset int) ([]*entities.Timesheet, error)
}
