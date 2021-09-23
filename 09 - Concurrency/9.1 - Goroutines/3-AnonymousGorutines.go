package main

import "fmt"

import "time"

func foo() {
	fmt.Println("This is foo!")
}

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Print(i, " ")
		}
	}()
	time.Sleep(10 * time.Millisecond) // -> 0 1 2 3 4 This is foo!
	foo()
}

// Without time delay :
// 		This is foo!
// 		0
