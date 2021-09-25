package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		messages <- 1
	}()
	go func() {
		time.Sleep(time.Second * 2)
		messages <- 2
	}()
	go func() {
		time.Sleep(time.Second * 1)
		messages <- 3
	}()
	for i := 0; i < 3; i++ {
		fmt.Println(<-messages)
	}
}

// Reference: https://nathanleclaire.com/blog/2014/02/21/how-to-wait-for-all-goroutines-to-finish-executing-before-continuing-part-two-fixing-my-ooops/

// messages <- 1 ===> data 1 is stored at memory loc message : Channel : more abt it in next module
