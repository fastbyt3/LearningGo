package main

import "fmt"
import "time"

func f1(done chan bool) {
	fmt.Println("The goroutine is")
	fmt.Println("doing something")
	fmt.Println("....Heheh  ༼ つ ◕_◕ ༽つ  ....")

	time.Sleep(2 * time.Millisecond)

	done <- true
}

func main() {
	done := make(chan bool)

	go f1(done)

	fmt.Println("Status of goroutine: Done =", <-done)
}
