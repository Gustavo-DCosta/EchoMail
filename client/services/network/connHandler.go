package services

import (
	"fmt"
)

func ConnHandler(newUser bool) {
	phoneNumber, emailAddress := GetCredentials(newUser)
	if phoneNumber == "" || emailAddress == "" {
		return
	}

	uuid, err := SendConnCredentials(phoneNumber, emailAddress, newUser)
	if err != nil {
		fmt.Println("an error occured, please try again", err)
		Check(err)
		return
	}

	go RunSaveUUID(uuid)
	token := Getotp()
	accessToken, err := SendOtp(uuid, token)
	if err != nil {
		fmt.Println("an error occured, please try again", err)
		Check(err)
		return
	}
	go RunSaveJWT(accessToken)
	AppUnlocked(emailAddress)
}
