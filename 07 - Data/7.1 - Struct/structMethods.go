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

func (conf Config) ConfInfo() string {
	fmt.Println("Env: ", conf.Env)
	fmt.Println("Proxy: ", conf.proxy)

	return "---------------------"
}

func main() {
	c := Config{
		Env: "DEBUG: TRUE",
		proxy: Proxy{
			addr: "http://localhost/",
			port: 8080,
		},
	}

	fmt.Println(c.ConfInfo())
}
