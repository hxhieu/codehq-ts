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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timesheetCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timesheetCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
