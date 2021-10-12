/*
Copyright © 2021 Hugh Hoang <hugh.hoang@codehq.nz>
For internal use only
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource",
	Long:  `Create a Timesheet related resource`,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
