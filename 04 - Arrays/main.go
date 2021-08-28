package main

import "fmt"

func standardArrays() {
	// var a = new([5]int64)
	// i := 0
	// len := 5
	// for i < len {
	// 	fmt.Scanf("%d", &a[i])
	// 	i++
	// }	Dont know why but skips from 1 to 2 without taking input
	a := [5]int64{1, 2, 3, 4, 5}
	// fmt.Println("Displaying elements in arr:")
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(a[i])
	// }
	fmt.Println(len(a)) // print len of array = 5

	arr2d := [2][2]int{
		{1, 1},
		{1, 0},
	}

	fmt.Println(arr2d)

	// Arr of strings
	strArr := [3]string{
		"Aldeesh",
		"Mike",
		"Bob",
	}
	fmt.Printf("%s\n", strArr) // [Aldeesh Mike Bob]
	fmt.Printf("%q\n", strArr) // ["Aldeesh" "Mike" "Bob"] : %q can be used for inserting quotes
}

func slices() {
	arr := [5]string{
		"Alice",
		"Bob",
		"Charlie",
		"Danny",
		"Eugene",
	}

	// slice := make([]string, 3)
	slice := make([]string, 3, 5)
	slice = arr[0:2]

	// fmt.Println(arr)
	fmt.Printf("%q\n", slice)
	fmt.Println(cap(slice))
}

func usingRange() {
	nums := [5]int{
		1, 2, 3, 4, 5,
	}
	fmt.Printf("Index\t|\tValue\n")
	for i, v := range nums {
		fmt.Printf("%d\t-\t%d\n", i, v)
	}
}

func main() {
	// standardArrays()
	// slices()
	usingRange()
}
