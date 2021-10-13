/*
Copyright © 2021 Hugh Hoang <hugh.hoang@codehq.nz>
For internal use only
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resource(s)",
	Long:  `Get Timesheet related resource(s)`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
