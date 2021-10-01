package main

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"github.com/hxhieu/codehq-ts/models"
)

func main() {
	dsn := "sqlserver://dev:!!Dev_123!!@mssql.k8s.local:32001?database=Timesheet&parseTime=true"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		test := []models.Timesheet{}
		result := db.Limit(2).Where(&models.Timesheet{EmployeeId: "hugh.hoang"}).Find(&test)
		if result.Error != nil {
			OutputError(&result.Error)
		} else {
			OutputResult(test)
		}
	}
}
