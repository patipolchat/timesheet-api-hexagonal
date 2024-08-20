package entities

import (
	"github.com/google/uuid"
	"github.com/patipolchat/timesheet-api-hexagonal/pkg/types"
)

type Holiday struct {
	ID   uuid.UUID
	Name string
	Date *types.PureDate
}
