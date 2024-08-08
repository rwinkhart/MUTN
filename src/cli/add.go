package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/rwinkhart/libmutton/core"
)

// AddEntry creates a new entry at targetLocation by taking user input via CLI prompts
// entryType: 0 = standard (password), 1 = auto-generated password, 2 = note
func AddEntry(targetLocation string, hideSecrets bool, entryType uint8) {
	// ensure target location does not already exist
	_, isAccessible := core.TargetIsFile(targetLocation, false, 0)
	if isAccessible {
		fmt.Println(core.AnsiError + "Target location already exists" + core.AnsiReset)
		os.Exit(1)
	}

	// ensure target containing directory exists and is a directory (not a file)
	core.TargetIsFile(targetLocation[:strings.LastIndex(targetLocation, "/")], true, 1)

	var unencryptedEntry []string

	if entryType < 2 {
		username := input("Username:")

		// determine whether to generate the password
		var password string
		if entryType == 0 {
			password = inputHidden("Password:")
		} else {
			password = core.StringGen(inputInt("Password length:", -1), inputBinary("Generate a complex (special characters) password?"), 0.2, false)
		}

		totp := inputHidden("TOTP secret:")
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
	writeEntryCLI(targetLocation, unencryptedEntry, hideSecrets, false)
}
