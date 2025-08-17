/*
* File to write log files when an error happens
* The file will be stored on log/err.json
 */

package services

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Gustavo-DCosta/EchoMail/client/model"
)

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
