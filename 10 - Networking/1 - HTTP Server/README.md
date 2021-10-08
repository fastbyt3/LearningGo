# Web Server in Golang

- Main module: `net/http`.
- To start the web server:

```golang
http.ListenAndServe(":9999", nil)
```

- To create routes / pages:

```golang
http.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
	// ....
	// ....
})

// --------------------------
// |		or 				|
// --------------------------

http.HandleFunc("/newRoute", foo())

//outside main()

func foo() http.Handler {
	return htt.HandleFunc(func(w http.ResponseWriter, r http.Request) {
		// ...
	})
}
```

### Server MUX

> Request multiplexer : compares incoming requests against a list of predefined URL paths, and calls the associated handler for the path whenever a match is found.


