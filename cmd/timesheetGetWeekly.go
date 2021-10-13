/*
Copyright © 2021 Hugh Hoang <hugh.hoang@codehq.nz>
For internal use only
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/hxhieu/codehq-ts/services"
	"github.com/spf13/cobra"
)

// timesheetGetWeeklyCmd represents the timesheetGetWeekly command
var timesheetGetWeeklyCmd = &cobra.Command{
	Use:   "weekly",
	Short: "Get the employee weekly record(s)",
	Long:  `Get the employee weekly record(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(EmployeeId) == 0 {
			log.Fatal("Employee ID is required.")
		} else {
			if date, err := time.ParseInLocation(DateLayout, GetTimesheetDate, time.Local); err != nil {
				log.Fatal(fmt.Sprintf("Failed to parse date. '%s' is not a valid 'DD/MM/YYYY'", GetTimesheetDate))
			} else {
				if result, err := services.Timesheet().GetWeeklyTimesheet(EmployeeId, &date); err != nil {
					OutputErrorJson(&err)
				} else {
					switch OutputFormat {

					case "table":
						tableData := [][]string{}
						for _, record := range *result {
							tableData = append(tableData, record.ToStringArray())
						}
						OutputResultTable(tableData, []string{
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
						OutputResultJson(result)
					}
				}
			}
		}
	},
}

func init() {
	timesheetGetCmd.AddCommand(timesheetGetWeeklyCmd)
}
