package main

import (
	"fmt"

	"example.com/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	presentOptions()

	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----")
	}

	fmt.Println("Welcome to Go Bank!")

	for {

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:

			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("How much you want to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount you can't withdraw more than you have")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			return
		}

	}
}
