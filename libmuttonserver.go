package main

import (
	"bufio"
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"github.com/rwinkhart/MUTN/src/cli"
	"github.com/rwinkhart/MUTN/src/sync"
	"os"
	"strconv"
	"strings"
)

// Field separator key:
// \x1d = path separator
// \x1e = space/list separator
// \x1f = misc. field separator (if \x1e is already used)

func main() {
	args := os.Args
	if len(args) < 2 {
		helpServer()
	}

	// check if stdin was provided
	stdinInfo, _ := os.Stdin.Stat()
	stdinPresent := stdinInfo.Mode()&os.ModeNamedPipe != 0

	var stdin []string
	if stdinPresent {
		// read stdin, appending each line to the stdin slice
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin = append(stdin, scanner.Text())
		}
	}

	switch args[1] {
	case "fetch":
		// print all information needed for syncing to stdout for interpretation by the client
		// stdin[0] is expected to be the device ID
		sync.GetRemoteDataFromServer(stdin[0])
	case "rename":
		// move an entry to a new location before using fallthrough to add its previous iteration to the deletions directory
		// stdin[0] is evaluated after fallthrough
		// stdin[1] is expected to be the OLD incomplete target location with "\x1d" representing path separators - always pass in UNIX format
		// stdin[2] is expected to be the NEW incomplete target location with "\x1d" representing path separators - always pass in UNIX format
		sync.RenameLocal(strings.ReplaceAll(stdin[1], "\x1d", "/"), strings.ReplaceAll(stdin[2], "\x1d", "/"))
		fallthrough // fallthrough to add the old entry to the deletions directory
	case "shear":
		// shear an entry from the server and add it to the deletions directory
		// stdin[0] is expected to be the device ID
		// stdin[1] is expected to be the incomplete target location with "\x1d" representing path separators - always pass in UNIX format
		sync.ShearLocal(strings.ReplaceAll(stdin[1], "\x1d", "/"), stdin[0])
	case "addfolder":
		// add a new folder to the server
		// stdin[0] is expected to be the incomplete target location with "\x1d" representing path separators - always pass in UNIX format
		sync.AddFolderLocal(strings.ReplaceAll(stdin[0], "\x1d", "/"))
	case "register":
		// register a new device ID
		// stdin[0] is expected to be the device ID
		_, _ = os.Create(backend.ConfigDir + backend.PathSeparator + "devices" + backend.PathSeparator + stdin[0]) // error ignored; failure unlikely to occur if init was successful; "register" is not a user-facing argument and thus the error would not be visible
		// print EntryRoot and bool indicating OS type to stdout for client to store in config
		fmt.Print(backend.EntryRoot + "\x1e" + strconv.FormatBool(backend.IsWindows))
	case "init":
		// create the necessary directories for libmuttonserver to function
		backend.DirInit(false)
		_ = os.MkdirAll(backend.ConfigDir+backend.PathSeparator+"deletions", 0700) // error ignored; failure would have occurred by this point in backend.DirInit
		fmt.Println("libmuttonserver directories initialized")
	case "version":
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
 help                    Bring up this menu
 version                 Display version and license information
 init                    Create the necessary directories for libmuttonserver to function` + "\n\n")
	os.Exit(0)
}

func versionServer() {
	cli.MITLicense()
	fmt.Print(cli.AnsiBold + "\n\n              libmuttonserver" + backend.AnsiReset + " Version " + backend.LibmuttonVersion + `

           Copyright (c) 2024 Randall Winkhart` + "\n\n")
}
