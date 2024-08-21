package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type PureDate struct {
	Year  int `json:"year" validate:"required,numeric,min=2000,max=2100"`
	Month int `json:"month" validate:"required,numeric,min=1,max=12"`
	Day   int `json:"day" validate:"required,numeric,min=1,max=31"`
}

func (d PureDate) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

func (d PureDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *PureDate) UnmarshalJSON(data []byte) error {
	date := string(data)
	_, err := parseDate(date, d)
	return err
}

func (d *PureDate) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal date value:", value))
	}
	if _, err := fmt.Sscanf(str, `%04d-%02d-%02d`, &d.Year, &d.Month, &d.Day); err != nil {
		return err
	}
	return nil
}

func (d PureDate) Value() (driver.Value, error) {
	return d.String(), nil
}

func parseDate(date string, d *PureDate) (PureDate, error) {
	var year, month, day int
	if _, err := fmt.Sscanf(date, `"%04d-%02d-%02d"`, &year, &month, &day); err != nil {
		return PureDate{}, err
	}
	*d = PureDate{Year: year, Month: month, Day: day}
	validate := validator.New()
	if err := validate.Struct(d); err != nil {
		return PureDate{}, err
	}
	return *d, nil
}
