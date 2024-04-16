package main

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"github.com/rwinkhart/MUTN/src/cli"
	"os"
	"strings"
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
			targetLocation := backend.TargetLocationFormat(args[1][1:])

			// entry reader shortcut (if no other arguments are supplied)
			if argsCount == 2 {
				cli.EntryReaderDecrypt(targetLocation, true, false)
				// perform other operations on the entry (if other arguments are supplied)
			} else if argsCount == 3 || (argsCount == 4 && (args[3] == "show" || args[3] == "-s")) {
				if argsCount == 3 { // default to "password" if no field is specified (for copy, edit, and add)
					switch args[2] {
					case "show", "-s":
						cli.EntryReaderDecrypt(targetLocation, false, false)
					case "copy":
						backend.CopyArgument(targetLocation, 0, args[0])
					case "edit":
						cli.EditEntryField(targetLocation, true, 0)
					case "gen":
						cli.AddEntry(targetLocation, true, 1)
					case "add":
						cli.AddEntry(targetLocation, true, 0)
					case "shear":
						backend.Shear(targetLocation)
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

				// declare entry field to perform target operation on

				switch args[2] {
				case "copy":
					var field int // indicates which (numbered) field to copy
					switch args[3] {
					case "password", "-pw":
						field = 0
					case "username", "-u":
						field = 1
					case "url", "-l":
						field = 2
					case "note", "-n":
						field = 3
					case "totp", "-t":
						fmt.Println("TOTP code will be copied to clipboard - your clipboard will be kept up to date with the current TOTP code until this process is closed")
						field = 5 // TODO Update field after removed from notes (breaking sshyp entry compatibility)
					default:
						cli.HelpCopy()
					}
					backend.CopyArgument(targetLocation, field, args[0])
				case "edit":
					var field int // indicates which field to edit
					switch args[3] {
					case "password", "-pw":
						field = 0
					case "username", "-u":
						field = 1
					case "url", "-l":
						field = 2
					case "note", "-n":
						field = 3
					case "rename", "-r":
						cli.RenameCli(targetLocation)
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
						backend.AddFolder(targetLocation)
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
				backend.ClipClearArgument()
			case "sync":
				if !backend.Windows {
					cli.SshypSync() // TODO Remove after native sync is implemented
				} else {
					fmt.Println(backend.AnsiError + "\"sync\" is not yet implemented" + backend.AnsiReset)
				}
				os.Exit(0)
			case "init":
				cli.TempInitCli()
			case "tweak":
				fmt.Println(backend.AnsiError + "\"tweak\" is not yet implemented" + backend.AnsiReset)
				os.Exit(0)
				// TODO backend.Tweak()
			case "copy":
				cli.HelpCopy()
			case "edit":
				cli.HelpEdit()
			case "gen":
				cli.HelpGen()
			case "add":
				cli.HelpAdd()
			case "version", "-v":
				cli.Version()
			default:
				cli.HelpMain()
			}
		}
	}
}
