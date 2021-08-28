# Variables in Golang

### Golang Strings

```go
var s1 = "s1 is a String variable"
```

- In go, string indexing starts at _0_

> We can use _printf_ by calling `fmt.Printf()`
> Here we can use **Format specifiers** like `%d, %c`

- String concat is also possible

```go
s := "hi " + "there!!"
s += " this is fastbyte."
```

> `:=` is a short variable declaration with no specific type
> When using `:=` only a NEW variable can be used

- Concat is also possible by using the **Strings** package

```go
s := Strings.Join([]string{"hello", "world"}, ", ")
```


### Other variable types(int, float, boolean, pointers)

- multiple variable declaration(of diff types) is possible

```go
var (
	i int
	b bool
	s string
	f float32
	p pointer
)
```

> float can be defined only as `float32` or `float64`
> i think `float64` is the default type

- multiple values can be assigned at same time

```go
i, b, s := 1, True, "xyz"
```

> At variable initialization all variables are given "zero" value of the variable type


### `nil` value

> The nil designator is used to represent a “zero value” of interface, function, maps, slices, and channels.If you do not specify the type of variable, the compiler will not be able to compile your code because it cannot guess a specific type.


---