package sync

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"os"
)

// GetRemoteDataFromServer prints to stdout the remote entries, mod times, folders, and deletions
// lists in output are separated by "\x1f"
// output is meant to be captured over SSH for interpretation by the client
func GetRemoteDataFromServer() {
	entryList, dirList := WalkEntryDir()
	modList := getModTimes(entryList)
	deletionsList, err := os.ReadDir(backend.ConfigDir + "/deletions")
	if err != nil {
		fmt.Println(backend.AnsiError + "Failed to read the deletions directory: " + err.Error() + backend.AnsiReset)
		os.Exit(1)
	}

	// print the lists to stdout
	for _, entry := range entryList {
		fmt.Print("\n" + entry)
	}
	fmt.Print("\x1f")
	for _, mod := range modList {
		fmt.Print("\n")
		fmt.Print(mod)
	}
	fmt.Print("\x1f")
	for _, dir := range dirList {
		fmt.Print("\n" + dir)
	}
	fmt.Print("\x1f")
	for _, deletion := range deletionsList {
		fmt.Print("\n" + deletion.Name())
	}
}
