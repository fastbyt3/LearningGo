package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	vari := 10

	varj := retNewDataObj()
	*varj = 100
	fmt.Println(*varj)

	fmt.Println(vari)

	// 	foo(vari)
	foo(&vari)

	fmt.Println(vari)
}

func retNewDataObj() *int {
	return new(int)
}

func foo(vari *int) {
	*vari = 100
	fmt.Println(*vari)
}

// func foo(vari int) {
// 	vari = 100
// 	fmt.Println(vari)
// }
