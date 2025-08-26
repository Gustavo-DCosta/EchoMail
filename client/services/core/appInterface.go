package services

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func HelpCommand() {
	fmt.Print("Enter ")
	color.RGB(224, 166, 31).Print("command")
	fmt.Print(" (")
	color.RGB(96, 168, 74).Print("'help' ")
	color.RGB(92, 132, 232).Print(" for ")
	fmt.Print(" options): \n")
}

func AppUnlocked(EmailAddress string) {
	CenterElement("Welcome "+EmailAddress+"\n", true)
	HelpCommand()
	IOParser()
}

func IOParser() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("=> !")
		scanner.Scan()
		err := scanner.Err()

		if err != nil {
			Check(err)
			fmt.Println("Sorry something went wrong") // Letting the user know that it failed
			// Can't read input so should leave
			return // swapped os.Exit(1) by returning to the last function
		}

		cmdInput := scanner.Text()

		switch cmdInput {
		case "send":
			MsgWS()
		case "help":
			StdOutInterHelp()
		case "exit":
			os.Exit(0)
			// fixed exit status
		case "clear":
			ClearUI()
		case "esc":
			fmt.Println("Going back to lockscreen...")
			CenterElement("[EchoMail]", false)
			LockScreenUX()
		default:
			fmt.Println("😺 Ayo captain Whiskers McGraw doesn't know that command!!")
		}

	}
}
