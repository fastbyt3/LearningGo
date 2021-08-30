# Struct data in Go

- Declaration

```go
type structName struct {
	var1 type
	var2 type
}
```

- creating an instance of declared struct

```go
var instance1 structName

var instance2 = structName{
	var1: "bar",
	var2: "foo",
}

var instance3 = new(structName)

var instance4 = &structName{...}
```

- instance4 is an instance created using pointers
