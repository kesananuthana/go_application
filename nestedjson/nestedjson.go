package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string `json:"city"`
	Pincode int    `json:"pincode"`
}

type Person struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

func main() {
	jsonData := `[
		{
			"id": 1,
			"name": "Nutana",
			"address": {
				"city": "Hyderabad",
				"pincode": 500001
			}
		},
		{
			"id": 2,
			"name": "Mahi",
			"address": {
				"city": "Warangal",
				"pincode": 506002
			}
		}
	]`

	var person []Person
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Print(person[0].Name)
	fmt.Print("\n", person[0].Address.City)

	fmt.Print("\n", person[1].Name)
	fmt.Print("\n", person[1].Address.City)
}
