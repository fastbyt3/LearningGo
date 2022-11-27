# Arrays in Go

- General syntax:
```
[capacity]data_type{element_values}
```
- indexing starts at 0
- array initialization

```go
var arr = new([5]int)
```

- all values are populated with Zero of respective data type

- entering values into arr

```go
var arr = [5]int{1,2,3,4,5}

// For arr with non-fixed length
var arr = [...]int{1,2,3,4,5}
var arr = []int{1,2,3,4,5}
```


### Creating 2D arrays

```go
a := [2][2]int{
	{1,1},
	{1,0},
}
```

### Array Slice

> Slice is a reference to a continuous fragment of an array (which we call the relevant array, usually anonymously), so the slice is a reference type (not the same as an array).


- Slice can be an entire arr or a subset 
- general rep : `arr[start:end]`
- len of slice _can be modified during runtime_:
	+ shortening size : slicing
	+ increasing size : using `append()` 

- In Go, slices are more versatile

- To create slices we can use the built in function : `make()`

```go
make([] T, len, cap)
```

- len is the size of the slice
- cap -> optional; denotes capacity

```go
make([]int, 10, 50) 
// This allocates an array with a 50 int value, 
// and creates a slice v with a length of 10 
// with a capacity of 50, which points to the first 10 elements of the array.
```

> more details : https://golangr.com/go-slice/


### Range

- iterates over elements in an array or dict or other DS
- syntax:

```go
for i, val := range arr {
	// ...
}
```

- if we dont have a use for index we can use `_`


### Notes

- In code below:
    ```golang
    func mod1(a []int) {
        for i := range a {
            a[i] = 5
        }

        fmt.Println("1:", a)
    }

    func mod2(a []int) {
        a = append(a, 125) // !

        for i := range a {
            a[i] = 5
        }

        fmt.Println("2:", a)
    }

    func main() {
        s1 := []int{1, 2, 3, 4}
        mod1(s1)
        fmt.Println("1:", s1)

        s2 := []int{1, 2, 3, 4}
        mod2(s2)
        fmt.Println("2:", s2)
    }
    ```
    - The output is `1 [5 5 5 5] 1 [5 5 5 5] 2 [5 5 5 5 5] 2 [1 2 3 4]`
    - This is because in mod1 we pass `a` as a Link so any changes made get reflected
    - The same case applies for mod2 but using `append(a, 125)` ends up increasing size of arr thus creating a new arr
    - so `a` in mod2 is an entirely new arr
