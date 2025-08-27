package core

import (
	"encoding/json"
	"fmt"
	"time"

	"os"
	"path/filepath"

	inoutput "github.com/Gustavo-DCosta/EchoMail/client/services/io"
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
	if err := areFoldersCreated(); err != nil {
		inoutput.Check(err)
		fmt.Println("Required folders were missing. Attempted to create them.")
		return
	}

	if !checkInstallationFile(InstallationLogPath) {
		file, err := os.Create(InstallationLogPath)

		if err != nil {
			inoutput.Check(err)
			fmt.Println("Error creating installation file")
		}
		writeInstallationState(InstallationLogPath)
		defer file.Close()
	}
}

func areFoldersCreated() error {
	folders := []string{"jwt", "session", "config", "log"}

	for _, folder := range folders {
		fullPath := filepath.Join(".", folder)
		info, err := os.Stat(fullPath)

		if os.IsNotExist(err) {
			fmt.Printf("📁Folder missing: %s\n", fullPath)
			err := os.MkdirAll(fullPath, 0755)
			if err != nil {
				inoutput.Check(err)
				return fmt.Errorf("problem creating folder %s: %v", fullPath, err)
			}
		} else if err != nil {
			inoutput.Check(err)
			return fmt.Errorf("error accessing %s: %v", fullPath, err)
		} else if !info.IsDir() {
			return fmt.Errorf("not a directory: %s", fullPath)
		}
	}
	return nil
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
	state := InstallState{
		State:       "success",
		InstalledAt: time.Now().Format(time.RFC3339),
		Version:     "Engine version 0.1",
	}

	data, err := json.MarshalIndent(state, "", " ")
	if err != nil {
		inoutput.Check(err)
		fmt.Println("Err marshaling state", err)
		return
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		inoutput.Check(err)
		fmt.Println("Err writing installation file", err)
		return
	}
}
