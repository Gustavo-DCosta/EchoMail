package main

// Clear the screen by printing \x0c.
import (
	"fmt"
	"log"
	"strings"
	"time"

	//	"github.com/Gustavo-DCosta/EchoMail/client/services/core"
	"github.com/Gustavo-DCosta/EchoMail/client/services/core"
	inoutput "github.com/Gustavo-DCosta/EchoMail/client/services/io"
	"github.com/Gustavo-DCosta/EchoMail/client/services/shared"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		// Log the error but don't exit. You might want to exit
		// if your application can't run without these variables.
		log.Fatal(`
									Error loading .env file at path: ./config/.env
									The app can't continue without it.
									Please contact the support team to fix this issue.
`)
	}
}

func init() {
	go core.Launcher()
	go inoutput.CacheEmailfromFile()
	// Function to get the email from the fail
	// And cache it using cache function on cache package
	// ENGLISH YES

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
	shared.ClearUI() //services.ClearUI() -> core.ClearUI()
}

func main() {
	shared.ClearUI()                          //services.ClearUI() -> core.ClearUI()
	shared.CenterElement("[EchoMail]", false) //services.CenterElement() -> core.CenterElement()
	shared.HelpCommand()                      //lockscreen -> core.LockScreen()
	core.LockScreenPrompt()                   // services.LockScreenUX() -> core.LockScreenUX()
}
