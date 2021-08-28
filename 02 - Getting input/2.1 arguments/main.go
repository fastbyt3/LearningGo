package main

import (
	"fmt"
	"os"
)

func main() {
	argsSupplied := os.Args[1:]
	fmt.Println("Array of all args: ")
	fmt.Println(argsSupplied)
	fmt.Println("Arg at index 2: ")
	fmt.Println(argsSupplied[1])
}
