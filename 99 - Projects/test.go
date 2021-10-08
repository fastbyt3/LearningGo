package main

import "fmt"

func foo(s string, c chan int) {
	c <- len(s)
}

func main() {
	c := make(chan int)
	go foo("this", c)
	go foo("routines", c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
