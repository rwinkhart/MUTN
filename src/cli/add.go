package cli

import (
	"github.com/rwinkhart/libmutton/core"
)

// AddEntry creates a new entry at targetLocation by taking user input via CLI prompts.
// Requires: entryType (0 = standard password entry, 1 = auto-generated password entry, 2 = note-only entry).
func AddEntry(targetLocation string, hideSecrets bool, entryType uint8) {
	// ensure targetLocation is valid
	core.EntryAddPrecheck(targetLocation)

	var unencryptedEntry []string

	if entryType < 2 {
		username := input("Username:")

		// determine whether to generate the password
		var password string
		if entryType == 0 {
			password = string(inputHidden("Password:"))
		} else {
			password = inputPasswordGen()
		}

		totp := string(inputHidden("TOTP secret:"))
		url := input("URL:")
		if inputBinary("Add a note to this entry?") {
			note, _ := editNote([]string{})
			unencryptedEntry = append([]string{password, username, totp, url}, note...)
		} else {
			unencryptedEntry = []string{password, username, totp, url}
		}
	} else {
		note, _ := editNote([]string{})
		unencryptedEntry = append([]string{"", "", "", ""}, note...)
	}

	// write and preview the new entry
	writeEntryCLI(targetLocation, unencryptedEntry, hideSecrets)
}
