package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	var appUser User

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser = User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	outputUserDetails(appUser)
}

func outputUserDetails(u User) {
	fmt.Println("First Name: ", u.firstName)
	fmt.Println("Last Name: ", u.lastName)
	fmt.Println("Birthdate: ", u.birthdate)

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
