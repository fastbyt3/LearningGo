package main

import "fmt"
import "time"

func f1(c chan string) {
	c <- "Inside function f1"
	time.Sleep(2 * time.Second)
}

func f2(c chan string) {
	c <- "Inside function f2"
	time.Sleep(1 * time.Second)
}

func f3(c chan int, value int) {
	c <- value
}

func main() {
	var c chan string = make(chan string)
	go f1(c)
	go f2(c)
	fmt.Println(<-c)
	// time.Sleep(10 * time.Second)
	fmt.Println(<-c)

	v := 5
	fmt.Println("Original value of v:", v)
	chInt := make(chan int)
	go f3(chInt, 100)
	v = <-chInt
	time.Sleep(1 * time.Second)
	fmt.Println("Value of v recieved from channel:", v)
}
