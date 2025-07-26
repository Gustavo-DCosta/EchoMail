package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Gustavo-DCosta/EchoPulse/client/model"
)

func SendRegisterCredentials(PhoneNumber string, EmailAddress string) (string, error) {
	url := os.Getenv("Server_Register_EndPoint")
	if url == "" {
		fmt.Println("There occurred a problem getting url")
		return "", errors.New("missing endpoint")
	}

	RequestCredentials := model.SignupRequest{
		Phone_Number:  PhoneNumber,
		Email_Address: EmailAddress,
	}

	jsonRequestCredentials, err := json.Marshal(RequestCredentials)
	if err != nil {
		fmt.Println("Error encoding into JSON:", err)
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestCredentials))
	if err != nil {
		fmt.Println("Error creating new HTTP request:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		return "", errors.New("bad status code")
	}

	//reads UUID (body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	type UUIDResponse struct {
		UUID string `json:"uuid"`
	}

	var response UUIDResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return "", err
	}

	fmt.Println("Received UUID from server:", response.UUID)
	return response.UUID, nil
}
