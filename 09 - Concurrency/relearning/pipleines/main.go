package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 1; x <= 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for {
			num, ok := <-naturals
			if !ok {
				close(squares)
				return
			}
			squares <- num * num
		}
	}()

	for {
		sq, ok := <-squares
		if !ok {
			return
		}
		fmt.Println(sq)
	}
}
