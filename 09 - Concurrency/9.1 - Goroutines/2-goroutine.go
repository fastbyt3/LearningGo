package main

import (
	"fmt"
	"time"
)

func foo() {
	// var i int
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
		time.Sleep(10 * time.Millisecond) // -> 0 0 1 1 2 2 3 3 4 4
	}
}

func main() {
	go foo()
	// time.Sleep(10 * time.Millisecond)	// -> 0 1 2 3 4 0 1 2 3 4
	foo()
	// time.Sleep(10 * time.Millisecond) // -> 0 1 2 3 4 0 1 2 3 4
}
