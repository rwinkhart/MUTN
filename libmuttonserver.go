package main

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"github.com/rwinkhart/MUTN/src/sync"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(backend.AnsiError + "No arguments supplied" + backend.AnsiReset) // TODO add server help menu
		os.Exit(0)
	}

	switch args[1] {
	case "fetch":
		// print all information needed for syncing to stdout for interpretation by the client
		sync.GetRemoteDataFromServer(args[2])
	case "shear":
		// shear an entry from the server and add it to the deletions directory
		deviceIDTargetLocation := strings.Split(args[2], "\x1d")
		targetLocationIncomplete := strings.ReplaceAll(strings.ReplaceAll(deviceIDTargetLocation[1], "\x1f", " "), "\x1e", backend.PathSeparator)
		sync.Shear(targetLocationIncomplete, deviceIDTargetLocation[0])
	case "addfolder":
		// add a new folder to the server
		sync.AddFolder(strings.ReplaceAll(args[2], "\x1f", " "), true)
	case "register":
		// register a new device ID
		os.Create(backend.ConfigDir + backend.PathSeparator + "devices" + backend.PathSeparator + args[2])
	case "init":
		// create the necessary directories for libmuttonserver to function TODO consider clearing out old directories first
		backend.DirInit()
		os.MkdirAll(backend.ConfigDir+backend.PathSeparator+"deletions", 0700)
	default:
		fmt.Println(backend.AnsiError + "Invalid argument" + backend.AnsiReset) // TODO add server help menu
	}
}
