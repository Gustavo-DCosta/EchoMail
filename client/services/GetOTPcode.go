package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func GetSMScode() string {
	StdOutOTPcode := color.RGB(149, 177, 219)
	StdOutOTPcode.Print("Please input the code received on your phone: ")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error sanitizing the input", err)
	}

	if input == "" {
		StdOutOTPcode.Print("Code was not saved, please rewrite it: ")
	}

	OTPcode := strings.TrimSpace(input)

	return OTPcode
}
