package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// Program: Infinitely ping a set of URLs passed in as ENV var

func main() {
	urls := strings.Split(os.Getenv("URLS"), ",")
	if len(urls) == 1 && urls[0] == "" {
		panic("Environment var URLS not set")
	}

	for _, url := range urls {
		go pingURL(url)
	}

	// Create a channel
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	// we can use someVar := <-stop if we are getting some val from that channel
	<-stop // allows main() to proceed after receiving termination from all channels

	log.Println("Terminating....")
}

// Send a GET request to URL and validate status using error
func pingURL(url string) {
	for {
		log.Println("Pinging ", url)
		_, err := http.Get(url)

		if err != nil {
			log.Println("Error pinging ", url, " Error: ", err)
		}

		time.Sleep(5 * time.Second)
	}
}
