package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	// define a struct
	type Person struct {
		name string `json:"name"`
		age  int    `json:"age"`
		role string `json:"role"`
	}
	var person Person

	// create a json string
	// p1 := `{"name": "fastbyte", "age": 19, "role": "student"}`
	// fmt.Println("p1:", p1)

	//read from file
	file, e := ioutil.ReadFile("test.json")
	if e != nil {
		fmt.Println(e)
	}
	// err := json.Unmarshal([]byte(p1), &person)

	err := json.Unmarshal([]byte(file), &person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("person:", person)
	// fmt.Printf("\n----- Person details -----\nName: %s\nAge: %d\nRole: %s\n-----------", person.name, person.age, person.role)
}
