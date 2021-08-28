package main

import "fmt"

func main() {
	l := 5
	b := 7

	// similar to anonymous fn but can access variables within the function in which it is defined
	func() {
		area := l * b
		fmt.Println(area)
	}()
}
