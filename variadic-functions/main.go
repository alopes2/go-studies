package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 2, 3, 4)

	anotherSum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, val := range numbers {
		sum += val
	}

	return sum
}
