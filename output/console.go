package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hxhieu/codehq-ts/models"
	"github.com/olekukonko/tablewriter"
)

// ConsoleErrorJson serialises Go error to a predictable JSON string
func ConsoleErrorJson(err *error) {
	bytes, _ := json.Marshal(models.Response{Error: (*err).Error()})
	// Send output
	fmt.Println(string(bytes))
	// Quit the process successfully
	// because the consumer would need to parse the result
	// if it is an error or not
	os.Exit(0)
}

// ConsoleResultJson serialises any object to JSON string
func ConsoleResultJson(data interface{}) {
	bytes, _ := json.Marshal(models.Response{Data: data})
	// Send output
	fmt.Println(string(bytes))
}

// ConsoleResultTable render result as ASCII table
func ConsoleResultTable(data [][]string, header []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	for _, v := range data {
		table.Append(v)
	}
	// Send output
	table.Render()
}
