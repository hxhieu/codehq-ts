/*
Copyright © 2021 Hugh Hoang <hugh.hoang@codehq.nz>
For internal use only
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// timesheetGetCmd represents the timesheetGet command
var timesheetGetCmd = &cobra.Command{
	Use:   "timesheet",
	Short: "Get employee timesheet record(s)",
	Long:  `Get details about the employee timesheet record(s)`,
}

func init() {
	getCmd.AddCommand(timesheetGetCmd)
}
