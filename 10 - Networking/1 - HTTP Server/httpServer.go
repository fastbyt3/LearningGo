package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "!!! This is a WebServer running on Golang !!!")
	})

	// Handling diff routes
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Yo!!\nYou have reached '/test'\nSee ya ✌")
	})

	http.HandleFunc("/hidden", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Congratss!!!\nYou have found the hidden page!!\n✨✨✨")
	})

	fmt.Println("[\\] Starting the webserver....")
	http.ListenAndServe(":9999", nil)
}
