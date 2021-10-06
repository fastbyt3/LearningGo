# WaitGroups

> Waitgroup is a blocking mechanism that blocks when none of the GoRoutines which is "inside" that group has been executed
> If a goroutine has finished exec -> unblocks the group

- Belongs to `sync` package
- creation:

```golang
var wg sync.WaitGroup
```

- to show that a goroutine func is using this WaitGroup pass that as _pointer_ arg 
- And in the calling func(in eg. it is `main`) use : 

```golang
wg.Add(1)	// Adds 1 goroutine : Counter
wg.Wait()	// Waits till wg.Done() is called
```

- To indicate that the Goroutine has finished execution

```golang
wg.Done()
```

**IMPORTANCE**

GoLang Waitgroups are important because they allow a goroutine to block the thread and execute it. Without it, we need to manually sleep the main thread to let the goroutines execute.

----