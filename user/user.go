package user

import (
	"errors"
	"fmt"
	"time"
)

//createdAt in this case is a nested struct
type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

//Use anonymous fields to use methods of the embedded struct
type Admin struct {
	email string
	password string
	User
}
//Added receiving arguments
func (u *User) OutputUserDetails () {

	//This works, because pointers to struct is an exception of dereferencing.
	//Go allows this.
	fmt.Println(u.firstName, u.lastName, u.birthDate)

	//Technically, it should be
	// fmt.Println((*u).firstName,(*u).lastName,(*u).birthDate)
}

func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("user input is invalid")
	}

	return &User{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}