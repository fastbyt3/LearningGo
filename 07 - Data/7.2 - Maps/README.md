# Maps in Go

> Maps are unordered collections of key-value pairs
> For every key there exists a **Unique** value

- can also be referred to as *Associativity array* or *hash table*
- declaring/initializing a map

```go
map1 := make(map[string]string)

var map2 = map[string]int{}

var map3 = map[string]int{"A": 1, "B": 2, "C": 3}
```

- adding new elements to the map can be done as:

```go
// adding elements
map3["D": 4]

// Updating items
map3["D": 99]
```

- deleting items from maps can be done too:

```go
delete(map3, "C") // deletes "C" key
```

- can be iterated over similar to arrays using `range`

