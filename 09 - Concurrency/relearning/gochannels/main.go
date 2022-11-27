package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)

	go foo("This is boring...", c)

	fmt.Println("Listening....")

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("I'm done!!")
}

func foo(m string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%d %s", i, m)

		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
