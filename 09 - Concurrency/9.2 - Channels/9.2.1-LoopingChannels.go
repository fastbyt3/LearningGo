package main

import "fmt"

func loopChannel(ch chan int, v int) {
	ch <- v
	ch <- v * 2
	ch <- v * 4
	ch <- v * 8
	close(ch) // If not used, msg: "fatal error: all goroutines are asleep - deadlock!" is issued
}

func main() {
	ch := make(chan int)

	go loopChannel(ch, 5)

	// Using `range` we can iterate over the values sent via the Channel
	for i := range ch {
		fmt.Println(i)
	}
}
