package services

import (
	"fmt"
)

func saveEmaillAdr(emailAddress string) error {
	fmt.Println("sucess, saved: ", emailAddress)

	return nil
}

func RunSaveEmail(str string) {
	if err := saveEmaillAdr(str); err != nil {
		Check(err)
		fmt.Println("Error saving the email adress")
	}
}
