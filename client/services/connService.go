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

func SendConnCredentials(phoneNumber, emaillAddress string, newUser bool) (string, error) {
	url := os.Getenv("ServerConnUrl")
	if url == "" {
		InfoLogs("Couldn't get the url from the env file")
		return "", fmt.Errorf("environment variable ServerConnUrl is not set")
	}

	payloadStruct := model.SignupRequest{
		StructPhone:     phoneNumber,
		StructEmaill:    emaillAddress,
		StructAccStatus: newUser,
	}

	reqPayload, err := json.Marshal(payloadStruct)
	if err != nil {
		fmt.Println("Error marsheling the payload structure | ERROR: ", err)
		Check(err)
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqPayload))
	if err != nil {
		fmt.Println("Error making a new request | ERROR: ", err)
		Check(err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		Check(err)
		fmt.Println("Couldn't send the http request | ERROR :", err)
		return "", nil
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected HTTP status: %d %s", resp.StatusCode, resp.Status)
	}

	var serverResponse model.UUIDResponse
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		Check(err)
		fmt.Println("Couldn't read the response body | ERROR:", err)
		return "", err
	}

	if err := json.Unmarshal(bodyBytes, &serverResponse); err != nil {
		fmt.Println("Couldn't unmarshal the uuid response | ERROR:", err)
		Check(err)
		return "", err
	}

	return serverResponse.StructUUID, nil
}
