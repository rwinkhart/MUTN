package cli

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/offline"
	"os"
)

// AddPasswordEntry prompts the user for all information relevant to a password entry and writes the entry to an encrypted file at targetLocation
func AddPasswordEntry(targetLocation string, hidePassword bool) {
	_, isAccessible := offline.TargetIsFile(targetLocation, false, 0)
	if isAccessible {
		fmt.Println(offline.AnsiError + "Target location already exists" + offline.AnsiReset)
		os.Exit(1)
	}

	username := input("Username: ")
	password := inputHidden("Password: ")
	url := input("URL: ")
	// TODO implement notes editor
	notes := ""

	unencryptedEntry := []string{password, username, url, notes}
	offline.WriteEntry(targetLocation, unencryptedEntry)
	fmt.Println(ansiBold + "\nEntry Preview:" + offline.AnsiReset)
	EntryReader(unencryptedEntry, hidePassword)
}
