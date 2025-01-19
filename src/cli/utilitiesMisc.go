package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rwinkhart/libmutton/core"
	"github.com/rwinkhart/libmutton/sync"
	"golang.org/x/term"
)

// input prompts the user for input and returns the input as a string.
func input(prompt string) string {
	fmt.Print("\n" + prompt + " ")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	return strings.TrimRight(userInput, "\n\r ") // remove trailing newlines, carriage returns, and spaces
}

// inputHidden prompts the user for input and returns the input as a byte array, hiding the input from the terminal.
func inputHidden(prompt string) []byte {
	fmt.Print("\n" + prompt + " ")
	byteInput, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return byteInput
}

// inputInt prompts the user for input and returns the input as an integer.
// A maxValue of 0 will cause the function to return 0, an error - a negative maxValue will disable the maxValue check.
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

// inputBinary prompts the user with a yes/no question and returns the response as a boolean.
func inputBinary(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n" + prompt + " (y/N) ")
	char, _, _ := reader.ReadRune()
	if char == 'y' || char == 'Y' {
		return true
	}
	return false
}

// inputMenuGen prompts the user with a menu and returns the user's choice as an integer.
func inputMenuGen(prompt string, options []string) int {
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	return inputInt(prompt, len(options))
}

// writeEntryCLI writes an entry to targetLocation and previews it (errors if no data is supplied).
func writeEntryCLI(targetLocation string, unencryptedEntry []string, hideSecrets bool) {
	if core.EntryIsNotEmpty(unencryptedEntry) {
		// write the entry to the target location
		core.WriteEntry(targetLocation, unencryptedEntry)
		// preview the entry
		fmt.Println(AnsiBold + "\nEntry Preview:" + core.AnsiReset)
		EntryReader(unencryptedEntry, hideSecrets, true)
	} else {
		fmt.Println(core.AnsiError + "No data supplied for entry" + core.AnsiReset)
		os.Exit(core.ErrorTargetNotFound)
	}
}

// RunJobWrapper is a wrapper for sync.RunJob that sets the passphrase input function to inputHidden.
func RunJobWrapper(manualSync bool) {
	core.PassphraseInputFunction = inputHidden
	sync.RunJob(manualSync, false)
}
