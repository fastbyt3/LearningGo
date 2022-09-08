# LearningGo
Documenting my learning process in Go

## Initializing workspace

> I always forget this so made a note of it

```bash
$ mkdir hello

$ cd hello

$ go mod init org/user/hello
...

$ cat go.mod
...
```

To build the file and get an executable:

```bash
go install org/user/hello
```

The binary is stored in `$GOBIN` (if that is set), in `$GOPATH/bin`, or in `~/go/`

To set the GO variables:

```bash
$ go env -W GOBIN='' # To set a variable

$ go env -u GOBIN='' # To unset a variable
```

----
