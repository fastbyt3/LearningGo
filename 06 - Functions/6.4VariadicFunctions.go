package main

import "fmt"

func main() {
	fmt.Println(sumOfNumbers(1, 2, 3, 4, 5))
	fmt.Println(sumOfNumbers(59, 15, 3, 74, 2, 48, 36, 4, 48, 1, 2))
}

/*
-------- Variadic Functions ----------

these func can take ANY number of args
i.e. args are not explicitly defined

eg : fmt.Print()

defining these func:
func name(<varName> ...<Type>) {}
--------------------------------------
*/

func sumOfNumbers(noList ...int) int {
	// fmt.Println(noList)
	total := 0
	for _, n := range noList {
		total += n
	}
	return total
}
