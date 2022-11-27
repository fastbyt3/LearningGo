package main

import "fmt"

func main() {
	const n = 100

	leftmost := make(chan int)

	left := leftmost
	right := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go foo(left, right)
		left = right
	}

	// Initialize first value
	go func(c chan int) {
		c <- 1
	}(right)

	fmt.Println(<-leftmost)
}

func foo(left, right chan int) {
	left <- 1 + <-right
}
