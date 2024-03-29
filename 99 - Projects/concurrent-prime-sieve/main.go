package main

import "fmt"

// Output series of numbers from 2...inf
func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// Remove non-primes from input channel (multiples of prime)
// Send to out channel
func filter(in, out chan int, prime int) {
	for {
		i := <-in

		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	ch := make(chan int)

	go generate(ch)

	const n = 50
	for i := 0; i < n; i++ {
		prime := <-ch
		fmt.Println(prime)

		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}

}
