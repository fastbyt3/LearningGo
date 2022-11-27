# Race conditions

## When does race condition happen

- Race conditions occur when 2/more concurrent goroutines(contexts) perform READ/WRITE operation on same variable leading to unexpected results
- Shared data between goroutines is potential race condition as critical section(modification of shared var) may happen at the same time

## Example of race 

```go
var counter uint32
var wg sync.WaitGroup

func main() {
    wg.Add(3)

    go inc()
    go inc()
    go inc()
    wg.Wait()
    fmt.Println(counter)
}

func inc() {
	for _ : range 3000 {
        counter++
    }
    wg.Done()
}
```

`counter` is the shared resource which each `inc()` function increments a random amt of times
Since there is no guarentee whether the 2 goroutines wont access the var at once we might get unexpected results
like at one point counter = 12 and 2 gorountines access it and try incrementing it once
Expected op = 14 but we might get op=13 as both perform `12+1`

## Ways to handle data race

### Blocking using channels

- Blocking by waiting for message from channel is also viable solution
- Similar to token passing where a token allows access to a resource

```go
// In G1
channel <- true

// In G2
for msg := range <-channel {
    if msg { break; }
}
```

### Atomic operations

- ensuring each operation that happens in Critical section is atomic
- Go provides `sync/atomic` to achieve synchronization during concurrency
- For example: Using `atomic.AddUint32`

```go
var counter uint32
var wg sync.WaitGroup

func main() {
    wg.Add(3)

    go inc()
    go inc()
    go inc()
    wg.Wait()
    fmt.Println(counter)
}

func inc() {
	for _ : range 3000 {
        atomic.AddUint32(counter, 1) // Ensures atomic increment
    }
    wg.Done()
}
```

### Mutex locks

- Mutual Exclusion lock: grants single-point authority on a shared resource
- `var mutex sync.Mutex`, it has 2 main functions: `mutex.Lock()` and `mutex.Unlock()`
- `mutex.RLock()` and `mutex.RUnlock()` are Read locks that multiple goroutines can possess at same time. 
- Multiple goroutines can acquire RLock when one goroutine has RLock but if a Goroutine requires Lock it must wait till all RLocks are completed.
- Similarly when Lock is currently active, goroutnes requesting RLock will have to wait. 
- working of mutex:
    ```
    Request Lock -> Acquire Lock -> Modify data -> Release lock
    ```

```go
var counter uint32
var wg sync.WaitGroup
var m sync.mutex

func main() {
    wg.Add(3)

    go inc()
    go inc()
    go inc()
    wg.Wait()
    fmt.Println(counter)
}

func inc() {
    m.Lock()
	for _ : range 3000 {
        counter++
    }
    m.Unlock()
    wg.Done()
}
```

### Atomic vs Mutex

- Atomic uses CPU instructions but Mutex implements locking
- Atomic is faster for updation of integers
- Mutex are preferred handling concurrency in complex data structures as its the only choice

### Another way -> utilize go channels

- Sharing memory by communicating
- Using send and receive channels to start/pause execution of goroutines