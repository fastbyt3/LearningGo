# Goroutines

> A Goroutine is a func that can run _concurrently_
> Goroutines are units of Execution in Go

- Go routines are lightweight threads in Go

- To invoke a goroutine, write `Go` before the function call

```go
go func1()
```

- All programs have atleast 1 goroutine that is implictly defined : **Main function**

- Go doesnt wait for a GoRoutine to be finished, it returns the execution to next line
- If there is no further instructions Go would finish the program which would result in not executing the goroutine
- To fix this we use `time.Sleep()` to introduce some delay

```go
func main() {
	go printSomething("This is what is being run concurrently : GoRoutine")
	printSomething("Testing next line")
	fmt.Println("Exitting....")
	// fmt.Scanln()
}
```

In this case only "Testing next line \n Exitting...." is printed. To ensure that the second Goroutine runs lets add some delay

```go
func main() {
	go printSomething("This is what is being run concurrently : GoRoutine")
	printSomething("Testing next line")
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Exitting....")
	// fmt.Scanln()
}
```

Output:

```
Testing next line
This is what is being run concurrently : GoRoutine
Exitting....
```

- In `2-goroutine.go` a different example is used
- In that program, it shows that the placement of the delay is important to get the results we need

- The third program uses Goroutines using anonymous functions

### Why use GoRoutines

- If the problem task has segments that can be done individually then Goroutines can make it faster
- making multiple requests to endpoints
- Multithread workload


---