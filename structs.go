package main

import (
	"fmt"

	"github.com/jxnhiro/go-practice/user"
)

func structs() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	var err error

	appUser, err = user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@test.com", "test123")

	admin.OutputUserDetails()
	admin.ClearUsername()
	admin.OutputUserDetails()
	//alternative notation
	// var appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now()
	// }

	// ... do something awesome with that gathered data!

	appUser.OutputUserDetails()
	appUser.ClearUsername()
	appUser.OutputUserDetails()
	// fmt.Println(firstName, lastName, birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
