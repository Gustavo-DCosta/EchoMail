package services

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

type CreateAccCredentials struct {
	EmailDomain string `json:"email_domain"`
	PhoneNumber string `json:"phone_number"`
}

func RegisterLogic() {
	RegisterCredentials()
}

func RegisterCredentials() {
	url := "http://127.0.0.1:8080/api/register"
	color.RGB(204, 190, 214).Print("Please insert a display name (it will be added to your email, example [displayName<>echomail.dev]: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred getting the user input, please retry...")
		return
	}

	// Remove the newline character from the input
	displayName := strings.TrimSpace(input)

	// Create the full email address
	fullEmail := displayName + "<>echomail.dev"

	color.RGB(204, 190, 214).Print("Please insert your phone number (think about including the country digits too please): ")
	reader = bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred getting the user input, please retry...")
		return
	}

	phoneNumber := strings.TrimSpace(input)

	// Giving values to the structure variables
	createAccCredentials := CreateAccCredentials{
		EmailDomain: fullEmail,
		PhoneNumber: phoneNumber,
	}

	// Marshaling the values
	jsonCredentials, err := json.Marshal(createAccCredentials)
	if err != nil {
		fmt.Println("Error marshaling credentials:", err)
		return // Handle the error appropriately
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonCredentials))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return // Handle the error appropriately
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return // Handle the error appropriately
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		// Handle the error appropriately
	}
}
