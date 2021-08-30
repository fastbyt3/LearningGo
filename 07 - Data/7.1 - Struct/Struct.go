package main

import "fmt"

type Person struct {
	name string
	job  string
}

func learn() {
	var p1 Person

	p1.name = "Aldeesh"
	p1.job = "Student"

	// fmt.Printf("Name: %s\nProfession: %s\n", p1.name, p1.job)

	fmt.Println(p1)

	var p2 = &Person{name: "Alice", job: "Painter"}
	fmt.Println(p2)

	p3 := p2
	p3.name = "Bob"
	p3.job = "Teacher"

	fmt.Println(p3)

	p4 := p1
	p4.name = "Charlie"
	fmt.Println(p4)
}

type house struct {
	noRooms int
	price   int
	city    string
}

func diy() {
	var house1 house
	var house2 house
	house1.noRooms = 3
	house1.price = 10000
	house1.city = "Chennai"

	house2.noRooms = 2
	house2.price = 7000
	house2.city = "Selaiyur"

	fmt.Println("Available housing options: ")
	fmt.Printf("1. House1: \n\tavailable rooms: %d\n\tPrice: %d\n\tCity: %s", house1.noRooms, house1.price, house1.city)
	fmt.Printf("\n2. House2: \n\tavailable rooms: %d\n\tPrice: %d\n\tCity: %s", house2.noRooms, house2.price, house2.city)
}

func main() {
	// diy()
	learn()
}
