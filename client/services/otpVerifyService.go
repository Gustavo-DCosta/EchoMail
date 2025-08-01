package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Gustavo-DCosta/EchoPulse/client/model"
)

func SendOtp(uuid, token string) (string, error) {
	url := os.Getenv("ServerVerificationUrl")
	if url == "" {
		return "", fmt.Errorf("environment variable ServerConnUrl is not set")
	}

	payloadStruct := model.VerifyOTPrequest{
		StructUuid:  uuid,
		StructToken: token,
	}

	reqPayload, err := json.Marshal(payloadStruct)
	if err != nil {
		fmt.Println("Couldn't marshal the payload | ERROR :", err)
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqPayload))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending the request")
		return "", nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected HTTP status: %d %s", resp.StatusCode, resp.Status)
	}

	var serverResponse model.VerifySupabaseResponse
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Couldn't read the response body | ERROR:", err)
		return "", err
	}

	if err := json.Unmarshal(bodyBytes, &serverResponse); err != nil {
		fmt.Println("Couldn't unmarshal the uuid response | ERROR:", err)
		return "", err
	}

	return serverResponse.StructAccessToken, nil
}
