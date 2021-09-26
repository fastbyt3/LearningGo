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

### Directional Channels

- Channels are _unidirectional_ in nature when declared like, 

```golang
c := make(chan int)
```

- Directional channels allow setting a channel either to only recieve or send data

```golang
ch := make(chan<- int)	// send only

ch := make(<-chan int)	// recv-only
```

### Closing a Channel

```golang
close(ch)
```

- `close()` ret a _bool_ val which can be used to check if the channel is closed or not

### Reference

1. https://golangdocs.com/channels-in-golang
2. https://golangr.com/channels/

---