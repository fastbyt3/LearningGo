package main

import "fmt"

func main() {
	// var smallMap = make(map[string]string)
	// smallMap["Alice"] = "Bob"
	// smallMap["Charlie"] = "Damion"
	// fmt.Println(smallMap)
	// delete(smallMap, "Alice")
	// fmt.Println(smallMap)

	var rollNo = map[string]int{
		"Aldeesh": 1,
		"Brandon": 2,
		"Charlie": 3,
	}
	// fmt.Println("Roll number of Aldeesh: ", rollNo["Aldeesh"])
	// rollNo["Charlie"] = 99
	// fmt.Println(rollNo)

	for k, v := range rollNo {
		fmt.Printf("Name: %s\tRollNo: %d\n", k, v)
	}

	// Checking if an element exists in the map
	c, ok := rollNo["Aldeesh"]
	fmt.Println(c, ok)
	c, ok = rollNo["Zed"]
	fmt.Println(c, ok)

}
