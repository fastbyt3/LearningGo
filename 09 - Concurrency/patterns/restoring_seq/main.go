package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	msg  string
	wait chan bool
}

func main() {
	c := bar(foo("Ann"), foo("Bob"))

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.msg)

		msg2 := <-c
		fmt.Println(msg2.msg)

		msg1.wait <- true
		msg2.wait <- true
	}

	fmt.Println("Quitting....")
}

func foo(m string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", m, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			<-waitForIt
		}
	}()
	return c
}

func bar(in1, in2 <-chan Message) <-chan Message {
	c := make(chan Message)

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
