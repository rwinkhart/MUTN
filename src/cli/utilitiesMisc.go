package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rwinkhart/libmutton/core"
	"github.com/rwinkhart/libmutton/sync"
	"github.com/rwinkhart/rcw/daemon"
	"github.com/rwinkhart/rcw/wrappers"
	"golang.org/x/term"
)

// GetRCWPassphrase retrieves the RCW passphrase from the most accessible source.
// It first attempts to retrieve from the RCW daemon, then from user input.
// If it does not find an RCW daemon, it launches one to cache the passphrase.
func GetRCWPassphrase() []byte {
	// request passphrase from RCW daemon
	passphrase := daemon.CallDaemonIfOpen()

	// if no passphrase was retrieved, request it from the user
	if passphrase == nil {
		for {
			passphrase = inputHidden("RCW Passphrase:")
			err := wrappers.RunSanityCheck(core.ConfigDir+"/sanity.rcw", passphrase)
			if err != nil {
				fmt.Println(core.AnsiError + "Failed to encrypt - " + err.Error() + core.AnsiReset)
				continue
			}
			break
		}
		// since no daemon was found earlier, start a new one
		core.LaunchRCWDProcess(string(passphrase))
	}
	return passphrase
}

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

// inputPasswordGen prompts the user for password generation parameters and returns a generated password as a string.
func inputPasswordGen() string {
	passLength := inputInt("Password length:", -1)
	fmt.Println()
	passCharset := uint8(inputMenuGen("Password complexity:", []string{"Simple", "Complex", "Ultra Complex (not compatible with many services)"}))
	var complexity float64
	switch passCharset {
	case 1:
		complexity = 0 // simple
		// 1 indicates string gen for filenames, but since complexity is 0, only the base charset is used
	default:
		complexity = 0.2 // (ultra) complex
		// 2 and 3 indicate complex and ultra complex charsets, respectively
	}
	return core.StringGen(passLength, complexity, passCharset)
}

// writeEntryCLI writes an entry to targetLocation and previews it (errors if no data is supplied).
func writeEntryCLI(targetLocation string, unencryptedEntry []string, hideSecrets bool) {
	if core.EntryIsNotEmpty(unencryptedEntry) {
		// write the entry to the target location
		core.WriteEntry(targetLocation, []byte(strings.Join(unencryptedEntry, "\n")), GetRCWPassphrase())
		// preview the entry
		fmt.Println(AnsiBold + "\nEntry Preview:" + core.AnsiReset)
		EntryReader(unencryptedEntry, hideSecrets, true)
	} else {
		core.PrintError("No data supplied for entry", core.ErrorTargetNotFound, true)
	}
}

// RunJobWrapper is a wrapper for sync.RunJob that sets the passphrase input function to inputHidden.
func RunJobWrapper(manualSync bool) {
	core.PassphraseInputFunction = inputHidden
	sync.RunJob(manualSync, false)
}
