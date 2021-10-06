package main

import (
	"fmt"
	"sync"
	"time"
)

func foo(wg *sync.WaitGroup) {
	fmt.Println("Doing foobar work....")
	time.Sleep(3 * time.Second)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go foo(&wg)
	wg.Wait()
	fmt.Println("[+] Goroutine has finished...")
}
