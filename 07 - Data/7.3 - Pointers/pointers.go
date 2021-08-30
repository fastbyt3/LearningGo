package main

import "fmt"

func main() {
	x := 5
	fmt.Printf("Value of x = %d", x)
	// fmt.Printf("\nAddr of x = %x", &x)

	var pt *int
	// *pt = x -> Not possible
	pt = &x
	fmt.Printf("\nValue of pt = %d", *pt)
	// fmt.Printf("\nAddr of pt = %x", pt)

	*pt = 20 // changes x too as `pt` points to x
	fmt.Printf("\n\n----- Changed pt to 20 -----\n")
	fmt.Printf("Value of x = %d", x)
	// fmt.Printf("\nAddr of x = %x", &x)
	fmt.Printf("\nValue of pt = %d", *pt)
	// fmt.Printf("\nAddr of pt = %x", pt)

}
