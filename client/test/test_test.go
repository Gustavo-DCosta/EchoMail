package test

import (
	"testing"
)

func TestGetcreds(t *testing.T) {
	var email string = "luna>>!**$ryx"
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
