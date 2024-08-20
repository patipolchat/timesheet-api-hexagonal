package types

import (
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

func ParseDate(date string) (PureDate, error) {
	var d PureDate
	_, err := parseDate(date, &d)
	return d, err
}
