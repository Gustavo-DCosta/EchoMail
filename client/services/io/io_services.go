package inoutput

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"os"
	"path/filepath"

	"github.com/Gustavo-DCosta/EchoMail/client/cache"
	"github.com/Gustavo-DCosta/EchoMail/client/model"
)

func SaveAccessToken(accessToken string) error {
	// Build file path
	dir := "jwt"
	filename := "jwt.json"
	path := filepath.Join(dir, filename)

	// Ensure folder exists
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create jwt folder: %w", err)
	}

	// Create the struct with the token
	tokenObj := model.AccessTkJsonObject{
		StructAccessTk: accessToken,
	}

	// Marshal into JSON
	jsonData, err := json.MarshalIndent(tokenObj, "", "  ")
	if err != nil {
		Check(err)
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Write to file (overwrites if exists)
	if err := os.WriteFile(path, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	// deleted debug message
	return nil
}

func RunSaveJWT(acessToken string) {
	if err := SaveAccessToken(acessToken); err != nil {
		Check(err)
		fmt.Println("Couldn't save the jwt | ERROR: ", err)
	}
}

func saveEmaillAdr(emailAddress string) {
	path := "config/email.json"

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Couldn't open email file |ERROR|", err)
		return
	}
	defer file.Close()

	entry := model.EmailObject{
		StructEmailObject: emailAddress,
	}

	encoder := json.NewEncoder(file)
	encoder.SetEscapeHTML(false) // this fixes the \u003c \u003e problem
	// appears now <>

	if err := encoder.Encode(entry); err != nil {
		fmt.Println("Couldn't encode JSON object |ERROR|", err)
	}
}

func CacheEmailfromFile() {
	// Should check if the file is there
	// If it isn't it returns
	// If the file is indeed there it reads rhe json object and decodes the email adress
	// And then if the email adress is empty it returns
	// If not it caches it

	_, err := os.Stat("config/email.json")

	if errors.Is(err, os.ErrNotExist) {
		return
	} else {
		content, err := ioutil.ReadFile("config/email.json")
		if err != nil {
			fmt.Println("Couldn't open the configuration email json file |ERROR|", err)
			return
		}

		var data model.EmailObject
		err = json.Unmarshal(content, &data)
		if err != nil {
			fmt.Println("Error unmarsheling the configuration email json file |ERROR|", err)
			return
		}
		if data.StructEmailObject == "" {
			return
		} else {
			cache.Set("UserEmail", data.StructEmailObject)
		}
	}
}

func saveUUID(uuid string) error {
	// Define the file path
	path := filepath.Join("session", uuid+".json")

	// Ensure the session folder exists
	if err := os.MkdirAll("session", os.ModePerm); err != nil {
		Check(err)
		return fmt.Errorf("failed to create session dir: %w", err)
	}

	// Open the file with append + create + write-only
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Check(err)
		return fmt.Errorf("failed to open or create file: %w", err)
	}
	defer file.Close()

	// Write the UUID as a line (or you could use JSON here)
	_, err = file.WriteString(uuid + "\n")
	if err != nil {
		Check(err)
		return fmt.Errorf("failed to write UUID: %w", err)
	}

	return nil
}

func RunSaveUUID(uuid string) {
	if err := saveUUID(uuid); err != nil {
		Check(err)
		fmt.Println("Error saving the uuid in /session folder | ERROR:", err)
	}
}

func Check(e error) {
	if e == nil {
		return
	}

	path := "log/err.json"
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Couldn't open logs file |ERROR|", err)
		return
	}
	defer file.Close()

	object := model.ErrorObject{
		StructTime:  time.Now(),
		StructError: e.Error(), // convert error to string
	}

	writeObj, err := json.Marshal(object)
	if err != nil {
		fmt.Println("Couldn't marshal log object |ERROR|", err)
		return
	}

	// Append a newline so logs are readable & valid JSONL (JSON lines)
	_, err = file.WriteString(string(writeObj) + "\n")
	if err != nil {
		fmt.Println("Couldn't write to logs file |ERROR|", err)
	}
}

func InfoLogs(message string) {
	path := "log/log.json"

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open the file |ERROR|", err)
		return
	}
	defer file.Close()

	entry := model.LogsObject{
		StructTimeStamp: time.Now(),
		StructMessage:   message,
	}

	encoder := json.NewEncoder(file)
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(entry); err != nil {
		fmt.Println("Failed to decode the message |ERROR|", err)
	}
}
