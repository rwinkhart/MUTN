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
func GetRemoteDataFromServer(clientDeviceID string, clientIsWindows bool) {
	entryList, dirList := WalkEntryDir()
	modList := getModTimes(entryList)
	deletionsList, err := os.ReadDir(backend.ConfigDir + backend.PathSeparator + "deletions")
	if err != nil {
		fmt.Println(backend.AnsiError + "Failed to read the deletions directory: " + err.Error() + backend.AnsiReset)
		os.Exit(1)
	}

	// determine path separator expected by the client
	var clientPathSeparator string
	switch clientIsWindows {
	case false:
		clientPathSeparator = "/"
	case true:
		clientPathSeparator = "\\"
	}

	// print the lists to stdout

	// entry list
	for _, entry := range entryList {
		printWithClientPathSeparator(entry, clientPathSeparator, clientIsWindows)
	}

	// modification time list
	fmt.Print("\x1d")
	for _, mod := range modList {
		fmt.Print("\x1f")
		fmt.Print(mod)
	}

	// directory/folder list
	fmt.Print("\x1d")
	for _, dir := range dirList {
		printWithClientPathSeparator(dir, clientPathSeparator, clientIsWindows)
	}

	// deletions list
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

// printWithClientPathSeparator prints the given string with the client's path separator
// it does not alter the string if the client and server are of the same OS family
func printWithClientPathSeparator(printable, clientPathSeparator string, clientIsWindows bool) {
	if clientIsWindows == backend.IsWindows {
		fmt.Print("\x1f" + printable)
	} else {
		fmt.Print("\x1f" + strings.ReplaceAll(printable, backend.PathSeparator, clientPathSeparator))
	}
}
