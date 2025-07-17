package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/term"
)

func CenterAppName(AppName string) {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80
	}

	txtLen := len(AppName)

	if txtLen >= width {
		fmt.Println(AppName)
		return
	}

	padding := (width - txtLen) / 2
	fmt.Println(strings.Repeat(" ", padding) + AppName)
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
			LoginLogic()
			// logic here

		case "register":
			fmt.Println("command to register")
			//logic here

		case "exit":
			fmt.Println("command to exit")
			os.Exit(0)
		default:
			fmt.Println("😺 Ayo captain Whiskers McGraw doesn't know that command!!")
		}
	}
}
