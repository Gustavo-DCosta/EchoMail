package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func GetCredentialsRegister() (string, string) {
	StdOutCred := color.RGB(204, 190, 214)

	StdOutCred.Print("Display name (for [name<>echomail.dev]): ")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error sanitizing input", err)
	}

	name := strings.TrimSpace(input)

	EmailAdress := name + "<>echomail.dev"

	StdOutCred.Print("Phone number (include country code): ")

	reader = bufio.NewReader(os.Stdin)

	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error sanitizing input", err)
	}

	phoneNumber := strings.TrimSpace(input)

	return phoneNumber, EmailAdress
}
