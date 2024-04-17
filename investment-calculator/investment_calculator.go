package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

	outputText("Input the investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Input the number of years: ")
	fmt.Scan(&years)

	outputText("Input the return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// fmt.Println(futureValue)
	// fmt.Println(futureRealValue)

	fmt.Printf(`Future Value: %.1f 
	
	Adjusted with inflation: %.1f`, futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, investimentYears float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, float64(investimentYears))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, investimentYears)
	return futureValue, futureRealValue
}

// func calculateFutureValues(investmentAmount, expectedReturnRate, investimentYears float64) (futureValue float64, futureRealValue float64) {
// 	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, float64(investimentYears))
// 	futureRealValue = futureValue / math.Pow(1+inflationRate/100, investimentYears)
// 	return
// }
