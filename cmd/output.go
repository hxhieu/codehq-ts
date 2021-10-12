package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hxhieu/codehq-ts/models"
)

// OutputError serialises Go error to a predictable JSON string
func OutputErrorJson(err *error) {
	bytes, jsonErr := json.Marshal(models.Response{Error: (*err).Error()})
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(string(bytes))
}

// OutputResult serialises any object to JSON string
func OutputResultJson(any interface{}) {
	bytes, jsonErr := json.Marshal(models.Response{Data: any})
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(string(bytes))
}
