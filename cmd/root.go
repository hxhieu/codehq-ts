/*
Copyright © 2021 Hugh Hoang <hugh.hoang@codehq.nz>
For internal use only
*/
package cmd

import (
	"github.com/hxhieu/codehq-ts/configuration"
	"github.com/spf13/cobra"
)

var EmployeeId string
var OutputFormat string
var DateLayout string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "codehq-ts",
	Short: "CodeHQ Timesheet CLI",
	Long:  `A command line tool to interact with CodeHQ Timesheet ecosystem`,
}

func init() {
	DateLayout = configuration.Get().DateFormat

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	rootCmd.PersistentFlags().StringVarP(
		&OutputFormat,
		"output",
		"o",
		"json",
		"The output format. Can be either 'json', or 'table'",
	)

	rootCmd.PersistentFlags().StringVarP(
		&EmployeeId,
		"employee",
		"e",
		"",
		"The employee ID, should be firstname.lastname",
	)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
