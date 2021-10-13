/*
Copyright © 2021 Hugh Hoang <hugh.hoang@codehq.nz>
For internal use only
*/
package cmd

import (
	"github.com/hxhieu/codehq-ts/output"
	"github.com/hxhieu/codehq-ts/services"
	"github.com/spf13/cobra"
)

// timesheetGetWeeklyCmd represents the timesheetGetWeekly command
var timesheetGetWeeklyCmd = &cobra.Command{
	Use:   "weekly",
	Short: "Get the employee weekly record(s)",
	Long:  `Get the employee weekly record(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ParseEmployeeId()
		if err != nil {
			output.ConsoleErrorJson(&err)
		}

		date, err := ParseInputDate()
		if err != nil {
			output.ConsoleErrorJson(&err)
		}

		// Service call
		result, err := services.Timesheet().GetWeeklyTimesheet(EmployeeId, date)
		if err != nil {
			output.ConsoleErrorJson(&err)
		}

		// Output
		switch OutputFormat {
		case "table":
			tableData := [][]string{}
			for _, record := range *result {
				tableData = append(tableData, record.ToStringArray())
			}
			output.ConsoleResultTable(tableData, []string{
				"Id",
				"EmployeeId",
				"ProjectId",
				"Phase",
				"Code",
				"Charge",
				"Description",
				"Start",
				"End",
				"Duration",
			})
		default:
			output.ConsoleResultJson(result)
		}
	},
}

func init() {
	timesheetGetCmd.AddCommand(timesheetGetWeeklyCmd)
}
