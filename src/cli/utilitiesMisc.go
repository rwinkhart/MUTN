package cli

import (
	"bufio"
	"fmt"
	"github.com/rwinkhart/MUTN/src/offline"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/exec"
	"strings"
)

// TODO remove after native sync is implemented
func SshypSync() {
	cmd := exec.Command("sshyp", "sync")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// input prompts the user for input and returns the input as a string
func input(prompt string) string {
	fmt.Print("\n" + prompt + " ")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	return strings.TrimRight(userInput, "\n\r ") // remove trailing newlines, carriage returns, and spaces
}

// inputHidden prompts the user for input and returns the input as a string, hiding the input from the terminal
func inputHidden(prompt string) string {
	fmt.Print("\n" + prompt + " ")
	byteInput, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	password := string(byteInput)
	fmt.Println()
	return password
}

// inputInt prompts the user for input and returns the input as an integer
// a maxValue of 0 will cause the function to return 0, an error - a negative maxValue will disable the maxValue check
func inputInt(prompt string, maxValue int) int {
	if maxValue == 0 {
		return 0
	}

	// loop until a valid integer is entered
	for {
		fmt.Print("\n" + prompt + " ")
		var userInput int
		_, err := fmt.Scanln(&userInput)
		if err == nil && userInput > 0 && (userInput <= maxValue || maxValue < 0) {
			return userInput
		}
	}
}

// inputBinary prompts the user with a yes/no question and returns the response as a boolean
func inputBinary(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n" + prompt + " (y/N) ")
	char, _, _ := reader.ReadRune()
	if char == 'y' || char == 'Y' {
		return true
	}
	return false
}

// inputMenuGen prompts the user with a menu and returns the user's choice as an integer
func inputMenuGen(prompt string, options []string) int {
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	return inputInt(prompt, len(options))
}

// writeEntryShortcut writes an entry to targetLocation and previews it (errors if no data is supplied)
func writeEntryShortcut(targetLocation string, unencryptedEntry []string, hidePassword bool) {
	if offline.EntryIsNotEmpty(unencryptedEntry) {
		// write the entry to the target location
		offline.WriteEntry(targetLocation, unencryptedEntry)
		// preview the entry
		fmt.Println(ansiBold + "\nEntry Preview:" + offline.AnsiReset)
		EntryReader(unencryptedEntry, hidePassword, true)
	} else {
		fmt.Println(offline.AnsiError + "No data supplied for entry" + offline.AnsiReset)
		os.Exit(1)
	}
}
