# Golang Testing

## Unit testing

### Folder structure

```
math
|
|_ math.go
|_ math_test.go
```

> Test GO files should always start with name of file to test appended with `_test.go`

### Testing methods

- test methods must use Pascal naming scheme
- method name starts with `Test` followed by name of the method to be tested

For testing `Foo(...)` we use the test method: `TestFoo(t *testing.T)`

For testing data types(say `T`) being used in a function `(T) Foo(...)` we have a test method: `TestT_Foo`

```go
struct T {....}

func (T) Foo(...) {}
```

### Usage of `testing`

- import `testing` pkg
- create `testF` where F is the func name
- create test cases

#### Types of messages in `testing`

1. `t.Log()` -> similar to `fmt.Println()`
2. `t.Fail()` -> show a test case has failed in results
3. `t.FailNow()` -> `t.Fail()` + exit immediately
4. `t.Error()` -> unexpected state : `t.Log() + t.Fail()`
5. `t.Fatal()` -> `t.Log() + t.FailNow()` 

#### Basic Unit tests

```bash
[fastbyte@omen math]$ go test -v
=== RUN   TestAbs
    math_test.go:7: Abs(-1) =  1
--- PASS: TestAbs (0.00s)
PASS
ok  	learn/testing/math	0.002s

[fastbyte@omen math]$ go test -v
=== RUN   TestAbs
    math_test.go:7: Abs(-1) =  -1
    math_test.go:9: Expected 1
--- FAIL: TestAbs (0.00s)
FAIL
exit status 1
FAIL	learn/testing/math	0.002s
```

#### Subset of test

- To create a subset f of Test t, use `t.Run(name, f func(t *T))`

```golang
t.Run("PositiveNumber", func(t *testing.T) {
    if Abs(100) != 100 {
        t.Error("Unexpected")
    }
})
```

Now all subtests get executed under main func test name:

```
[fastbyte@omen math]$ go test -v
=== RUN   TestAbs
=== RUN   TestAbs/Positive
=== RUN   TestAbs/Zero
=== RUN   TestAbs/Negative
--- PASS: TestAbs (0.00s)
    --- PASS: TestAbs/Positive (0.00s)
    --- PASS: TestAbs/Zero (0.00s)
    --- PASS: TestAbs/Negative (0.00s)
PASS
ok  	learn/testing/math	0.007s
```

#### Cleanup function

- Kinda like a deferred method but this is **concurrently safe**
- Helpful when testing involves opening one or a shit ton of files

#### Parallel execution of tests

- Use `t.Parallel()` to execute test functions parallely

