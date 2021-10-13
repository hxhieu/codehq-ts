/*
Copyright © 2021 Hugh Hoang <hugh.hoang@codehq.nz>
For internal use only
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/hxhieu/codehq-ts/models"
	"github.com/hxhieu/codehq-ts/output"
	"github.com/hxhieu/codehq-ts/services"
	"github.com/spf13/cobra"
)

var projectId int32
var phase string
var code string
var nonCharge bool
var desc string
var startString string
var endString string

// timesheetCreateCmd represents the timesheetCreate command
var timesheetCreateCmd = &cobra.Command{
	Use:   "timesheet",
	Short: "Create employee timesheet record(s)",
	Long:  `Add the Timesheet record(s) for the employee`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ParseEmployeeId()
		if err != nil {
			output.ConsoleErrorJson(&err)
		}

		if projectId <= 0 {
			err = errors.New("project ID is required")
			output.ConsoleErrorJson(&err)
		}

		if len(phase) == 0 {
			err = errors.New("project phase is required")
			output.ConsoleErrorJson(&err)
		}

		start, err := ParseTime(startString, "start time")
		if err != nil {
			output.ConsoleErrorJson(&err)
		}

		end, err := ParseTime(endString, "end time")
		if err != nil {
			output.ConsoleErrorJson(&err)
		}

		// New record
		record := models.Timesheet{
			EmployeeId:  EmployeeId,
			ProjectId:   projectId,
			Phase:       phase,
			Code:        code,
			Description: desc,
			Start:       *start,
			End:         *end,
		}
		if nonCharge {
			record.Charge = "N/C"
		}

		// Insert
		if newId, err := services.Timesheet().AddTimesheet(&record); err != nil {
			output.ConsoleErrorJson(&err)
		} else {
			switch OutputFormat {

			case "table":
				output.ConsoleResultTable([][]string{
					{fmt.Sprintf("%d", newId)},
				}, []string{
					"NewId",
				})

			default:
				output.ConsoleResultJson(newId)
			}
		}
	},
}

func init() {
	createCmd.AddCommand(timesheetCreateCmd)

	timesheetCreateCmd.Flags().Int32VarP(&projectId, "project", "p", 0, "The Project ID")
	timesheetCreateCmd.Flags().StringVar(&phase, "phase", "", "The Project phase")
	timesheetCreateCmd.Flags().StringVar(&code, "code", "ANLYS", "The Timesheet activity")
	timesheetCreateCmd.Flags().BoolVar(&nonCharge, "nc", false, "If the record is N/C?")
	timesheetCreateCmd.Flags().StringVar(&desc, "desc", "", "The Timesheet description")
	timesheetCreateCmd.Flags().StringVar(&startString, "start", "", "The Timesheet start time (HHmm)")
	timesheetCreateCmd.Flags().StringVar(&endString, "end", "", "The Timesheet end time (HHmm)")
}
