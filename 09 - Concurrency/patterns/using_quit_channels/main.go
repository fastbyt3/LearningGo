package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan string)

	c := foo("Alice", quit)
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	fmt.Println(x)
	for i := x; i > 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "I'm Done"
	fmt.Println("Received on quit channel: ", <-quit)
}

func foo(msg string, quit chan string) chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
			case <-quit:
				quit <- "Kay bye!"
				return
			}
		}
	}()

	return c
}
