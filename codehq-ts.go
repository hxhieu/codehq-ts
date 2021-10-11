package main

import (
	"time"

	"github.com/hxhieu/codehq-ts/services"
)

func main() {
	testDate := time.Date(2021, 9, 1, 0, 0, 0, 0, time.Local)
	test, err := services.Timesheet().GetWeeklyTimesheet("hugh.hoang", &testDate)
	if err != nil {
		OutputError(&err)
	} else {
		OutputResult(test)
	}
}
