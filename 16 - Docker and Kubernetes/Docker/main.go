package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Author    string
	Msg       string
	Timestamp time.Time
}

func main() {
	fmt.Println("Initializing server... at http://localhost:8080/")
	http.HandleFunc("/", IndexPagePutMsg)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func IndexPagePutMsg(resp http.ResponseWriter, req *http.Request) {
	Message := Message{
		Author:    "fastbyte",
		Msg:       "Hello there",
		Timestamp: time.Now(),
	}

	json.NewEncoder(resp).Encode(Message)

	fmt.Println("Returned: ", Message)
}
