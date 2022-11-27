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
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		}
	}()
	return c
}

func bar(in1, in2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-in1
		}
	}()

	go func() {
		for {
			c <- <-in2
		}
	}()

	return c
}
