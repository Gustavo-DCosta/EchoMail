package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Gustavo-DCosta/EchoMail/client/model"
)

func SendOtp(uuid, token string) (string, error) {
	url := os.Getenv("ServerVerificationUrl")
	if url == "" {
		return "", fmt.Errorf("environment variable ServerVerificationUrl is not set")
	}

	payloadStruct := model.VerifyOTPrequest{
		StructUuid:  uuid,
		StructToken: token,
	}

	reqPayload, err := json.Marshal(payloadStruct)
	if err != nil {
		Check(err)
		fmt.Println("Couldn't marshal the payload | ERROR :", err)
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqPayload))
	if err != nil {
		Check(err)
		fmt.Println("Error creating HTTP request:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		Check(err)
		fmt.Println("Error sending the request:", err)
		return "", err // Return the actual error
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected HTTP status: %d %s", resp.StatusCode, resp.Status)
	}

	var serverResponse model.VerifySupabaseResponse
	fmt.Println("Server HTTP status:", resp.StatusCode)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		Check(err)
		return "", err
	}

	fmt.Println("Raw response from server:", string(bodyBytes))

	if err := json.Unmarshal(bodyBytes, &serverResponse); err != nil {
		fmt.Println("Couldn't unmarshal the uuid response | ERROR:", err)
		return "", err
	}

	fmt.Println("DEBUG:	", serverResponse.StructAccessToken)

	return serverResponse.StructAccessToken, nil
}
