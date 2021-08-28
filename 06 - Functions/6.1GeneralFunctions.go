package main

import "fmt"

func f1() {
	fmt.Println("hi")
}

func f2(s1 string) {
	fmt.Println(s1)
}

func f3(a int, b int) int {
	return a * b
}

func f4(a int, b int) (res int) {
	res = a * b
	return
}

func f5(a int, b int) (add int, sub int) {
	add = a + b
	sub = a - b
	return
}

func main() {
	f1()
	f2("My custom string")
	fmt.Println(f3(10, 9))
	fmt.Println(f4(10, 9))
	fmt.Println(f5(10, 9))
	x := 10
	y := 20
	swap(&x, &y)
	// fmt.Printf("\nAfter swap: %d and %d", x, y) : 20 10
	// After swapping w/ address even the orignal vals change
}

// Functions can take mem address as param too(like in C)
func swap(a *int, b *int) {
	fmt.Printf("Before swap: %d and %d", *a, *b)
	tmp := *a
	*a = *b
	*b = tmp
	fmt.Printf("\nAfter swap: %d and %d", *a, *b)
}
