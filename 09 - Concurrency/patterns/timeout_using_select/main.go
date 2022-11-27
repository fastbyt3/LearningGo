package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	// 	timeoutEachMsg()
	timeoutEntireConversation()
	fmt.Println(time.Since(start))
}

func timeoutEntireConversation() {
	c := foo("Alice")
	timeout := time.After(5 * time.Second)

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("Exiting cos you took long!")
			return
		}
	}
}

func timeoutEachMsg() {
	c := foo("Alice")

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("Exiting cos you took long!")
			return
		}
	}
}
func foo(msg string) chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			x := time.Duration(rand.Intn(2e3))
			fmt.Printf("Sleep for %d ms\n", x)
			time.Sleep(x * time.Millisecond)
		}
	}()

	return c
}
