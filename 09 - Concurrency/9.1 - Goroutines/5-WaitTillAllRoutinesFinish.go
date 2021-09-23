package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"time"
)

func get(URL string) {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	sb := string(body)
	fmt.Println(sb)
}

func main() {
	fmt.Println(runtime.NumGoroutine())
	go get("https://fastbyt3.github.io/")
	go get("https://pronoun.is/he")
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(20 * time.Second)
	fmt.Println(runtime.NumGoroutine())

	// i := 1
	// for runtime.NumGoroutine() != 1 {
	// 	fmt.Println("Delay: ", i)
	// 	fmt.Println("No of goroutines: ", runtime.NumGoroutine())
	// 	time.Sleep(5 * time.Second)
	// 	i++
	// }
	// time.Sleep(10 * time.Millisecond)
}
