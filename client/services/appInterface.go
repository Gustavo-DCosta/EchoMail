package services

import "github.com/fatih/color"

func AppUnlocked(EmailAddress string) {
	CenterElement("Welcome"+EmailAddress, true)

	stdOutAppUnlocked := color.RGB(221, 211, 115)

	for {
		stdOutAppUnlocked.Print("=> !")
		IOParser()
	}
}
