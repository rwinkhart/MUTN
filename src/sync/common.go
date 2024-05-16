package sync

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// WalkEntryDir walks the entry directory and returns lists of all files and directories found (two separate lists)
// initCommand is used to specify to the end user how to generate the entry directory if it does not exists
func WalkEntryDir() ([]string, []string) {
	// define file/directory containing slices so that they may be accessed by the anonymous WalkDir function
	var fileList []string
	var dirList []string

	// walk entry directory
	_ = filepath.WalkDir(backend.EntryRoot,
		func(fullPath string, entry fs.DirEntry, err error) error {

			// check for errors encountered while walking directory
			if err != nil {
				if os.IsNotExist(err) {
					fmt.Println(backend.AnsiError+"The entry directory does not exist - run \""+os.Args[0], "init"+"\" to create it"+backend.AnsiReset) // TODO implement init command for libmuttonserver
				} else {
					// otherwise, print the source of the error
					fmt.Println(backend.AnsiError + "An unexpected error occurred while generating the entry list: " + err.Error() + backend.AnsiReset)
				}
				os.Exit(1)
			}

			// trim root path from each path before storing
			trimmedPath := fullPath[rootLength:]

			// create separate slices for entries and directories
			if !entry.IsDir() {
				fileList = append(fileList, trimmedPath)
			} else {
				dirList = append(dirList, trimmedPath)
			}

			return nil
		})

	return fileList, dirList
}

func getModTimes(entryList []string) []int64 {
	// get a list of all entry modification times
	var modList []int64
	for _, file := range entryList {
		modTime, _ := os.Stat(backend.EntryRoot + file)
		modList = append(modList, modTime.ModTime().Unix())
	}

	return modList
}

// Shear removes the target file or directory from the system
// if running on the server, it will also add the target to the deletions list
// if running on the client, it will call the server to add the target to the deletions list
func Shear(targetLocationIncomplete string, deviceID string) {
	// get the full targetLocation path and remove the target
	targetLocationComplete := backend.TargetLocationFormat(targetLocationIncomplete[1:])
	backend.TargetIsFile(targetLocationComplete, true, 0) // needed because os.RemoveAll does not return an error if target does not exist
	err := os.RemoveAll(targetLocationComplete)
	if err != nil {
		fmt.Println(backend.AnsiError + "Failed to remove local target: " + err.Error() + backend.AnsiReset)
		os.Exit(1)
	}

	// read the devices directory
	var deviceIDList []os.DirEntry
	deviceIDList, err = os.ReadDir(backend.ConfigDir + backend.PathSeparator + "devices")
	if err != nil {
		fmt.Println(backend.AnsiError + "Failed to read the devices directory: " + err.Error() + backend.AnsiReset)
		os.Exit(1)
	}

	// add the sheared target (incomplete, vanity) to the deletions list
	if deviceID != "" { // if running on the server...
		for _, device := range deviceIDList {
			if device.Name() != deviceID {
				_, err = os.Create(backend.ConfigDir + backend.PathSeparator + "deletions" + backend.PathSeparator + device.Name() + "\x1d" + strings.ReplaceAll(targetLocationIncomplete, backend.PathSeparator, "\x1e"))
				if err != nil {
					// do not print error as there is currently no way of seeing server-side errors
					os.Exit(1)
				}
			}
		}
	} else {
		backend.ReadConfig([]string{"sshUser"}, "0") // exits if value does not exist (indicates offline mode)
		// if running on the client in online mode...
		// determine client device ID (to send to server, avoids creating a deletion file for the client device)
		deviceID = deviceIDList[0].Name()
		// below: deviceID and targetLocationIncomplete are separated by \x1d, path separators are replaced with \x1e, and spaces are replaced with \x1f
		getSSHOutput("libmuttonserver shear "+deviceID+"\x1d"+strings.ReplaceAll(strings.ReplaceAll(targetLocationIncomplete, backend.PathSeparator, "\x1e"), " ", "\x1f"), true)
	}

	os.Exit(0)
}
