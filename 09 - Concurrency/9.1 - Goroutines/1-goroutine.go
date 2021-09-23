package main

import (
	"fmt"
	"time"
)

func printSomething(msg string) {
	fmt.Println(msg)
}

func main() {
	go printSomething("This is what is being run concurrently : GoRoutine")
	printSomething("Testing next line")
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Exitting....")
	// fmt.Scanln()
}
