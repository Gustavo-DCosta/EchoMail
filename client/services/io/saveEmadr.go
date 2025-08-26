package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Gustavo-DCosta/EchoMail/client/cache"
	"github.com/Gustavo-DCosta/EchoMail/client/model"
)

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
