package services

import (
	"encoding/json"
	"fmt"
	"os"

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

/*func CacheEmailfromFile() {

}*/
