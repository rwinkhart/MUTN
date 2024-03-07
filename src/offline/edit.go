package offline

import (
	"fmt"
	"os"
)

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

	// TODO If in online mode, check if oldLocation is a directory and rename it on the server
	os.Exit(0)
}
