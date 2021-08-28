package main

import (
	"bufio"
	"fmt"
	"os"
)

func usingFmt() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello there, %s", name)
}

func usingBufIO() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter anything: ")

	// str1 := reader.ReadString('\n') doesnt work cos reader.ReadString returns two vals
	str1, _ := reader.ReadString('\n')

	fmt.Print("Your anything is : " + str1)
	fmt.Print("\nenter one more? ")
	str2, _ := reader.ReadString('\n')
	fmt.Print("Nope cant process " + str2)
}

func main() {
	// usingFmt()
	usingBufIO()
}
