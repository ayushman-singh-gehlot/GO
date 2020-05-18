package main

import "fmt"

func main() {
	price, availableFunds := 100, 120
	var tax, total float64
	taxRate := 0.08

	fmt.Println("price is", price, "dollars.")
	tax = float64(price) * taxRate
	fmt.Println("tax is", tax, "dollars.")
	total = tax + float64(price)
	fmt.Println("total cost is", total, "dollars.")
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("within budget?", float64(availableFunds) >= total)
}
