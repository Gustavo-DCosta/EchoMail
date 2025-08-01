package services

import (
	"fmt"
)

func ConnHandler(accountStatus bool) {
	phoneNumber, emailAddress := GetCredentials()

	uuid, err := SendConnCredentials(phoneNumber, emailAddress, accountStatus)
	if err != nil {
		fmt.Println("an error occured, please try again")
	}

	go RunSaveUUID(uuid)
	token := Getotp()
	accessToken, err := SendOtp(uuid, token)
	if err != nil {
		fmt.Println("an error occured, please try again")
	}
	go RunSaveJWT(accessToken)
	AppUnlocked(emailAddress)
	return
}
