package services

import (
	"github.com/fatih/color"
)

func StdOutHelp() {
	StdOutColor := color.RGB(239, 207, 167)
	StdOutColor.Println(`
	register		-command to create a new account
	login 			-command to connect to an existent account
	clear			-command to clear the screen
	exit			-command to exit the app (why would you??)`)
}
