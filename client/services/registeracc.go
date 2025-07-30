package services

import "fmt"

func RegisterAccount() {
	PhoneNumber, EmailAddress := GetCredentialsRegister()

	uuid, err := SendRegisterCredentials(PhoneNumber, EmailAddress)
	if err != nil {
		fmt.Println("Couldn't send credentials to the server")
		return
	}

	Token := GetSMScode()

	err = SendOTPcode(uuid, Token)
	if err != nil {
		fmt.Println("Couldn't create your account: ", err)
		return
	} else {
		AppUnlocked(EmailAddress)
	}
}
