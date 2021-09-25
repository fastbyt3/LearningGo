# Channels in Golang

### What are channels

> Channels are a way for goroutines to communicate with each other

- defined using the keyword : `chan`
- Channels help communication between goroutines as data can't be simply passed between goroutines

### How do channels work?

```golang
c <- "message"
```

- "message" gets sent to channel

```golang
msg := <- c
```

- ^^ means a msg is recieved from the channel


### Creating channels

```golang
channel1 := make(chan string)	// channel that carries string
channel1 := make(chan int)		// channel that carries int only
```