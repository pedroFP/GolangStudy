package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthDate (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, error := user.New(userFirstName, userLastName, userBirthDate)

	admin := user.NewAdmin("test@example.com", "test123")
	admin.User.OutputUserDetail()
	admin.User.ClearUserName()
	admin.User.OutputUserDetail()

	if error != nil {
		fmt.Println(error)
		return
	}

	appUser.OutputUserDetail()
	appUser.ClearUserName()
	appUser.OutputUserDetail()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
