package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Gustavo-DCosta/EchoPulse/client/model"
)

type OTPResponse struct {
	JWT     string `json:"jwt"`
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

func SendOTPcode(UUID string, token string) error {
	url := os.Getenv("Verification_OTP_endpoint")

	OTPdataPayload := model.VerifyOTPpayload{
		UUID:  UUID,
		Token: token,
	}

	jsonPayload, err := json.Marshal(OTPdataPayload)
	if err != nil {
		fmt.Println("Error marshaling the data", err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error creating new request", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d, body: %s\n", resp.StatusCode, string(body))
		return fmt.Errorf("server returned status %d", resp.StatusCode)
	}

	// Parse the response
	var otpResponse OTPResponse
	err = json.Unmarshal(body, &otpResponse)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return err
	}

	if !otpResponse.Success {
		return fmt.Errorf("OTP verification failed: %s", otpResponse.Error)
	}

	// Create jwt directory if it doesn't exist
	jwtDir := "./jwt"
	err = os.MkdirAll(jwtDir, 0755)
	if err != nil {
		fmt.Println("Error creating jwt directory:", err)
		return err
	}

	// Save JWT to file
	jwtPath := filepath.Join(jwtDir, "token.jwt")
	err = os.WriteFile(jwtPath, []byte(otpResponse.JWT), 0644)
	if err != nil {
		fmt.Println("Error saving JWT to file:", err)
		return err
	}

	fmt.Printf("JWT successfully saved to %s\n", jwtPath)
	return nil
}
