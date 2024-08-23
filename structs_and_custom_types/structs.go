package main

import (
	"fmt"
	"time"
	"errors"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u *User) outputUserDetails() {
	fmt.Println("First Name: ", u.firstName)
	fmt.Println("Last Name: ", u.lastName)
	fmt.Println("Birthdate: ", u.birthdate)
	fmt.Println("Created At: ", u.createdAt)

}

func (u *User) clearUserNames() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First Name, Last Name and Birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	var appUserPtr *User

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUserPtr, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}


	appUserPtr.outputUserDetails()
	appUserPtr.clearUserNames()
	appUserPtr.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
