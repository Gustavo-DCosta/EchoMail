package services

import (
	"fmt"
	"os"
	"path/filepath"
)

func saveUUID(uuid string) error {
	// Define the file path
	path := filepath.Join("session", uuid+".json")

	// Ensure the session folder exists
	if err := os.MkdirAll("session", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create session dir: %w", err)
	}

	// Open the file with append + create + write-only
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open or create file: %w", err)
	}
	defer file.Close()

	// Write the UUID as a line (or you could use JSON here)
	_, err = file.WriteString(uuid + "\n")
	if err != nil {
		return fmt.Errorf("failed to write UUID: %w", err)
	}

	return nil
}

func RunSaveUUID(uuid string) {
	if err := saveUUID(uuid); err != nil {
		fmt.Println("Error saving the uuid in /session folder | ERROR:", err)
	}
}
