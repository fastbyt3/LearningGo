# Gin web framework

## Base setup

```go
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
```

- `gin.Default()` -> Creates a router

- we have `r.GET(...)` and `r.POST(...)`
- by default the server is started on port 8080 by `r.run()`


