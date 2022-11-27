package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("vim-go")

	go foo("This is boring...")

	fmt.Println("Listening....")
	time.Sleep(2 * time.Second)
	fmt.Println("I'm done!!")
}

func foo(m string) {
	for i := 0; ; i++ {
		fmt.Println(m)

		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
