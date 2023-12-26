package main

import (
	"fmt"
	"time"
)

//createdAt in this case is a nested struct
type user struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//struct literal notation
	var appUser user = user{
		firstName: userFirstName,
		lastName: userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	//alternative notation
	// var appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now()
	// }

	// ... do something awesome with that gathered data!

	fmt.Println(firstName, lastName, birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
