package services

// config files

/*
-> {
	after jwt received make a account numbers number variable
	store n as number of accounts

	account n
		email
		phone number

	if n > 1 on free subscription then delete 1 account, only letting accountfor free subscription
-> }  // TODO: later
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Gustavo-DCosta/EchoMail/client/cache"
	"github.com/fatih/color"
)

func GetCredentials(newUser bool) (string, string) {
	stdErr := color.RGB(255, 0, 0)
	stdOutCred := color.RGB(204, 190, 214)
	reader := bufio.NewReader(os.Stdin)

	var emailAddress string
	if newUser {
		for {
			stdOutCred.Print("Display name (for [name<>echomail.dev]): ")
			input, _ := reader.ReadString('\n')
			name := strings.TrimSpace(input)
			if name != "" {
				emailAddress = name + "<>echomail.dev"
				go saveEmaillAdr(emailAddress)
				// call func in gorountines to save rhe email adress in a json object
				// inside config folder
				// cofig/profile.json
				// email_address: example<>echomail.dev
				cache.Set("UserEmail", emailAddress)
				break
			}
			stdErr.Println("Please insert a name.")
		}
	} else {
		val, ok := cache.Get("UserEmail")
		if !ok {
			stdErr.Println("No cached email found. Please create a new user.")
			return GetCredentials(true) // fallback to new user flow
		}
		emailAddress = val
	}

	var phoneNumber string
	for {
		stdOutCred.Print("Phone number (include country code): ")
		input, _ := reader.ReadString('\n')
		phoneNumber = strings.TrimSpace(input)
		if phoneNumber != "" {
			break
		}
		stdErr.Println("Please insert a phone number.")
	}

	return phoneNumber, emailAddress
}

func Getotp() string {
	//Added protection agaisn't empty code upload
	const max = 6
	var token string
	stdOutToken := color.RGB(224, 180, 130)

	reader := bufio.NewReader(os.Stdin)

	// Ask once before loop
	stdOutToken.Print("Please insert the code received on your phone: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		Check(err)
		fmt.Println("error sanitizing input", err)
		return ""
	}
	input = strings.TrimSpace(input)

	if len(input) != max {
		fmt.Println("Incorrect code, please insert the real code!")
		for {
			stdOutToken.Print("Please insert the code received on your phone: ")
			input, err = reader.ReadString('\n')
			if err != nil {
				Check(err)
				fmt.Println("error sanitizing input", err)
				return "" // return if it isn't possible to read the input

			}
			input = strings.TrimSpace(input)
			if len(input) == max {
				token = input
				break
			}
		}
	} else {
		token = input
	}

	return token
}
