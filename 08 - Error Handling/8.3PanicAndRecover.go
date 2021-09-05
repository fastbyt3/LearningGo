package main

import "fmt"

func execPanic() {
	defer recoverFromPanic()
	panic("Panic time!!!")
	fmt.Println("End: execPanic function")
}

func recoverFromPanic() {
	e := recover() // e = panic message : "Panic time!!!"
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("End: recover() function")
}

func main() {
	fmt.Println("Start MAIN function")
	execPanic()
	fmt.Println("End: MAIN function")
}
