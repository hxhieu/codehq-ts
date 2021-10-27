package cmd

import (
	"errors"
	"fmt"
	"time"
)

func ParseEmployeeId() error {
	if len(EmployeeId) == 0 {
		return errors.New("employee ID is required")
	}
	return nil
}

func ParseInputDate() (*time.Time, error) {
	if date, err := time.ParseInLocation(DateLayout, InputDate, time.Local); err != nil {
		return nil, fmt.Errorf("failed to parse date. '%s' is not a valid 'DDMMYYYY'", InputDate)
	} else {
		return &date, nil
	}
}

func ParseTime(timeString string, name string) (*time.Time, error) {
	if t, err := time.ParseInLocation(TimeLayout, timeString, time.Local); err != nil {
		return nil, fmt.Errorf("failed to parse %s. '%s' is not a valid 'HHmm'", name, timeString)
	} else {
		if t.Minute()%15 != 0 {
			return nil, fmt.Errorf("invalid %s minutes, the minutes should be divisible by 15mins i.e. 15, 30, 45 and 60", name)
		}
		if date, err := ParseInputDate(); err != nil {
			return nil, err
		} else {
			duration := time.Hour*time.Duration(t.Hour()) + time.Minute*time.Duration(t.Minute())
			// Clone the date
			result := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
			result = result.Add(duration)
			return &result, nil
		}
	}
}
