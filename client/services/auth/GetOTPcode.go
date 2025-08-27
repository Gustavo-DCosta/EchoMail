package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	inoutput "github.com/Gustavo-DCosta/EchoMail/client/services/io"
	"github.com/fatih/color"
)

func GetSMScode() string {
	stdOutOTPcode := color.RGB(149, 177, 219)
	stdOutOTPcode.Print("Please input the code received on your phone: ")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		inoutput.Check(err)
		fmt.Println("Error sanitizing the input", err)
		return ""
	}

	if input == "" {
		stdOutOTPcode.Print("The code seems empty Whiskers McGraw doesn't aprove it\n")
		stdOutOTPcode.Print("Please wirte it again:	")
		inoutput.InfoLogs("There was a problem to receive the OTP from user input")
	}

	OTPcode := strings.TrimSpace(input)

	return OTPcode
}
