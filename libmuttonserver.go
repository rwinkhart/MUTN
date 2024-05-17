package main

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"github.com/rwinkhart/MUTN/src/cli"
	"github.com/rwinkhart/MUTN/src/sync"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		helpServer()
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
		// create the necessary directories for libmuttonserver to function
		backend.DirInit()
		os.MkdirAll(backend.ConfigDir+backend.PathSeparator+"deletions", 0700)
	case "version", "-v":
		versionServer()
	default:
		helpServer()
	}
}

func helpServer() {
	fmt.Print(cli.AnsiBold + "\nlibmuttonserver | Copyright (c) 2024 Randall Winkhart\n" + backend.AnsiReset + `
This software exists under the MIT license; you may redistribute it under certain conditions.
This program comes with absolutely no warranty; type "libmuttonserver version" for details.

` + cli.AnsiBold + "Usage:" + backend.AnsiReset + ` libmuttonserver <argument>
	
` + cli.AnsiBold + "Arguments (user):" + backend.AnsiReset + `
 help|-h                 Bring up this menu
 version|-v              Display version and license information
 init                    Create the necessary directories for libmuttonserver to function` + "\n\n")
	os.Exit(0)
}

func versionServer() {
	cli.MITLicense()
	fmt.Print(cli.AnsiBold + "\n\n              libmuttonserver" + backend.AnsiReset + " Version " + backend.LibmuttonVersion + `

           Copyright (c) 2024 Randall Winkhart` + "\n\n")
}
