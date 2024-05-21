package backend

import (
	"fmt"
	"os"
)

// GetOldEntryData decrypts and returns old entry data (with all required lines present)
func GetOldEntryData(targetLocation string, field int) []string {
	// ensure targetLocation exists
	TargetIsFile(targetLocation, true, 2)

	// read old entry data
	unencryptedEntry := DecryptGPG(targetLocation)

	// return the old entry data with all required lines present
	if field > 0 {
		return EnsureSliceLength(unencryptedEntry, field)
	} else {
		return unencryptedEntry
	}
}

// Rename renames oldLocation to newLocation
func Rename(oldLocation string, newLocation string) {
	// ensure newLocation does not exist
	_, isAccessible := TargetIsFile(newLocation, false, 0)
	if isAccessible {
		fmt.Println(AnsiError + "\"" + newLocation + "\" already exists" + AnsiReset)
		os.Exit(1)
	}

	// rename oldLocation to newLocation
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		fmt.Println(AnsiError + "Failed to rename - does the target containing directory exists?" + AnsiReset)
	}

	// TODO implement synced renaming
	os.Exit(0)
}

// EnsureSliceLength ensures slice is long enough to contain the specified index
func EnsureSliceLength(slice []string, index int) []string {
	for len(slice) <= index {
		slice = append(slice, "")
	}
	return slice
}
