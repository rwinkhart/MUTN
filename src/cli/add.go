package cli

import (
	"github.com/rwinkhart/go-boilerplate/back"
	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
	"github.com/rwinkhart/libmutton/core"
)

// AddEntry creates a new entry at realPath by taking user input via CLI prompts.
// Requires: entryType (0 = standard password entry, 1 = auto-generated password entry, 2 = note-only entry).
func AddEntry(realPath string, entryType uint8) {
	// ensure realPath is valid
	_, err := core.EntryAddPrecheck(realPath)
	if err != nil {
		other.PrintError("Failed to add entry: "+err.Error(), back.ErrorWrite)
	}

	var decSlice []string

	var password []byte
	if entryType < 2 {
		username := front.Input("Username:")

		// determine whether to generate the password
		if entryType == 0 {
			password = front.InputSecret("Password:")
		} else {
			password = inputPasswordGen()
		}

		totp := string(front.InputSecret("TOTP secret:"))
		url := front.Input("URL:")
		if front.InputBinary("Add a note to this entry?") {
			note, _ := editNote([]string{})
			decSlice = append([]string{string(password), username, totp, url}, note...)
		} else {
			decSlice = []string{string(password), username, totp, url}
		}
	} else {
		note, _ := editNote([]string{})
		decSlice = append([]string{"", "", "", ""}, note...)
	}

	// write and preview the new entry
	if password != nil {
		writeEntryCLI(realPath, decSlice, true, nil)
	} else {
		writeEntryCLI(realPath, decSlice, false, nil)
	}
}
