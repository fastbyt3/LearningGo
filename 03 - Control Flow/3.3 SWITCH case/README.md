# Switch case : Go

basic syntax:

```go
switch var {
	case val1:
		...
	case val2: 
		...
	default:
		...
}
```

> Values as well as conditions(like x > 0) can be used in the cases

- `break` statements is not required as it exits the switch block if cond is satisfied
- In case we want to exec the subsequent cases despite failing the cond we can use `fallthrough`

---