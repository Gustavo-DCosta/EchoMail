package core

import (
	"bufio"
	"fmt"
	"os"

	inoutput "github.com/Gustavo-DCosta/EchoMail/client/services/io"
	"github.com/Gustavo-DCosta/EchoMail/client/services/shared"
)

func AppUnlocked(EmailAddress string) {
	shared.CenterElement("Welcome "+EmailAddress+"\n", true)
	shared.HelpCommand()
	IOParser()
}

func IOParser() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("=> !")
		scanner.Scan()
		err := scanner.Err()

		if err != nil {
			inoutput.Check(err)
			fmt.Println("Sorry something went wrong") // Letting the user know that it failed
			// Can't read input so should leave
			return // swapped os.Exit(1) by returning to the last function
		}

		cmdInput := scanner.Text()

		switch cmdInput {
		case "send":
			//MsgWS()
			fmt.Println("yo")
		case "help":
			shared.StdOutInterHelp()
		case "exit":
			os.Exit(0)
			// fixed exit status
		case "clear":
			shared.ClearUI()
		case "esc":
			fmt.Println("Going back to lockscreen...")
			shared.CenterElement("[EchoMail]", false)
			LockScreenPrompt()
		default:
			fmt.Println("😺 Ayo captain Whiskers McGraw doesn't know that command!!")
		}

	}
}
