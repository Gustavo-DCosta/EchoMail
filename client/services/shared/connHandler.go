package shared

import (
	"fmt"

	"github.com/Gustavo-DCosta/EchoMail/client/services/auth"
	//"github.com/Gustavo-DCosta/EchoMail/client/services/core"
	inoutput "github.com/Gustavo-DCosta/EchoMail/client/services/io"
	"github.com/Gustavo-DCosta/EchoMail/client/services/network"
)

func ConnHandler(newUser bool) error {
	phoneNumber, emailAddress := auth.GetCredentials(newUser)
	if phoneNumber == "" || emailAddress == "" {
		return fmt.Errorf("The input fields are empty!")
	}

	uuid, err := network.SendConnCredentials(phoneNumber, emailAddress, newUser)
	if err != nil {
		fmt.Println("an error occured, please try again", err)
		inoutput.Check(err)
		return err
	}

	go inoutput.RunSaveUUID(uuid)
	token := auth.Getotp()
	accessToken, err := network.SendOtp(uuid, token)
	if err != nil {
		fmt.Println("an error occured, please try again", err)
		inoutput.Check(err)
		return err
	}
	go inoutput.RunSaveJWT(accessToken)

	return nil
}
