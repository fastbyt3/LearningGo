package main

import "fmt"

func main() {
	res := fib(10)
	fmt.Println(res)
}

func fib(n int) int {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		fmt.Println(x)
		x, y = y, x+y
	}

	return x
}
