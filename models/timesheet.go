package models

import (
	"fmt"
	"time"
)

// Timesheet represents the Timesheet2 table
type Timesheet struct {
	Id          int32     `gorm:"primaryKey;column:TimesheetID"`
	EmployeeId  string    `gorm:"column:EmployeeId"`
	ProjectId   int32     `gorm:"column:ProjectId"`
	Phase       string    `gorm:"column:Phase"`
	Code        string    `gorm:"column:Code"`
	Charge      string    `gorm:"column:Charge"`
	Description string    `gorm:"column:Description"`
	Start       time.Time `gorm:"column:Start"`
	End         time.Time `gorm:"column:End"`
	Duration    float32   `gorm:"column:Duration"`
}

func (Timesheet) TableName() string {
	return "Timesheet2"
}

func (record *Timesheet) ToStringArray() []string {
	return []string{
		fmt.Sprintf("%d", record.Id),
		record.EmployeeId,
		fmt.Sprintf("%d", record.ProjectId),
		record.Phase,
		record.Code,
		record.Charge,
		record.Description,
		record.Start.Format(time.RFC822Z),
		record.End.Format(time.RFC822Z),
		fmt.Sprintf("%.2f", record.Duration),
	}
}
