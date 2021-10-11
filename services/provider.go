package services

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type provider struct {
	db *gorm.DB
}

var timesheetService TimesheetService

func init() {
	dsn := "sqlserver://dev:!!Dev_123!!@mssql.k8s.local:32001?database=Timesheet&parseTime=true"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to init db: ", err)
	}
	timesheetService = &provider{db}
}

func Timesheet() TimesheetService { return timesheetService }
