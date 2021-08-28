# Hello World - Go

My very first program in go!! Excited✨!!

Running go applications:

```bash
go run main.go

# In case we want to run multiple files
go run .
``` 

Building go applications - creates binary

```bash
go build
```

## Program

- create a `main.go` file
- code:

```go
package main
import "fmt"

func main(){
	fmt.Println("Hello, World!! This is my very first program!")
}
```

## Setting up Go in sublime

- Install `Golang Build` 

- From Preferences -> Package Settings -> Golang config -> User settings, add the `PATH` and `GOPATH` var

```json
{
    "PATH": "",
    "GOPATH": ""
}
```

- For Go formatting, install `Gofmt`

---
