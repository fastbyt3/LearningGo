package main

import "fmt"

// def anonymous function
var (
	area = func(l int, b int) int {
		return l * b
	}
)

func main() {
	area := area(5, 2)
	fmt.Println(area)
}
