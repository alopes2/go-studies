package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	hobbies := [3]string{"Sports", "Cooking", "Gaming"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// mainHobbies := hobbies[0:2]
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:cap(mainHobbies)]
	fmt.Println(mainHobbies)

	courseGoals := []string{"Learn Go!", "Learn all the basics!"}
	fmt.Println(courseGoals)

	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn the basics!")
	fmt.Println(courseGoals)

	products := []Product{
		{
			id:    "first-product",
			title: "A First Product",
			price: 12.99,
		},
		{
			id:    "second-product",
			title: "A Second Product",
			price: 129.99,
		}}

	fmt.Println(products)

	newProduct := Product{
		id:    "third-product",
		title: "A Third Product",
		price: 15.99,
	}

	products = append(products, newProduct)

	fmt.Println(products)
}
