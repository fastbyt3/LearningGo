# Logging in Golang

### `Log` package

- `Log` is a built in pkg that suffices basic logging needs
- doesn't come with logging levels

```go
// basic example
log.Println("First log!!")
```

#### Logging to a file

- create/open a new file 
- set `log` package to write to the opened file

```golang
log.SetOutput(file)
```

#### Custom loggers

- to create a new logger:

```go
log.New(file, prefix, flag)
```

- prefix: string to represent the level like _INFO_, _WARNING_, _ERROR_
- flag	: 
	+ to add different log flags separated by '|' 
	+ provides context to the log
	+ example:
		*  _log.Ldate_	-> (year/month/date)
		*  _log.Ltime_
		*  _log.Lshortfile_	-> (displays which line in the file log was registered)

### Logging frameworks

> Logging frameworks help standardizing log data and to create more structured log data

**glog** and **logrus** are the two most popular logging frameworks in Go.


---