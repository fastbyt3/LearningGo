package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	age  int
}

func getDetails(c chan Person, p Person) {
	c <- p
	time.Sleep(1 * time.Second)
}

func main() {
	p := Person{
		"Aldeesh",
		19,
	}
	c := make(chan Person)

	go getDetails(c, p)

	// time.Sleep(1 * time.Second)

	fmt.Println("Name of person:", (<-c).name)
	// fmt.Println("Age of person:", (<-c).age) : Errors since the channel is closed as the work is over
}
