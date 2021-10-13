package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/hxhieu/codehq-ts/models"
	"github.com/olekukonko/tablewriter"
)

// OutputErrorJson serialises Go error to a predictable JSON string
func OutputErrorJson(err *error) {
	bytes, jsonErr := json.Marshal(models.Response{Error: (*err).Error()})
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	// Send output
	fmt.Println(string(bytes))
}

// OutputResultJson serialises any object to JSON string
func OutputResultJson(data interface{}) {
	bytes, jsonErr := json.Marshal(models.Response{Data: data})
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	// Send output
	fmt.Println(string(bytes))
}

// OutputResultTable render result as ASCII table
func OutputResultTable(data [][]string, header []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	for _, v := range data {
		table.Append(v)
	}
	// Send output
	table.Render()
}
