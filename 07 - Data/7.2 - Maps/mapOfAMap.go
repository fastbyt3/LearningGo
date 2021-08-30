package main

import "fmt"

func main() {
	processors := map[string]map[string]int{
		"Intel": map[string]int{
			"i3": 30000,
			"i5": 50000,
			"i7": 60000,
			"i9": 80000,
		},
		"AMD": map[string]int{
			"3000": 25000,
			"5000": 50000,
		},
	}

	// fmt.Println(processors)
	fmt.Println(processors["Intel"]["i5"])
}
