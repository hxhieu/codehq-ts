package services

import (
	"time"

	"github.com/hxhieu/codehq-ts/models"
)

type TimesheetService interface {
	AddTimesheet(entry *models.Timesheet) error
	GetWeeklyTimesheet(employeeId string, start *time.Time) (*[]models.Timesheet, error)
}

func getWeekRange(start *time.Time) (*time.Time, *time.Time) {
	input := time.Now()
	if start != nil {
		input = *start
	}
	// Diff from Monday
	diffStart := int(input.Weekday()) - 1
	// Week range
	weekStart := input.AddDate(0, 0, -diffStart) // Monday
	weekEnd := weekStart.AddDate(0, 0, 7)        // 0:00 next Monday so pretty much end of Sunday
	return &weekStart, &weekEnd
}

func (orm *provider) AddTimesheet(entry *models.Timesheet) (err error) {
	return
}

func (orm *provider) GetWeeklyTimesheet(employeeId string, start *time.Time) (entries *[]models.Timesheet, err error) {
	weekStart, weekEnd := getWeekRange(start)
	records := []models.Timesheet{}
	result := orm.db.Where("EmployeeId = ? AND Start BETWEEN ? AND ?", employeeId, weekStart, weekEnd).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}
	return &records, nil
}
