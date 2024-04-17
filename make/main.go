package main

import "fmt"

type courseRatings map[string]float64

func (m *courseRatings) output() {
	fmt.Println(*m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Yay"

	userNames = append(userNames, "Andre")
	userNames = append(userNames, "Lopes")

	fmt.Println(userNames)

	// courseRatings := make(map[string]float64, 3)
	courseRatings := make(courseRatings, 3)

	courseRatings["Go"] = 4.7
	courseRatings["React"] = 4.8
	courseRatings["Angular"] = 4.6

	// fmt.Println(courseRatings)
	courseRatings.output()
}
