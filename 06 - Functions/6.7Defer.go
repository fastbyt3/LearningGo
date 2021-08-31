package main

import "fmt"

/*

--------------------- Defer ---------------------

```
defer <function_name>
```
Using 'defer' will ensure that the mentioned
function/method will be executed at end of
current function

-------------------------------------------------
*/

func printHello() {
	fmt.Println("Hello, World!!")
}

func bye() {
	fmt.Println("--- Called after main has finished exec ---")
	fmt.Println("Goodbye ✌")
}

func main() {
	defer bye()
	printHello()
	fmt.Println("---- Inside main ----")
	fmt.Println("blah blah blah")
}
