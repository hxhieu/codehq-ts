/*
Copyright © 2021 Hugh Hoang <hugh.hoang@codehq.nz>
For internal use only
*/
package cmd

import (
	"time"

	"github.com/hxhieu/codehq-ts/configuration"
	"github.com/spf13/cobra"
)

var GetTimesheetDate string

// timesheetGetCmd represents the timesheetGet command
var timesheetGetCmd = &cobra.Command{
	Use:   "timesheet",
	Short: "Get employee timesheet record(s)",
	Long:  `Get details about the employee timesheet record(s)`,
}

func init() {
	getCmd.AddCommand(timesheetGetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	timesheetGetCmd.PersistentFlags().StringVarP(
		&GetTimesheetDate,
		"date",
		"d",
		time.Now().Format(configuration.Get().DateFormat),
		"Records start date (DD/MM/YYYY)",
	)
}
