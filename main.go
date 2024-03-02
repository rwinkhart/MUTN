package main

import (
	"github.com/rwinkhart/MUTN/src/cli"
	"github.com/rwinkhart/MUTN/src/offline"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	argsCount := len(args)

	// if no arguments are supplied...
	if argsCount == 0 {
		// list all entries
		cli.EntryListGen()
	} else {

		// if the first argument is an entry...
		if strings.HasPrefix(args[0], "/") {

			// store location of target entry
			targetLocation := offline.EntryRoot + offline.PathSeparator + args[0][1:]

			// entry reader shortcut (if no other arguments are supplied)
			if argsCount == 1 {
				cli.EntryReaderShortcut(targetLocation, true)
				// perform other operations on the entry (if other arguments are supplied)
			} else if argsCount == 2 {

				switch args[1] {
				case "show", "-s":
					cli.EntryReaderShortcut(targetLocation, false)
				case "shear": // TODO offline.Shear(targetLocation), exit after run
				case "gen": // TODO offline.Gen(targetLocation), exit after run
				case "copy":
					cli.HelpCopy()
				case "edit":
					cli.HelpEdit()
				case "add":
					cli.HelpAdd()
				default:
					cli.HelpMain()
				}

			} else if argsCount == 3 {

				// declare entry field to perform target operation on

				switch args[1] {
				case "copy":
					var field uint8
					switch args[2] {
					case "password", "-p":
						field = 0
					case "username", "-u":
						field = 1
					case "url", "-l":
						field = 2
					case "note", "-n":
						field = 3
					default:
						cli.HelpCopy()
					}
					offline.CopyField(targetLocation, field)
				case "edit":
					var field rune
					field = field // TODO temporary to avoid unused variable error
					switch args[2] {
					case "password", "-p":
						field = 'p'
					case "username", "-u":
						field = 'u'
					case "url", "-l":
						field = 'l'
					case "note", "-n":
						field = 'n'
					case "rename", "-r": // TODO prompt for newLocation, offline.Rename(targetLocation, newLocation), exit after run
					default:
						cli.HelpCopy()
					}
					// TODO offline.EditEntry(targetLocation, field), exit after run
				case "add":
					var field rune
					field = field // TODO temporary to avoid unused variable error
					switch args[2] {
					case "password", "-p":
						field = 'p'
					case "note", "-n":
						field = 'n'
					case "folder", "-f": // TODO offline.AddFolder(targetLocation), exit after run
					default:
						cli.HelpAdd()
					}
					// TODO offline.AddEntry(targetLocation, field), exit after run
				case "gen":
					var field rune
					field = field // TODO temporary to avoid unused variable error
					switch args[2] {
					case "update", "-u": // TODO offline.GenUpdate(targetLocation), exit after run
					default:
						cli.HelpGen()
					}
				default:
					cli.HelpMain()
				}
			}

			// if the first argument is not an entry...
		} else {
			switch args[0] {
			case "sync": // TODO online.Sync(), exit after run
			case "init": // TODO offline.Init(), exit after run
			case "tweak": // TODO offline.Tweak(), exit after run

			case "add":
				cli.HelpAdd()
			case "edit":
				cli.HelpEdit()
			case "copy":
				cli.HelpCopy()
			case "gen":
				cli.HelpGen()
			case "version", "-v":
				cli.Version()
			default:
				cli.HelpMain()
			}
		}
	}
}
