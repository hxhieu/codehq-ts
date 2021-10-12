package models

import (
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
