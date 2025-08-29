package test

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetcreds(t *testing.T) {
	var email string = "lunaryx"
	ilegalChars := []rune{
		' ', '!', '#', '$', '%', '^', '&', '*', '(', ')', '=', '+', ',', '/', '\\', ';', ':', '\'', '"', '<', '>', '?', '[', ']', '{', '}', '|', '`', '~',
	}

	for _, c := range email {
		for _, ilegal := range ilegalChars {
			if c == ilegal {
				t.Errorf("Found an ilegal character")
			}
		}
	}
	input := email[0]

	if len(email) <= 3 || len(email) > 15 {
		t.Errorf("The email should be between 3 -> 15 characters")
	}
	if input == '.' {
		t.Errorf("No email should start with a .")
	}

}

func TestGetPhoneNumber(t *testing.T) {
	phone := "+33762691203"

	// occurs one or more times with the + SHOULD ONLY OCCUR 0 or one so we should use the *
	pattern := regexp.MustCompile(`^\+?\d{4,17}$`)
	//doesn't lets me do \+\ maybe it's /+/?
	// North Korea phone number goes until 7, and Nieu has a max lenght of 4
	// Check if the phone number follows the pattern
	ok := pattern.MatchString(phone)
	fmt.Println("Regex expression:	", ok)

	// returns false
}
