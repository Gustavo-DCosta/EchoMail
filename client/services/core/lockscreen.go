package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	//"github.com/Gustavo-DCosta/EchoMail/client/services/core"
	"github.com/Gustavo-DCosta/EchoMail/client/services/shared"
)

func LockScreenPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("=> !")

		if !scanner.Scan() {
			fmt.Println("Oh Uh something went wrong")
			break
		}
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))

		switch input {
		case "login":
			shared.ConnHandler(false)
		case "register":
			shared.ConnHandler(true)
		case "exit":
			fmt.Println("command to exit")
			os.Exit(0)
		case "help":
			shared.StdOutHelp()
		case "clear":
			shared.ClearUI()
		case "completion":
			fmt.Println("Actually this command doesn't work, because he isn't supposed to be here, sorry not sorry muah")
		case "guest":
			//core.AppUnlocked("guest")
		default:
			fmt.Println("😺 Ayo captain Whiskers McGraw doesn't know that command!!")
		}
	}
}
