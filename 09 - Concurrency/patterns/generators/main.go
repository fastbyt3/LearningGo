package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("vim-go")

	// 	ch := foo("LOLOLOL")
	//
	// 	for i := 0; i < 5; i++ {
	// 		fmt.Println("Recv: ", <-ch)
	// 	}

	p1 := foo("Alice")
	p2 := foo("Bobby")

	// Synchronous nature
	// Bobby can send value only if Alice sends value
	// else Bobby remains blocked
	for i := 0; i < 5; i++ {
		fmt.Println(<-p1)
		fmt.Println(<-p2)
	}

	fmt.Print("Im quitting....")
}

func foo(msg string) chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}
