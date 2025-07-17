package main

import (
	"github.com/Gustavo-DCosta/EchoPulse/client/cmd"
	"github.com/Gustavo-DCosta/EchoPulse/client/services"
)

func init() {
	go services.InstallationProcess()
}

func main() {
	cmd.Execute()
}
