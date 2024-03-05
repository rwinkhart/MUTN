package cli

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

// input prompts the user for input and returns the input as a string
func input(prompt string) string {
	fmt.Print("\n" + prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

// inputHidden prompts the user for input and returns the input as a string, hiding the input from the terminal
func inputHidden(prompt string) string {
	fmt.Print("\n" + prompt)
	byteInput, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	password := string(byteInput)
	fmt.Println()
	return password
}
