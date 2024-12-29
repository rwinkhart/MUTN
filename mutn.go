package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rwinkhart/MUTN/src/cli"
	"github.com/rwinkhart/libmutton/core"
	"github.com/rwinkhart/libmutton/sync"
)

func main() {

	args := os.Args
	argsCount := len(args)

	// if no arguments are supplied...
	if argsCount == 1 {
		// list all entries
		cli.EntryListGen()
	} else {

		// if the first argument is an entry...
		if strings.HasPrefix(args[1], "/") {

			// store location of target entry
			targetLocation := core.TargetLocationFormat(args[1])

			// entry reader shortcut (if no other arguments are supplied)
			if argsCount == 2 {
				cli.EntryReaderDecrypt(targetLocation, true)

				// perform other operations on the entry (if other arguments are supplied)
			} else if argsCount == 3 || (argsCount == 4 && (args[3] == "show" || args[3] == "-s")) {
				if argsCount == 3 { // default to "password" if no field is specified (for copy, edit, and add)
					switch args[2] {
					case "show", "-s":
						cli.EntryReaderDecrypt(targetLocation, false)
					case "copy":
						core.CopyArgument(targetLocation, 0)
					case "edit":
						cli.EditEntryField(targetLocation, true, 0)
					case "gen":
						cli.AddEntry(targetLocation, true, 1)
					case "add":
						cli.AddEntry(targetLocation, true, 0)
					case "shear":
						sync.ShearRemoteFromClient(args[1], false) // pass the incomplete path as the server and all clients (reading from the deletions directory) will have a different home directory
					default:
						cli.HelpMain()
					}
				} else { // handle "show" or "-s" argument for gen, edit, and add
					switch args[2] {
					case "edit":
						cli.EditEntryField(targetLocation, false, 0)
					case "gen":
						cli.AddEntry(targetLocation, false, 1)
					case "add":
						cli.AddEntry(targetLocation, false, 0)
					default:
						cli.HelpMain()
					}
				}

			} else if argsCount >= 4 {

				switch args[2] {
				case "copy":
					var field int // indicates which (numbered) field to copy
					switch args[3] {
					case "password", "-pw":
						field = 0
					case "username", "-u":
						field = 1
					case "totp", "-t":
						field = 2
					case "url", "-l":
						field = 3
					case "note", "-n":
						field = 4
					default:
						cli.HelpCopy()
					}
					core.CopyArgument(targetLocation, field)
				case "edit":
					var field int // indicates which (numbered) field to edit
					switch args[3] {
					case "password", "-pw":
						field = 0
					case "username", "-u":
						field = 1
					case "totp", "-t":
						field = 2
					case "url", "-l":
						field = 3
					case "note", "-n":
						field = 4
					case "rename", "-r":
						core.TargetIsFile(targetLocation, true, 0) // ensure location exists before prompting for new location
						cli.RenameCli(args[1])                     // pass the incomplete path as the server and all clients (reading from the deletions directory) will have a different home directory
					default:
						cli.HelpEdit()
					}
					if argsCount == 4 {
						cli.EditEntryField(targetLocation, true, field)
					} else {
						switch args[4] {
						case "show", "-s":
							cli.EditEntryField(targetLocation, false, field)
						default:
							cli.EditEntryField(targetLocation, true, field)
						}
					}
				case "gen":
					if argsCount == 4 {
						switch args[3] {
						case "update", "-u":
							cli.GenUpdate(targetLocation, true)
						default:
							cli.HelpGen()
						}
					} else if args[3] == "update" || args[3] == "-u" {
						switch args[4] {
						case "show", "-s":
							cli.GenUpdate(targetLocation, false)
						default:
							cli.GenUpdate(targetLocation, true)
						}
					}
					cli.HelpGen()
				case "add":
					switch args[3] {
					case "password", "-pw":
						if argsCount == 4 {
							cli.AddEntry(targetLocation, true, 0)
						} else {
							switch args[4] {
							case "show", "-s":
								cli.AddEntry(targetLocation, false, 0)
							default:
								cli.AddEntry(targetLocation, true, 0)
							}
						}
					case "note", "-n":
						cli.AddEntry(targetLocation, true, 2)
					case "folder", "-f":
						sync.AddFolderRemoteFromClient(args[1], false) // pass the incomplete path as the server will have a different home directory
					default:
						cli.HelpAdd()
					}
				default:
					cli.HelpMain()
				}
			}

			// if the first argument is not an entry...
		} else {
			switch args[1] {
			case "clipclear":
				core.ClipClearArgument()
			case "sync":
				cli.RunJobWrapper(true)
			case "init":
				cli.TempInitCli()
			case "tweak":
				fmt.Println(core.AnsiError + "\"tweak\" is not yet implemented" + core.AnsiReset)
				os.Exit(0)
			case "copy":
				cli.HelpCopy()
			case "edit":
				cli.HelpEdit()
			case "gen":
				cli.HelpGen()
			case "add":
				cli.HelpAdd()
			case "version":
				cli.Version()
			default:
				cli.HelpMain()
			}
		}
	}
}
