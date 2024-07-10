package sync

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"io/fs"
	"os"
	"strings"
)

// getModTimes returns a list of all entry modification times
func getModTimes(entryList []string) []int64 {
	var modList []int64
	for _, file := range entryList {
		modTime, _ := os.Stat(backend.TargetLocationFormat(file))
		modList = append(modList, modTime.ModTime().Unix())
	}

	return modList
}

// genDeviceIDList returns a pointer to a slice of all registered device IDs
func genDeviceIDList() *[]fs.DirEntry {
	// create a slice of all registered devices
	deviceIDList, err := os.ReadDir(backend.ConfigDir + backend.PathSeparator + "devices")
	if err != nil {
		fmt.Println(backend.AnsiError + "Failed to read the devices directory: " + err.Error() + backend.AnsiReset)
		os.Exit(1)
	}
	return &deviceIDList
}

// ShearLocal removes the target file or directory from the local system
// returns: deviceID (on client), for use in ShearRemoteFromClient
// if the local system is a server, it will also add the target to the deletions list for all clients (except the requesting client)
// this function should only be used directly by the server binary
func ShearLocal(targetLocationIncomplete, clientDeviceID string) string {
	// determine if running on a server
	var onServer bool
	if clientDeviceID != "" {
		onServer = true
	}

	deviceIDList := genDeviceIDList()

	// add the sheared target (incomplete, vanity) to the deletions list (if running on a server)
	if onServer {
		for _, device := range *deviceIDList {
			if device.Name() != clientDeviceID {
				_, err := os.Create(backend.ConfigDir + backend.PathSeparator + "deletions" + backend.PathSeparator + device.Name() + "\x1e" + strings.ReplaceAll(targetLocationIncomplete, "/", "\x1d"))
				if err != nil {
					// do not print error as there is currently no way of seeing server-side errors
					// failure to add the target to the deletions list will exit the program and result in a client re-uploading the target (non-critical)
					os.Exit(1)
				}
			}
		}
	}

	// get the full targetLocation path and remove the target
	targetLocationComplete := backend.TargetLocationFormat(targetLocationIncomplete)
	if !onServer { // error if target does not exist on client, needed because os.RemoveAll does not return an error if target does not exist
		backend.TargetIsFile(targetLocationComplete, true, 0)
	}
	err := os.RemoveAll(targetLocationComplete)
	if err != nil {
		fmt.Println(backend.AnsiError + "Failed to remove local target: " + err.Error() + backend.AnsiReset)
		os.Exit(1)
	}

	if !onServer && len(*deviceIDList) > 0 { // return the device ID if running on the client and a device ID exists (online mode)
		return (*deviceIDList)[0].Name()
	}
	return ""

	// do not exit program, as this function is used as part of ShearRemoteFromClient
}

// RenameLocal renames oldLocationIncomplete to newLocationIncomplete on the local system
// this function should only be used directly by the server binary
func RenameLocal(oldLocationIncomplete, newLocationIncomplete string) {
	// get full paths for both locations
	oldLocation := backend.TargetLocationFormat(oldLocationIncomplete)
	newLocation := backend.TargetLocationFormat(newLocationIncomplete)

	// ensure oldLocation exists
	backend.TargetIsFile(oldLocation, true, 0)

	// ensure newLocation does not exist
	_, isAccessible := backend.TargetIsFile(newLocation, false, 0)
	if isAccessible {
		fmt.Println(backend.AnsiError + "\"" + newLocation + "\" already exists" + backend.AnsiReset)
		os.Exit(1)
	}

	// rename oldLocation to newLocation
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		fmt.Println(backend.AnsiError + "Failed to rename - does the target containing directory exist?" + backend.AnsiReset)
	}

	// do not exit program, as this function is used as part of RenameRemoteFromClient
}

// AddFolderLocal creates a new entry-containing directory on the local system
// this function should only be used directly by the server binary
func AddFolderLocal(targetLocationIncomplete string) {
	// get the full targetLocation path and create the target
	targetLocationComplete := backend.TargetLocationFormat(targetLocationIncomplete)
	err := os.Mkdir(targetLocationComplete, 0700)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println(backend.AnsiError + "Directory already exists" + backend.AnsiReset)
			os.Exit(1)
		} else {
			fmt.Println(backend.AnsiError + "Failed to create directory: " + err.Error() + backend.AnsiReset)
			os.Exit(1)
		}
	}

	// do not exit program, as this function is used as part of AddFolderRemoteFromClient
}
