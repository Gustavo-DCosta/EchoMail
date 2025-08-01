package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func GetCredentials() (string, string) {
	stdOutCred := color.RGB(204, 190, 214)

	stdOutCred.Print("Display name (for [name<>echomail.dev]): ")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error sanitizing input", err)
	}

	name := strings.TrimSpace(input)

	EmailAdress := name + "<>echomail.dev"

	stdOutCred.Print("Phone number (include country code): ")

	reader = bufio.NewReader(os.Stdin)

	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error sanitizing input", err)
	}

	phoneNumber := strings.TrimSpace(input)

	return phoneNumber, EmailAdress
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
		fmt.Println("error sanitizing input", err)
	}
	input = strings.TrimSpace(input)

	if len(input) != max {
		fmt.Println("Incorrect code, please insert the real code!")
		for {
			stdOutToken.Print("Please insert the code received on your phone: ")
			input, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("error sanitizing input", err)
				continue
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
