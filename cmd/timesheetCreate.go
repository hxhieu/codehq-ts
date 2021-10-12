/*
Copyright © 2021 Hugh Hoang <hugh.hoang@codehq.nz>
For internal use only
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// timesheetCreateCmd represents the timesheetCreate command
var timesheetCreateCmd = &cobra.Command{
	Use:   "timesheet",
	Short: "Create employee timesheet record(s)",
	Long:  `Add the Timesheet record(s) for the employee`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("timesheetCreate called")
	},
}

func init() {
	createCmd.AddCommand(timesheetCreateCmd)
}
