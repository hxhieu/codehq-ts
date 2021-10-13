/*
Copyright © 2021 Hugh Hoang <hugh.hoang@codehq.nz>
For internal use only
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var projectId int64
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
		// if start, err := time.
	},
}

func init() {
	createCmd.AddCommand(timesheetCreateCmd)

	timesheetCreateCmd.Flags().Int64VarP(&projectId, "project", "p", 0, "The Project ID")
	timesheetCreateCmd.Flags().StringVar(&phase, "project-phase", "", "The Project phase")
	timesheetCreateCmd.Flags().StringVar(&code, "code", "ANLYS", "The Timesheet activity")
	timesheetCreateCmd.Flags().BoolVar(&nonCharge, "nc", false, "If the record is N/C?")
	timesheetCreateCmd.Flags().StringVar(&desc, "desc", "", "The Timesheet description")
	timesheetCreateCmd.Flags().StringVar(&startString, "start", "", "The Timesheet start time (HH:mm)")
	timesheetCreateCmd.Flags().StringVar(&endString, "end", "", "The Timesheet end time (HH:mm)")
}
