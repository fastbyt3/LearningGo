package main

import (
	"fmt"
)

func counter(out chan<- int) {
	for i := 1; i < 11; i++ {
		out <- i
	}

	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for {
		n, ok := <-in

		if !ok {
			close(out)
			return
		}

		out <- n * n
	}
}

func printer(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}

func main() {
	fmt.Println("vim-go")

	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)

	printer(squares)
}
