package services

import (
	"fmt"
	"os"
	"path/filepath"
)

func InstallationProcess() {
	if !areFoldersCreated() {
		fmt.Println("Required folders were missing. Attempted to create them.")
	} else {
		fmt.Println("All required folders are present.")
	}
}

func areFoldersCreated() bool {
	folders := []string{"jwt", "session", "config", "log"}
	allExist := true

	for _, folder := range folders {
		fullPath := filepath.Join(".", folder)
		info, err := os.Stat(fullPath)

		if os.IsNotExist(err) {
			fmt.Printf("📁 Folder missing: %s\n", fullPath)
			err := os.MkdirAll(fullPath, 0755)
			if err != nil {
				fmt.Printf("Problem creating folder %s: %v\n", fullPath, err)
				os.Exit(1)
			}
			allExist = false
		} else if err != nil {
			fmt.Printf("Error accessing %s: %v\n", fullPath, err)
			os.Exit(1)
		} else if !info.IsDir() {
			fmt.Printf("Not a directory: %s\n", fullPath)
			os.Exit(1)
		}
	}
	return allExist
}
