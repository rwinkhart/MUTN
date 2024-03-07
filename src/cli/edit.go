package cli

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/offline"
	"os"
)

// RenameCli renames an entry at oldLocation to a new location (user input)
func RenameCli(oldLocation string) {
	// ensure targetLocation exists
	offline.TargetIsFile(oldLocation, true, 0)

	// prompt user for new location and rename
	newLocation := offline.EntryRoot + offline.PathSeparator + input("New location:")
	offline.Rename(oldLocation, newLocation)

	// exit is done from offline.Rename
}

// EditEntry edits a field of an entry at targetLocation (user input), does not allow for editing notes
func EditEntry(targetLocation string, hidePassword bool, field int) {
	// ensure targetLocation exists
	offline.TargetIsFile(targetLocation, true, 2)

	// read old entry data
	unencryptedEntry := offline.DecryptGPG(targetLocation)

	// ensure slice is long enough for field
	for len(unencryptedEntry) <= field {
		unencryptedEntry = append(unencryptedEntry, "")
	}

	// edit the field
	switch field {
	case 0:
		unencryptedEntry[field] = inputHidden("Password:")
	case 1:
		unencryptedEntry[field] = input("Username:")
	case 2:
		unencryptedEntry[field] = input("URL:")
	}

	// write and preview the modified entry
	if offline.EntryIsNotEmpty(unencryptedEntry) {
		offline.WriteEntry(targetLocation, unencryptedEntry)
		fmt.Println(ansiBold + "\nEntry Preview:" + offline.AnsiReset)
		EntryReader(unencryptedEntry, hidePassword)
	} else {
		fmt.Println(offline.AnsiError + "No data supplied for entry" + offline.AnsiReset)
		os.Exit(1)
	}
}
