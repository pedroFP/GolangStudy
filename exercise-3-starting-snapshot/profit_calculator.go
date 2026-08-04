package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	writeResultsToFile(ebt, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (userInput float64, err error) {
	fmt.Print(infoText)
	_, err = fmt.Scan(&userInput)

	if err != nil {
		err = errors.New("Wrong input")
	} else if userInput <= 0.0 {
		err = errors.New("Invalid number, must be greater than zero (0)")
	}

	return userInput, err
}

func writeResultsToFile(ebt, profit, ratio float64) {
	resultsFile := "results.txt"

	result := fmt.Sprintf("ebt: %.2f\nprofit: %.2f\nratio: %.2f", ebt, profit, ratio)

	os.WriteFile(resultsFile, []byte(result), 0644)
}
