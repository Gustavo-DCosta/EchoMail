package services

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func AppUnlocked(EmailAddress string) {
	CenterElement("Welcome"+EmailAddress, true)
}

func IOParser() {
	stdOutAppUnlocked := color.RGB(221, 211, 115)

	stdOutAppUnlocked.Print("=> !")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		err := scanner.Err()

		if err != nil {
			fmt.Println("Error starting parser: ", err)
		}

		input := scanner.Text()

		fmt.Println(input)
	}
}
