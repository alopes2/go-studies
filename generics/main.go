package main

import "fmt"

func main() {
	// result := add[int](1, 2)
	result := add(1, 2)

	fmt.Println(result)

	// result2 := add[float64](1.5, 2.3)
	result2 := add(1.5, 2.3)

	fmt.Println(result2)

	// result3 := add[string]("Hello ", "World!")
	result3 := add("Hello ", "World!")

	fmt.Println(result3)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

// func add[T any](a, b T) T {
// 	return a + b
// }
