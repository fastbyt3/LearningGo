package main

import "fmt"

func main() {
	// x := 2
	// x := 0
	x := -1

	if x > 0 {
		fmt.Println("Positive number")
	} else if x < 0 {
		fmt.Println("Negative number")
	} else {
		fmt.Println("The number is zero")
	}
}
