package main

// Clear the screen by printing \x0c.
import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Gustavo-DCosta/EchoPulse/client/cmd"
	"github.com/Gustavo-DCosta/EchoPulse/client/services"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		// Log the error but don't exit. You might want to exit
		// if your application can't run without these variables.
		log.Println("Error loading .env file, continuing with existing environment variables.")
	}
}

func init() {
	go services.Launcher()
	const col = 50

	for i := 0; i <= col; i++ {
		percent := (i * 100) / col
		bar := strings.Repeat("=", i) + strings.Repeat(" ", col-i)
		fmt.Printf("\r[%s] %d%%", bar, percent)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println() // newline at the end

	// Clear screen after animation
	time.Sleep(200 * time.Millisecond)
	services.ClearUI()
}

func main() {
	cmd.Execute()
	services.ClearUI()
	services.CenterElement("[EchoMail]", false)
	services.LockScreen()
	services.LockScreenUX()
}
