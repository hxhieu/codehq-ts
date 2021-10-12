/*
Copyright © 2021 Hugh Hoang <hugh.hoang@codehq.nz>
For internal use only
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hxhieu/codehq-ts/configuration"
	"github.com/hxhieu/codehq-ts/models"
	"github.com/spf13/cobra"
)

var EmployeeId string
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

// OutputError serialises Go error to a predictable JSON string
func OutputError(err *error) {
	bytes, jsonErr := json.Marshal(models.Response{Error: (*err).Error()})
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(string(bytes))
}

// OutputResult serialises any object to JSON string
func OutputResult(any interface{}) {
	bytes, jsonErr := json.Marshal(models.Response{Data: any})
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(string(bytes))
}
