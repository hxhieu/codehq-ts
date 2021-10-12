package services

import (
	"log"
	"time"

	"github.com/hxhieu/codehq-ts/models"
)

type TimesheetService interface {
	AddTimesheet(entry *models.Timesheet) (int64, error)
	GetWeeklyTimesheet(employeeId string, start *time.Time) (*[]models.Timesheet, error)
	Dispose()
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

func (orm *provider) AddTimesheet(entry *models.Timesheet) (int64, error) {
	// HACK: Pretty bad one, just for the sake of Timesheet2's legacy...
	// Cannot use the Create method because the table triggers prevent returning of the new record ID
	// which is the default behaviour of GORM Create method

	// TODO: Must validate the input against the employee project etc

	var newId int64
	result := orm.db.Raw(`
		INSERT INTO Timesheet2 (
			[EmployeeId],
			[ProjectId],
			[Phase],
			[Code],
			[Charge],
			[Description],
			[Start],
			[End]
		) VALUES (
			@EmployeeId,
			@ProjectId,
			@Phase,
			@Code,
			@Charge,
			@Description,
			@Start,
			@End
		);
		SELECT SCOPE_IDENTITY()`, entry).Scan(&newId)
	if result.Error != nil {
		return 0, result.Error
	}
	return newId, nil
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

func (orm *provider) Dispose() {
	db, err := orm.db.DB()
	if err != nil {
		log.Fatal(err)
	} else {
		db.Close()
	}
}
