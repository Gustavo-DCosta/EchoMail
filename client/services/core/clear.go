package services

import (
	"fmt"
)

func ClearUI() {
	fmt.Print("\033[2J\033[H")
}
