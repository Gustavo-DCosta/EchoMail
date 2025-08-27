// This package will store shell widgets that are reused across multiple packages

package shared

import (
	"fmt"

	"os"

	"strings"

	inoutput "github.com/Gustavo-DCosta/EchoMail/client/services/io"
	"github.com/fatih/color"
	"golang.org/x/term"
)

func HelpCommand() {
	fmt.Print("Enter ")
	color.RGB(224, 166, 31).Print("command")
	fmt.Print(" (")
	color.RGB(96, 168, 74).Print("'help' ")
	color.RGB(92, 132, 232).Print(" for ")
	fmt.Print(" options): \n")
}

func CenterElement(Text string, param2 bool) {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		inoutput.Check(err)
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

func ClearUI() {
	fmt.Print("\033[2J\033[H")
}

func StdOutHelp() {
	StdOutColor := color.RGB(239, 207, 167)
	StdOutColor.Println(`

	register		-command to create a new account
	login 			-command to connect to an existent account
	clear			-command to clear the screen
	exit			-command to exit the app (why would you??)

	`)
}

func StdOutInterHelp() {
	StdOutColor := color.RGB(239, 207, 167)
	StdOutColor.Println(`
	
	send 			-command to send a message
	exit			-command to exit the app		
	clear			-command to clear the screen
	esc				-command to return to lockscreen

	`)
}
