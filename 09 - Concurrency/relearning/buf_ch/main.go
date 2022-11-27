package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")
	ch := make(chan int, 3)

	defer close(ch)

	go func() {
		for i := 1; i < 11; i++ {
			ch <- i
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(len(ch))
}
