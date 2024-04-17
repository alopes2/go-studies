package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		panic(err)
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		panic(err)
	}

	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		panic(err)
	}

	earningsBeforeTax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	earningsOutput := fmt.Sprintf("Revenue before taxes: %.2f\n", earningsBeforeTax)

	profitOutput := fmt.Sprintf("Profit: %.2f\n", profit)

	fmt.Print(earningsOutput, profitOutput)

	fmt.Printf("Ratio: %.2f\n", ratio)

	storeResults(earningsBeforeTax, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("Earnings Before Tax: %.2f\nProfit: %.2f\nRatio: %.3f\n", ebt, profit, ratio)

	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	return earningsBeforeTax, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number")
	}

	return userInput, nil
}
