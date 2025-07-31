package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Gustavo-DCosta/EchoPulse/client/model"
)

func SendLoginCreds(PhoneNumber string) error {
	url := os.Getenv("Login_Endpoint")
	if url == "" {
		fmt.Println("error getting the url, .env was not loaded")
	}

	loginstruct := model.LoginPayload{
		StructPhoneNumber: PhoneNumber,
	}

	reqPayload, err := json.Marshal(loginstruct)
	if err != nil {
		fmt.Println("Couldn't marshal the struct: ", err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqPayload))
	if err != nil {
		fmt.Println("Coudln't create a request | ERROR CODE: ", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request | ERROR CODE: ", err)
		return err
	}

	defer resp.Body.Close()
	return nil
}

func SendloginOTP(PhoneNumber, Token string) error {
	url := os.Getenv("Verification_URL")
	if url == "" {
		fmt.Println("Couldn't load the env")
	}
	return nil
}
