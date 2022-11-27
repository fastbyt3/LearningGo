package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := bar(foo("Ann"), foo("Bob"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("Quitting....")
}

func foo(m string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", m, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func bar(in1, in2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-in1:
				c <- s
			case s := <-in2:
				c <- s
			}
		}
	}()

	return c
}
