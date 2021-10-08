# Error handling

- Go's built-in error type uses _Interface_

```golang
type error interface {
	Error() string
}
```

- Unlike most programming languages Golang doesn't use `try`, `catch`, `finally` methods to handle errors
- In Golang, most functions return a error value if `nil` => no error

```golang
if res, err := foo(); err!=nil {
	// Error handling code
	// ....
}
```

- Since we know how the errors get defined we can create our own error:

```golang
type MyError string
func (e MyError) Error() string {
	return string(e)
}
```

- 