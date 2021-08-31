package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var b = []int{6, 7, 8, 9}

	slice1 := a[1:4]
	fmt.Println(a, b, slice1)

	slice2 := make([]string, 10, 20)
	fmt.Printf("\nLength of slice2= %v\nCapacity of slice2= %v\n", len(slice2), cap(slice2))

	// creating slices using `new` keyword
	var newSlice = new([10]int)[0:5] // max size of slice = 10
	fmt.Println("Slice created with `new` gets filled with 0: ", newSlice)
	fmt.Println("Size of newSlice = ", len(newSlice))

	var x = []int{10, 20, 30, 40, 50, 60, 70}
	// x = append(x, 40)
	// fmt.Println("Adding 40 to slice X: ", x)

	var y = []int{5, 10, 15, 20}
	fmt.Println("Contents of Slice X: ", x)
	fmt.Println("Contents of slice Y: ", y)

	// copy slice into another slice : copy keyword
	// copy(dest, src)
	// returns no of elements copied
	// copy can copy only the elements of src that fit the max size of dest
	noOfElementsCopied := copy(x, y)
	fmt.Printf("\n---- copy(x, y) ----\n")
	fmt.Println("No of elements copied: ", noOfElementsCopied)
	fmt.Println("Contents of slice X: ", x)
	fmt.Println("Contents of slice Y: ", y)

	// Appending one slice to an existing one
	var z = []int{2, 4, 6, 8, 10}
	z = append(z, y...) // appends [5,10,15,20] to slice 'z'
	fmt.Println(z)
}
