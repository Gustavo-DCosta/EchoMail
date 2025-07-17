package services

import (
	"encoding/json"
	"fmt"
	"time"

	"os"
	"path/filepath"
)

type InstallState struct {
	State       string `json:"state"`
	InstalledAt string `json:"installed_at"`
	Version     string `json:"version"`
}

func Launcher() {
	installationProcess()
}

func installationProcess() {
	InstallationLogPath := "./log/installation.json"
	if !areFoldersCreated() {
		fmt.Println("Required folders were missing. Attempted to create them.")
	} else {
		fmt.Println("All required folders are present.")
	}

	if !checkInstallationFile(InstallationLogPath) {
		file, err := os.Create(InstallationLogPath)

		if err != nil {
			fmt.Println("Error creating installation file")
		}
		defer file.Close()
		writeInstallationState(InstallationLogPath)
	}
}

func areFoldersCreated() bool {
	folders := []string{"jwt", "session", "config", "log"}
	allExist := true

	for _, folder := range folders {
		fullPath := filepath.Join(".", folder)
		info, err := os.Stat(fullPath)

		if os.IsNotExist(err) {
			fmt.Printf("📁Folder missing: %s\n", fullPath)
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

func checkInstallationFile(path string) bool { // We return if the file is there or not
	var fileState bool
	_, err := os.Stat(path)
	if err == nil {
		fileState = true
	} else {
		fileState = false
	}

	return fileState
}

func writeInstallationState(path string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Err writting down the file", err)
	}

	state := InstallState{
		State:       "success",
		InstalledAt: time.Now().Format(time.RFC3339),
		Version:     "1.0.0",
	}
	data, _ := json.MarshalIndent(state, "", " ")
	_ = os.WriteFile("logs/installation.json", data, 0644)

	defer file.Close()
}
