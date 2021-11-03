package main

import (
	"fmt"
	"net"
)

/*
----- This is a slow TCP Scanner -----

net.Dial(network, address string)

Close() -> sends FIN packet

--------------------------------------
*/

func main() {
	for i := 650; i < 660; i++ {
		addr := fmt.Sprintf("192.168.0.249:%d", i)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			// Port is closed
			fmt.Println("Port", i, "closed")
			continue
		}
		conn.Close()
		fmt.Println("- Port", i, "open")
	}
}
