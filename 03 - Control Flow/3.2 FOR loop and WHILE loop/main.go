package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	min := 0
	i := 10
	for i > min {
		fmt.Println(i)
		i--
	}
}
