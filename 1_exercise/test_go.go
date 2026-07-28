package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt, profit, ratio = getCalculations(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %.2f\n", ebt)

	fmt.Printf("Earnings after tax: %.2f\n", profit)

	fmt.Printf("Ratio: %.2f\n", ratio)
}

func getCalculations(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}
