package main

import (
	"errors"
	"fmt"
)

func factorial(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("Negative number was given as input")
	} else {
		fact := 1
		for i := n; i >= 0; i-- {
			fact *= i
		}
		return fact, nil
	}
}

func main() {
	n := 10
	res, err := factorial(n)
	if err != nil {
		fmt.Printf("factorial of %d = %d", n, res)
	} else {
		fmt.Println(err)
	}
}
