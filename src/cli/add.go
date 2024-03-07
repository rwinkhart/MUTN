package cli

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/offline"
	"os"
)

// AddEntry creates a new entry at targetLocation by taking user input via CLI prompts
// entryType: 0 = standard (password), 1 = auto-generated password, 2 = note
func AddEntry(targetLocation string, hidePassword bool, entryType uint8) {
	_, isAccessible := offline.TargetIsFile(targetLocation, false, 0)
	if isAccessible {
		fmt.Println(offline.AnsiError + "Target location already exists" + offline.AnsiReset)
		os.Exit(1)
	}

	var unencryptedEntry []string

	if entryType < 2 {
		username := input("Username:")

		// determine whether to generate the password
		var password string
		if entryType == 0 {
			password = inputHidden("Password:")
		} else {
			password = offline.StringGen(inputInt("Password length:"), inputBinary("Generate a complex (special characters) password?"), 0.2)
		}

		url := input("URL:")
		if inputBinary("Add a note to this entry?") {
			note := newNote()
			unencryptedEntry = append([]string{password, username, url}, note...)
		} else {
			unencryptedEntry = []string{password, username, url}
		}
	} else {
		note := newNote()
		unencryptedEntry = append([]string{"", "", ""}, note...)
	}

	// write and preview the new entry
	if offline.EntryIsNotEmpty(unencryptedEntry) {
		offline.WriteEntry(targetLocation, unencryptedEntry)
		fmt.Println(ansiBold + "\nEntry Preview:" + offline.AnsiReset)
		EntryReader(unencryptedEntry, hidePassword)
	} else {
		fmt.Println(offline.AnsiError + "No data supplied for entry" + offline.AnsiReset)
		os.Exit(1)
	}
}
