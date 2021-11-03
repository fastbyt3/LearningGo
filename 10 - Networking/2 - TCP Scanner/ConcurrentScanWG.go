package main

import (
	"fmt"
	"net"
	"sync"
)

/*

Concurrent Port Scanner

- uses Goroutines with WaitGroups
- This enables Synchronized scanning

*/

func main() {
	var wg sync.WaitGroup
	// openPorts := []int{}
	for i := 1; i < 1024; i++ {
		wg.Add(1)
		go func(pno int) {
			defer wg.Done() // Ensure that the waitgroup is closed
			addr := fmt.Sprintf("127.0.0.1:%d", pno)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				// Port is closed
				// fmt.Println("Port", i, "closed")
				return
			}
			conn.Close()
			// openPorts = append(openPorts, pno)
			fmt.Println("- Port", i, "open")
		}(i)
	}
	// fmt.Println(openPorts)
}
