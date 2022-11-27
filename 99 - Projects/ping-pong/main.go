package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Ball struct {
	hits int
}

func player(name string, table chan *Ball) {
	rand.Seed(time.Now().Unix())
	for {
		ball := <-table
		ball.hits++

		fmt.Printf("%s : %d\n", name, ball.hits)
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		table <- ball
	}
}

func main() {
	table := make(chan *Ball)

	go player("Ping", table)
	table <- new(Ball)

	go player("Pong", table)

	time.Sleep(1 * time.Second)
	<-table // game over
}
