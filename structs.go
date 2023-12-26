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

//Added receiving arguments
func (u user) outputUserDetails () {

	//This works, because pointers to struct is an exception of dereferencing.
	//Go allows this.
	fmt.Println(u.firstName, u.lastName, u.birthDate)

	//Technically, it should be
	// fmt.Println((*u).firstName,(*u).lastName,(*u).birthDate)
}

func (u *user) clearUsername() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthDate string) user {
	return user{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//struct literal notation
	var appUser user = newUser(userFirstName, userLastName, userBirthdate)

	//alternative notation
	// var appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now()
	// }

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
	appUser.clearUsername()
	appUser.outputUserDetails()
	// fmt.Println(firstName, lastName, birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
