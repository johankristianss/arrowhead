package main

import (
	"encoding/json"
	"fmt"

	arrowhead "github.com/johankristianss/arrowhead/pkg/arrowhead"
)

type Car struct {
	Brand string `json:"brand"`
	Color string `json:"color"`
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	framework, err := arrowhead.CreateFramework()
	checkError(err)

	// Buiid the request
	params := arrowhead.EmptyParams()
	car := Car{Brand: "Toyota", Color: "Red"}
	carJSON, err := json.Marshal(car)
	checkError(err)
	params.Payload = carJSON

	// Send the request
	res, err := framework.SendRequest("create-car", params)
	checkError(err)

	// Fetch cars
	res, err = framework.SendRequest("get-car", arrowhead.EmptyParams())
	checkError(err)

	// Print the response
	cars := []Car{}
	err = json.Unmarshal(res, &cars)
	checkError(err)
	for _, car := range cars {
		fmt.Println(car.Brand, car.Color)
	}
}
