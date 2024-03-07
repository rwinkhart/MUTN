package cli

import "github.com/rwinkhart/MUTN/src/offline"

// RenameCli renames an entry at oldLocation to a new location (user input)
func RenameCli(oldLocation string) {
	// ensure targetLocation exists
	offline.TargetIsFile(oldLocation, true, 0)

	// prompt user for new location and rename
	newLocation := offline.EntryRoot + offline.PathSeparator + input("New location:")
	offline.Rename(oldLocation, newLocation)

	// exit is done from offline.Rename
}
