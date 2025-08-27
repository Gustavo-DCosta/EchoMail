package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Gustavo-DCosta/EchoMail/client/model"
	inoutput "github.com/Gustavo-DCosta/EchoMail/client/services/io"
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
		inoutput.Check(err)
		fmt.Println("Couldn't marshal the payload | ERROR :", err)
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqPayload))
	if err != nil {
		inoutput.Check(err)
		fmt.Println("Error creating HTTP request:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		inoutput.Check(err)
		fmt.Println("Error sending the request:", err)
		return "", err // Return the actual error
	}
	defer resp.Body.Close()
	inoutput.InfoLogs("Sent an http request to the server")

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected HTTP status: %d %s", resp.StatusCode, resp.Status)
	}

	var serverResponse model.VerifySupabaseResponse // removed debug message, UX upgrade
	inoutput.InfoLogs("Rceived response from the server")

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		inoutput.Check(err)
		return "", err
	}

	if err := json.Unmarshal(bodyBytes, &serverResponse); err != nil {
		fmt.Println("Couldn't unmarshal the uuid response | ERROR:", err)
		return "", err
	}

	//deleted debug message

	return serverResponse.StructAccessToken, nil
}
