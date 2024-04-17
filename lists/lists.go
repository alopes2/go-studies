package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}

	fmt.Println(prices[1])
	fmt.Println(prices[:1])

	prices[1] = 9.99

	newPrices := append(prices, 5.99)

	fmt.Println(newPrices)

	newPrices = append(prices, 5.99, 12.99, 100.10)

	discoutPrices := []float64{101.99, 80.99, 20.59}

	prices = append(prices, discoutPrices...)

	fmt.Println(prices)
}

// type Product struct {
// 	title string
// 	id    string
// 	price float64
// }

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	productNames[2] = "A Car"

// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	featuredPrices := prices[1:3]

// 	fmt.Println(featuredPrices)

// 	firstTwoPrices := prices[:2]

// 	fmt.Println(firstTwoPrices)

// 	lastTwoPrices := prices[2:]

// 	fmt.Println(lastTwoPrices)

// 	lastTwoPrices[1] = 9999.99
// 	fmt.Println(prices)
// }
