package main

import (
	"time"

	"github.com/hxhieu/codehq-ts/models"
	"github.com/hxhieu/codehq-ts/services"
)

func main() {
	// testDate := time.Date(2021, 9, 1, 0, 0, 0, 0, time.Local)
	// test, err := services.Timesheet().GetWeeklyTimesheet("hugh.hoang", &testDate)
	// OutputResult(models.Response{Data: test})
	entry := models.Timesheet{
		EmployeeId:  "hugh.hoang",
		ProjectId:   100100,
		Phase:       "99",
		Code:        "ADMIN",
		Description: "Hello 123",
		Start:       time.Now(),
		End:         time.Now().Add(15 * time.Minute),
	}
	newId, err1 := services.Timesheet().AddTimesheet(&entry)
	services.Timesheet().Dispose()
	if err1 != nil {
		OutputError(&err1)
	} else {
		OutputResult(newId)
	}
}
