package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	var appUser *user.User
	// appUser.firstName = getUserData("Please enter your first name: ")
	// appUser.lastName = getUserData("Please enter your last name: ")
	// appUser.birthDate = getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(firstName, lastName, birthDate)
	// ... do something awesome with that gathered data!

	if err != nil {
		panic(err)
	}

	admin := user.NewAdmin("example@exmaple.com", "1234")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
