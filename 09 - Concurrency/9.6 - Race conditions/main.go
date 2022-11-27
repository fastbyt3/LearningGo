package main

import (
	"fmt"
	"sync"
	"time"
)

type Inventory struct {
	stock int
}

func (i *Inventory) getStock() int {
	return i.stock
}

func (i *Inventory) updateStock(amount int) {
	i.stock = amount
}

func buyPhone(i *Inventory, who string) {
	time.Sleep(100 * time.Millisecond)
	stock := i.getStock()

	time.Sleep(100 * time.Millisecond)
	if stock <= 0 {
		fmt.Printf("Sorry, %s we are out of stock\n", who)
		return
	}
	fmt.Printf("Hi %s, We have %d currently in stock\n", who, stock)

	time.Sleep(800 * time.Millisecond)
	fmt.Printf("Sold 1 phone to %s\n", who)
	i.updateStock(stock - 1)
}

func main() {
	var inventory Inventory
	inventory.updateStock(3)

	customers := []string{"Alice", "Bob", "Charles", "Dan", "Eugene", "Florentia", "Grundy", "Holton", "Iceman", "Jacob"}
	var wg sync.WaitGroup
	wg.Add(len(customers))

	for _, customer := range customers {
		go func(who string) {
			buyPhone(&inventory, who)
			wg.Done()
		}(customer)
		time.Sleep(500 * time.Millisecond)
	}

	wg.Wait()
	fmt.Printf("\nFinal inventory: Stock = %d\n", inventory.getStock())
}
