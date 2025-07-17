package main

// Clear the screen by printing \x0c.
import (
	"fmt"
	"strings"
	"time"

	"github.com/Gustavo-DCosta/EchoPulse/client/cmd"
	"github.com/Gustavo-DCosta/EchoPulse/client/services"
)

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

}

func main() {
	cmd.Execute()
}
