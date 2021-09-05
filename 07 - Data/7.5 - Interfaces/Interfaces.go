package main

import "fmt"

/*
----------------- Go : Interfaces ----------------------------------

An interface is an abstract concept which enables polymorphism in Go

Declaration: `type interfaceName interface{}`

Interfaces can be used when Polymorphism is needed in Go
--------------------------------------------------------------------
*/

type Student interface {
	PrintDetails()
}

// Create a type for the interface
type Stud struct {
	name string
	dept string
	year int
}

func (s Stud) PrintDetails() {
	fmt.Println("Student name: ", s.name)
	fmt.Println("Department: ", s.dept)
	fmt.Println("Current year: ", s.year)
}

func PrintInterface(i interface{}) {
	fmt.Printf("Interface: %t\n", i)
}

func main() {
	var emptyInterface interface{} // Emtpy interface
	fmt.Println("This is an empty interface: ", emptyInterface)

	var a interface{} = "some strings"
	// fmt.Println(a)
	PrintInterface(a)

	var s1 Student
	s1 = Stud{"Aldeesh", "CompSci", 3}
	s1.PrintDetails()
}
