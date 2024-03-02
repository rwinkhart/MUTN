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
				cli.EntryReaderShortcut(targetLocation)
				// perform other operations on the entry (if other arguments are supplied)
			} else if argsCount == 2 {

				switch args[1] {

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
				var field rune

				switch args[1] {
				case "copy":
					switch args[2] {
					case "password", "-p":
						field = 'p'
					case "username", "-u":
						field = 'u'
					case "url", "-l":
						field = 'l'
					case "note", "-n":
						field = 'n'
					default:
						cli.HelpCopy()
					}
					offline.CopyField(targetLocation, field)
				case "edit":
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
					offline.CopyField(targetLocation, field)
				case "add":
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
