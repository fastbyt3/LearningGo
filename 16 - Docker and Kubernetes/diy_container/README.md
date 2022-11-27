# DIY container (??)

> https://www.youtube.com/watch?v=Utf-A4rODH8&ab_channel=ContainerCamp

### Namespaces

- visibility of container

- created with syscalls:
    - UTS: Unix Timesharing System
    - PIDs
    - Mounts
    - Network
    - UIDs
    - IPC

### Rootless containers

```bash
$ go run main.go run echo hello
Executing [echo hello] as PID 1000
panic: fork/exec /proc/self/exe: operation not permitted

goroutine 1 [running]:
main.execWithoutFail(...)
	/home/fastbyte/Documents/go_journeys/16 - Docker and Kubernetes/diy_container/main.go:50
main.run()
	/home/fastbyte/Documents/go_journeys/16 - Docker and Kubernetes/diy_container/main.go:33 +0x312
main.main()
	/home/fastbyte/Documents/go_journeys/16 - Docker and Kubernetes/diy_container/main.go:13 +0x4d
exit status 2
```

