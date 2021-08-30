package main

import "fmt"

type Config struct {
	Env   string
	proxy Proxy
}

type Proxy struct {
	addr string
	port int
}

func main() {

	c := &Config{
		Env: "DEBUG: TRUE",
		proxy: Proxy{
			addr: "http://localhost/",
			port: 80,
		},
	}

	fmt.Println(c)

	var c2 Config
	c2.Env = "DEBUG: FALSE"
	c2.proxy.addr = "http://localhost"
	c2.proxy.port = 9001

	fmt.Println(c2)
}
