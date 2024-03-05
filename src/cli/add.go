package cli

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/offline"
	"os"
)

// AddEntry creates a new entry at targetLocation by taking user input via CLI prompts
func AddEntry(targetLocation string, hidePassword bool, isNote bool) {
	_, isAccessible := offline.TargetIsFile(targetLocation, false, 0)
	if isAccessible {
		fmt.Println(offline.AnsiError + "Target location already exists" + offline.AnsiReset)
		os.Exit(1)
	}

	var unencryptedEntry []string

	if !isNote {
		username := input("Username: ")
		password := inputHidden("Password: ")
		url := input("URL: ")
		// TODO prompt for optional note
		unencryptedEntry = []string{password, username, url}
	} else {
		note := newNote()
		unencryptedEntry = append([]string{"", "", ""}, note...)
		fmt.Println()
	}

	offline.WriteEntry(targetLocation, unencryptedEntry)
	fmt.Println(ansiBold + "\nEntry Preview:" + offline.AnsiReset)
	EntryReader(unencryptedEntry, hidePassword)
}
