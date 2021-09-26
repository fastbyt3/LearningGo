package main

import "fmt"

func f1(ch chan<- int, v int) {
	ch <- v
	// fmt.Println(<-ch) // Errors since it cant recieve from a send-only channel
}

func main() {
	ch := make(chan<- int) // send-only channel
	go f1(ch, 42)
	go f1(ch, 41)
	go f1(ch, 40)
}
