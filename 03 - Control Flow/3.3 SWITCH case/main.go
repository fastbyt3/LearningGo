package main

import "fmt"

func main() {
	switch in := 1; {
	case in == 1:
		fmt.Println("No = 1")
		// fallthrough
	case in == 2:
		fmt.Println("No = 2")
		// fallthrough
	case in == 3:
		fmt.Println("No = 3")
		// fallthrough
	default:
		fmt.Println("exceeds 3")
	}
}
