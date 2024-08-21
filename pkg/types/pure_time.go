package types

import (
	"database/sql/driver"
	"errors"
	"github.com/go-playground/validator/v10"

	"fmt"
)

type PureTime struct {
	Hour   int `json:"hour" validate:"required,numeric,min=0,max=23"`
	Minute int `json:"minute" validate:"required,numeric,min=0,max=59"`
}

func (t *PureTime) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal time value:", value))
	}
	if _, err := fmt.Sscanf(str, `%02d:%02d`, &t.Hour, &t.Minute); err != nil {
		return err
	}
	return nil
}

func (t PureTime) Value() (driver.Value, error) {
	return t.String(), nil
}

func (t *PureTime) CalculateWorkHour(endTime *PureTime) *PureTime {
	launchMin := 0
	if t.Hour <= 12 && endTime.Hour >= 13 {
		launchMin = 60
	}
	startMinutes := t.ToMinutes()
	endMinutes := endTime.ToMinutes()
	workMinutes := endMinutes - startMinutes - launchMin
	if workMinutes < 0 {
		workMinutes = 0
	}
	workTime := &PureTime{}
	workTime.FromMinutes(workMinutes)
	return workTime
}

func (t *PureTime) ToMinutes() int {
	return t.Hour*60 + t.Minute
}

func (t *PureTime) FromMinutes(minutes int) {
	t.Hour = minutes / 60
	t.Minute = minutes % 60
}

func (t PureTime) String() string {
	return fmt.Sprintf("%02d:%02d", t.Hour, t.Minute)
}

func (t PureTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

func (t *PureTime) UnmarshalJSON(data []byte) error {
	time := string(data)
	_, err := parseTime(time, t)
	if err != nil {
		return err
	}
	validate := validator.New()
	return validate.Struct(t)
}

func parseTime(time string, t *PureTime) (PureTime, error) {
	var hour, minute int
	if _, err := fmt.Sscanf(time, `"%02d:%02d"`, &hour, &minute); err != nil {
		return PureTime{}, err
	}
	*t = PureTime{Hour: hour, Minute: minute}
	return *t, nil
}
