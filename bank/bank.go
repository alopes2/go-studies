package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"

	"github.com/alopes2/go-studies/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------")
		// panic("Can't continue, sorry!")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println(randomdata.PhoneNumber())

	var choice int
	for choice != 4 {
		presentOptions()

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is %.2f\n", accountBalance)
			break
		case 2:
			var depositAmount float64
			fmt.Print("Deposit amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Printf("Your new balance is %.2f\n", accountBalance)

			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			break
		case 3:
			var withdrawAmount float64
			fmt.Print("Withdraw amount: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Printf("Your new balance is %.2f\n", accountBalance)

			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			break
		default:
			fmt.Println("Goodbye!")
		}
	}

	fmt.Println("Thanks for using our bank.")
}
