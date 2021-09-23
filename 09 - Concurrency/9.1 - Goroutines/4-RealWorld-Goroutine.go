package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func get(URL string) {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	sb := string(body)
	fmt.Println(sb)
}

func main() {
	go get("https://fastbyt3.github.io/")
	go get("https://pronoun.is/he")
	// get("https://jsonplaceholder.typicode.com/posts")
	// fmt.Println("exitting...")
	time.Sleep(20 * time.Second)
}

// Reference -> https://blog.logrocket.com/making-http-requests-in-go/
