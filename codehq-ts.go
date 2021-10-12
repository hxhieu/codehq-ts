package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	// // testDate := time.Date(2021, 9, 1, 0, 0, 0, 0, time.Local)
	// // test, err := services.Timesheet().GetWeeklyTimesheet("hugh.hoang", &testDate)
	// // OutputResult(models.Response{Data: test})
	// entry := models.Timesheet{
	// 	EmployeeId:  "hugh.hoang",
	// 	ProjectId:   100100,
	// 	Phase:       "99",
	// 	Code:        "ADMIN",
	// 	Description: "Hello 123",
	// 	Start:       time.Now(),
	// 	End:         time.Now().Add(15 * time.Minute),
	// }
	// newId, err1 := services.Timesheet().AddTimesheet(&entry)
	// services.Timesheet().Dispose()
	// if err1 != nil {
	// 	OutputError(&err1)
	// } else {
	// 	OutputResult(newId)
	// }
	fmt.Println("Please select table.")
	t := prompt.Input("codehq-ts > ", completer)
	fmt.Println("You selected " + t)
}
