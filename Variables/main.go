package main

import "fmt"

func main() {
	var s1 string = "This is a string"
	fmt.Println(s1)

	s2 := "Is this a string??"
	fmt.Println(s2)

	var i1 int = 1
	fmt.Printf("i1 = %d\n", i1)
	// i1 := 2 -> will throw error cos i1 is already defined
	// fmt.Println(i1)

	i2 := 2 // This works cos i2 is new var
	fmt.Println(i2)

	var (
		f1 float32
		f2 float64
		b  bool
	)
	f1, f2, b = 6.9, 4.20, true
	fmt.Println(f1, f2, b)
}
