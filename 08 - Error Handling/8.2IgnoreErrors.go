package main

import "fmt"
import "errors"

func throwError() (int, error) {
	return 1, errors.New("This is a error")
}

func main() {
	r, _ := throwError()
	fmt.Println("Error was ignored")
	fmt.Println("r = ", r)
}
