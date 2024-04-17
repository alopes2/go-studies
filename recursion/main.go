package main

import "fmt"

func main() {
	fact := factorial(5)

	fmt.Println(fact)
}

func factorial(number int) int {
	if number <= 1 {
		return number
	}

	return number * factorial(number-1)
}

// func factorial(number int) int {
// 	result := 1

// 	for i := 1; i < number; i++ {
// 		result = result * i
// 	}

// 	return result
// }
