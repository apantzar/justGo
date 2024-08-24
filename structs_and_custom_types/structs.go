package main

import (
	"fmt"
	"example.com/myapp/user"
)

func main() {
	var appUserPtr *user.User

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUserPtr, err :=  user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}


	admin := user.NewAdmin("test@mail.com", "password")
	admin.OutputUserDetails()


	appUserPtr.OutputUserDetails()
	appUserPtr.ClearUserNames()
	appUserPtr.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
