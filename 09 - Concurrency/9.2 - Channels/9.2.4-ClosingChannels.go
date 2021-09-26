package main

import "fmt"

func ClosingAChannel(ch chan string, s string) {
	ch <- s
	close(ch)
}

func main() {
	ch := make(chan string)
	go ClosingAChannel(ch, "Hi there")

	retValue, condition := <-ch // recive data from channel

	if !condition {
		fmt.Println("Channel is still open!")
	}

	fmt.Println("Returned value: ", retValue)
}
