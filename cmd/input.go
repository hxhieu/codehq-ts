package cmd

import (
	"fmt"
	"log"
	"time"
)

func ParseDate() *time.Time {
	if date, err := time.ParseInLocation(DateLayout, InputDate, time.Local); err != nil {
		log.Fatal(fmt.Sprintf("Failed to parse date. '%s' is not a valid 'DD/MM/YYYY'", InputDate))
	} else {
		return &date
	}
	return nil
}

func EnsureEmployeeId() {
	if len(EmployeeId) == 0 {
		log.Fatal("Employee ID is required.")
	}
}
