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

func SendOTPcode(UUID string, token string) error {
	url := os.Getenv("Verification_OTP_endpoint")

	// Debug: Check URL
	if url == "" {
		return fmt.Errorf("Verification_OTP_endpoint not set")
	}
	fmt.Printf("Using URL: %s\n", url)

	OTPdataPayload := model.VerifyOTPpayload{
		UUID:  UUID,
		Token: token,
	}

	jsonPayload, err := json.Marshal(OTPdataPayload)
	if err != nil {
		fmt.Println("Error marshaling the data", err)
		return err
	}

	// Debug: Show what we're sending
	fmt.Printf("Sending: %s\n", string(jsonPayload))

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

	// Check for empty body BEFORE status check
	if len(body) == 0 {
		fmt.Println("WARNING: Response body is empty!")
		return fmt.Errorf("server returned empty response")
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d, body: %s\n", resp.StatusCode, string(body))
		return fmt.Errorf("server returned status %d: %s", resp.StatusCode, string(body))
	}

	// Try to parse JSON
	var otpResponse model.VerifySupabaseResponse
	err = json.Unmarshal(body, &otpResponse)
	if err != nil {
		fmt.Printf("JSON Unmarshal Error: %v\n", err)
		fmt.Printf("Trying to parse: '%s'\n", string(body))

		// Let's try to see if it's valid JSON at all
		var raw interface{}
		if json.Unmarshal(body, &raw) != nil {
			fmt.Println("Response is not valid JSON at all!")
		} else {
			fmt.Printf("Response is valid JSON but doesn't match our struct: %+v\n", raw)
		}

		return fmt.Errorf("failed to parse response: %w", err)
	}

	// Save JWT to file
	jwtPath := filepath.Join("./jwt", "token.json")
	err = os.WriteFile(jwtPath, []byte(otpResponse.AccesToken), 0644)
	if err != nil {
		fmt.Println("Error saving JWT to file:", err)
		return err
	}

	fmt.Printf("JWT successfully saved to %s\n", jwtPath)
	return nil
}
