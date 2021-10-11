package services

import (
	"time"

	"github.com/hxhieu/codehq-ts/models"
)

type TimesheetService interface {
	AddTimesheet(entry *models.Timesheet) error
	GetWeeklyTimesheet(employeeId string, start *time.Time, end *time.Time) (*[]models.Timesheet, error)
}

func (orm *provider) AddTimesheet(entry *models.Timesheet) (err error) {
	return
}

func (orm *provider) GetWeeklyTimesheet(employeeId string, start *time.Time, end *time.Time) (entries *[]models.Timesheet, err error) {
	records := []models.Timesheet{}
	result := orm.db.Limit(2).Where(&models.Timesheet{EmployeeId: "hugh.hoang"}).Find(&records)

	if result.Error != nil {
		return nil, result.Error
	}

	return &records, nil
}
