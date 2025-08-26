package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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
