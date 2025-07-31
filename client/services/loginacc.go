package services

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Login() {
	stdOutLogin := color.RGB(81, 133, 124)
	stdOutLogin.Print("Phone number (include country code): ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error sanitizing input | ERROR CODE: ", err)
		return
	}

	phoneNumber := input
	SendLoginCreds(phoneNumber)

	stdOutLogin.Print("Thank you please insert the received code on your phone: ")

	reader = bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error sanitizing input | ERROR CODE: ", err)
		return
	}

	loginOTP := input
	SendloginOTP(phoneNumber, loginOTP)
}
