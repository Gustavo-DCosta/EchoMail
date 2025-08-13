package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/term"
)

func CenterElement(Text string, param2 bool) {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80
	}

	txtLen := len(Text)

	if txtLen >= width {
		fmt.Println(Text)
		return
	}

	padding := (width - txtLen) / 2
	if param2 == true {
		color.RGB(122, 77, 89).Print(strings.Repeat(" ", padding) + Text)
	} else {
		fmt.Println(strings.Repeat(" ", padding) + Text)
	}
}

func LockScreen() {
	fmt.Print("Enter ")
	color.RGB(224, 166, 31).Print("command")
	fmt.Print(" (")
	color.RGB(96, 168, 74).Print("'help' ")
	color.RGB(92, 132, 232).Print(" for ")
	fmt.Print(" options): \n")
}

func LockScreenUX() { // I didn't knew what to call this function if someone has a propusition for later I am down
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
			ConnHandler(false)
		case "register":
			ConnHandler(true)
		case "exit":
			fmt.Println("command to exit")
			os.Exit(0)
		case "help":
			StdOutHelp()
		case "clear":
			ClearUI()
		case "completion":
			fmt.Println("Actually this command doesn't work, because he isn't supposed to be here, sorry not sorry muah")
		case "guest":
			AppUnlocked("guest")
		default:
			fmt.Println("😺 Ayo captain Whiskers McGraw doesn't know that command!!")
		}
	}
}
