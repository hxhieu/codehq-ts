package services

import (
	"io"
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type provider struct {
	db *gorm.DB
}

var timesheetService TimesheetService

func init() {
	dsn := ""

	file, err := os.Open("connstr")
	// Can open the connstr file
	if err == nil {
		b, err := io.ReadAll(file)
		// Can read the connstr file
		if err == nil {
			dsn = string(b)
		}
		file.Close()
	}

	// Fallback to env var
	if len(dsn) == 0 {
		dsn = os.Getenv("CODEHQ_TS_CONNECTION_STRING")
	}

	// Default no logging as it would interfere with result parsing
	runLogger := logger.Default.LogMode(logger.Silent)
	// OR debug mode
	if os.Getenv("CODEHQ_TS_RUN_MODE") == "debug" {
		runLogger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: runLogger,
	})
	if err != nil {
		log.Fatal("Fail to init db: ", err)
	}
	timesheetService = &provider{db}
}

func Timesheet() TimesheetService { return timesheetService }
