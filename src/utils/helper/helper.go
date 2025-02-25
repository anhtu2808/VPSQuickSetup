package helper

import (
	"fmt"
	"os"
)

func ClearScreen() {
	os.Stdout.Sync()
	os.Stdout.WriteString("\033c")
}

func PauseScreen() {
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
}
