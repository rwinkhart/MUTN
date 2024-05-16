package sync

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"os"
	"strings"
)

// GetRemoteDataFromServer prints to stdout the remote entries, mod times, folders, and deletions
// lists in output are separated by "\x1d"
// output is meant to be captured over SSH for interpretation by the client
func GetRemoteDataFromServer(clientDeviceID string) {
	entryList, dirList := WalkEntryDir()
	modList := getModTimes(entryList)
	deletionsList, err := os.ReadDir(backend.ConfigDir + backend.PathSeparator + "deletions")
	if err != nil {
		fmt.Println(backend.AnsiError + "Failed to read the deletions directory: " + err.Error() + backend.AnsiReset)
		os.Exit(1)
	}

	// print the lists to stdout
	for _, entry := range entryList {
		fmt.Print("\x1f" + entry)
	}
	fmt.Print("\x1d")
	for _, mod := range modList {
		fmt.Print("\x1f")
		fmt.Print(mod)
	}
	fmt.Print("\x1d")
	for _, dir := range dirList {
		fmt.Print("\x1f" + dir)
	}
	fmt.Print("\x1d")
	for _, deletion := range deletionsList {
		// print deletion if it is relevant to the current client device
		affectedIDTargetLocationIncomplete := strings.Split(deletion.Name(), "\x1d")
		if affectedIDTargetLocationIncomplete[0] == clientDeviceID {
			fmt.Print("\x1f" + strings.ReplaceAll(affectedIDTargetLocationIncomplete[1], "\x1e", backend.PathSeparator))

			// assume successful client deletion and remove deletions file (if assumption is somehow false, worst case scenario is that the client will re-upload the deleted entry)
			os.Remove(backend.ConfigDir + backend.PathSeparator + "deletions" + backend.PathSeparator + deletion.Name())
		}
	}
}
