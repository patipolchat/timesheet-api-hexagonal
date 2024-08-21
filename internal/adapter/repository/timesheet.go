package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/patipolchat/timesheet-api-hexagonal/internal/core/entities"
	"github.com/patipolchat/timesheet-api-hexagonal/pkg/types"
	"gorm.io/gorm"
)

type TimesheetRepository struct {
	gormDB *gorm.DB
}

func (r *TimesheetRepository) Create(ctx context.Context, timesheet *entities.Timesheet) (*entities.Timesheet, error) {
	tx := r.gormDB.WithContext(ctx)
	id := uuid.New()
	err := tx.Exec(
		`INSERT INTO timesheets (id, date, start_time, end_time, work_hour, over_time_hour, leave_time, details, reason)
    		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		id,
		timesheet.Date.String(),
		timesheet.StartTime.String(),
		timesheet.EndTime.String(),
		timesheet.WorkHour.String(),
		timesheet.OverTimeHour.String(),
		timesheet.LeaveTime.String(),
		timesheet.Details,
		timesheet.Reason,
	).Error
	if err != nil {
		return nil, err
	}
	timesheet.ID = id
	return timesheet, nil
}

func (r *TimesheetRepository) Update(ctx context.Context, timesheet *entities.Timesheet) (*entities.Timesheet, error) {
	tx := r.gormDB.WithContext(ctx)
	err := tx.Exec(
		`UPDATE timesheets 
					SET start_time = ?, end_time = ?, work_hour = ?, over_time_hour = ?, leave_time = ?, details = ?, reason = ?
					WHERE id = ?`,
		timesheet.StartTime,
		timesheet.EndTime,
		timesheet.WorkHour,
		timesheet.OverTimeHour,
		timesheet.LeaveTime,
		timesheet.Details,
		timesheet.Reason,
		timesheet.ID,
	).Error
	if err != nil {
		return nil, err
	}
	return timesheet, nil
}

func (r *TimesheetRepository) FindByDate(ctx context.Context, date types.PureDate) (*entities.Timesheet, error) {
	tx := r.gormDB.WithContext(ctx)
	var result entities.Timesheet
	err := tx.Raw(`SELECT * FROM timesheets WHERE date = ? LIMIT 1`, date.String()).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *TimesheetRepository) DeleteByDate(ctx context.Context, date types.PureDate) error {
	tx := r.gormDB.WithContext(ctx)
	err := tx.Exec(`DELETE FROM timesheets WHERE date = ?`, date).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TimesheetRepository) GetAll(ctx context.Context, limit int, offset int) ([]*entities.Timesheet, error) {
	tx := r.gormDB.WithContext(ctx)
	var result []*entities.Timesheet
	err := tx.Raw(`SELECT * FROM timesheets LIMIT ? OFFSET ?`, limit, offset).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewTimesheetRepository(gormDB *gorm.DB) *TimesheetRepository {
	return &TimesheetRepository{gormDB: gormDB}
}
