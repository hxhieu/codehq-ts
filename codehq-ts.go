package main

import (
	"github.com/hxhieu/codehq-ts/services"
)

func main() {
	test, err := services.Timesheet().GetWeeklyTimesheet("aaa", nil, nil)
	if err != nil {
		OutputError(&err)
	} else {
		OutputResult(test)
	}
}
