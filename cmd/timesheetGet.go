/*
Copyright © 2021 Hugh Hoang <hugh.hoang@codehq.nz>
For internal use only
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var weeklyFlag string = "weekly"
var dateFlag string = "date"

// timesheetGetCmd represents the timesheetGet command
var timesheetGetCmd = &cobra.Command{
	Use:   "timesheet",
	Short: "Get employee timesheet record(s)",
	Long:  `Get details about the employee timesheet record(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Flags().GetString(dateFlag))
	},
}

func init() {
	getCmd.AddCommand(timesheetGetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	timesheetGetCmd.Flags().BoolP(
		weeklyFlag,
		weeklyFlag[0:1],
		true,
		"Get the weekly records",
	)

	timesheetGetCmd.Flags().StringP(
		dateFlag,
		dateFlag[0:1],
		time.Now().Format("02/01/2006"),
		"Records start date (DD/MM/YYYY)",
	)
}
