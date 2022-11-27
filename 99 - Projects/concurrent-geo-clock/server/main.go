package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("Concurrent Clock Server")

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(con)
	}
}

func handleConnection(con net.Conn) {
	defer con.Close()

	for {
		_, err := io.WriteString(con, time.Now().Format("10:30:15\n"))

		if err != nil {
			log.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
