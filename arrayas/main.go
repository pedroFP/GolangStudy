package main

import "fmt"

func main() {
	// var productNames [4]string

	prices := [4]float64{10.99, 20.0, 45.99, 20.0}

	// fmt.Println(prices)
	// fmt.Println(productNames)

	featuredPrices := prices[1:]
	featuredPrices[0] = 199.99
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3] // INFO: Originally highlightedPrices was [199.99] but now it will be [199.99 45.99 20]

	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}
