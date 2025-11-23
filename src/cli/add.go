package cli

import (
	"github.com/rwinkhart/go-boilerplate/back"
	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
	"github.com/rwinkhart/libmutton/core"
)

// AddEntry creates a new entry at realPath by taking user input via CLI prompts.
// Requires: entryType (0 = standard password entry, 1 = auto-generated password entry, 2 = note-only entry).
func AddEntry(realPath string, hideSecrets bool, entryType uint8) {
	// ensure realPath is valid
	_, err := core.EntryAddPrecheck(realPath)
	if err != nil {
		other.PrintError("Failed to add entry: "+err.Error(), back.ErrorWrite)
	}

	var decryptedEntry []string

	var password string
	if entryType < 2 {
		username := front.Input("Username:")

		// determine whether to generate the password
		if entryType == 0 {
			password = string(front.InputHidden("Password:"))
		} else {
			password = inputPasswordGen()
		}

		totp := string(front.InputHidden("TOTP secret:"))
		url := front.Input("URL:")
		if front.InputBinary("Add a note to this entry?") {
			note, _ := editNote([]string{})
			decryptedEntry = append([]string{password, username, totp, url}, note...)
		} else {
			decryptedEntry = []string{password, username, totp, url}
		}
	} else {
		note, _ := editNote([]string{})
		decryptedEntry = append([]string{"", "", "", ""}, note...)
	}

	// write and preview the new entry
	if password != "" {
		writeEntryCLI(realPath, decryptedEntry, hideSecrets, true, "")
	} else {
		writeEntryCLI(realPath, decryptedEntry, hideSecrets, false, "")
	}

}
