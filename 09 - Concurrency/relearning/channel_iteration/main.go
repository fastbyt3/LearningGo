package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("vim-go")

	channel := make(chan string)

	go foo(channel)

	for {
		msg, ok := <-channel
		if !ok {
			break
		}
		fmt.Println(msg)
	}

}

func foo(channel chan string) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprintf("%d. Score is %d", i, score)
	}

	close(channel)
}
