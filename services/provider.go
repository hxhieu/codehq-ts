package services

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/hxhieu/codehq-ts/configuration"
	"github.com/hxhieu/codehq-ts/output"
)

type provider struct {
	db *gorm.DB
}

// Services
var timesheetService TimesheetService

func init() {
	config := configuration.Get()

	// Default no logging as it would interfere with result parsing
	runLogger := logger.Default.LogMode(logger.Silent)
	// OR debug mode
	if config.Debug {
		runLogger = logger.Default.LogMode(logger.Info)
	}

	// Timesheet DB
	db, err := gorm.Open(sqlserver.Open(config.TimesheetDsn), &gorm.Config{
		Logger: runLogger,
	})
	if err != nil {
		output.ConsoleErrorJson(&err)
	}
	timesheetService = &provider{db}
}

func Timesheet() TimesheetService { return timesheetService }
