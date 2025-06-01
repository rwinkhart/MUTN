package main

import (
	"bytes"
	"cmp"
	"fmt"
	"os"
	"strings"

	"github.com/rwinkhart/MUTN/src/cli"
	"github.com/rwinkhart/go-boilerplate/back"
	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
	"github.com/rwinkhart/libmutton/core"
	"github.com/rwinkhart/libmutton/crypt"
	"github.com/rwinkhart/libmutton/global"
	"github.com/rwinkhart/libmutton/syncclient"
	"github.com/rwinkhart/libmutton/synccycles"
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
			targetLocation := global.TargetLocationFormat(args[1])

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
						err := core.CopyArgument(targetLocation, 0)
						if err != nil {
							other.PrintError("Failed to copy passphrase to clipboard: "+err.Error(), global.ErrorClipboard)
						}
					case "edit":
						cli.EditEntryField(targetLocation, true, 0)
					case "gen":
						cli.AddEntry(targetLocation, true, 1)
					case "add":
						cli.AddEntry(targetLocation, true, 0)
					case "shear":
						err := syncclient.ShearRemoteFromClient(args[1]) // pass the incomplete path as the server and all clients (reading from the deletions directory) will have a different home directory
						if err != nil {
							other.PrintError("Failed to shear target: "+err.Error(), back.ErrorWrite)
						}
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
					err := core.CopyArgument(targetLocation, field)
					if err != nil {
						other.PrintError("Failed to copy field to clipboard: "+err.Error(), global.ErrorClipboard)
					}
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
						isAccessible, _ := back.TargetIsFile(targetLocation, true) // error is ignored because dir/file status is irrelevant
						if !isAccessible {
							other.PrintError("Failed to access location ("+targetLocation+")", back.ErrorTargetNotFound)
						}
						cli.RenameCli(args[1]) // pass the incomplete path as the server and all clients (reading from the deletions directory) will have a different home directory
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
						err := syncclient.AddFolderRemoteFromClient(args[1]) // pass the incomplete path as the server will have a different home directory
						if err != nil {
							other.PrintError("Failed to add folder: "+err.Error(), back.ErrorWrite)
						}
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
				err := core.ClipClearArgument()
				if err != nil {
					other.PrintError("Failure occurred in clipboard clearing process: "+err.Error(), global.ErrorClipboard)
				}
			case "startrcwd":
				crypt.RCWDArgument()
			case "sync":
				_, err := syncclient.RunJob(true)
				if err != nil {
					other.PrintError("Failed to sync entries: "+err.Error(), global.ErrorSyncProcess)
				}
			case "init":
				err := core.LibmuttonInit(front.Input,
					[][3]string{{"MUTN", "textEditor", cmp.Or(front.Input("Text editor (leave blank for $EDITOR, falls back to \""+cli.FallbackEditor+"\"):"), os.Getenv("EDITOR"), cli.FallbackEditor)}},
					confirmRCWPassphrase("new"), false)
				if err != nil {
					other.PrintError("Initialization failed: "+err.Error(), 0)
				}
			case "tweak":
				choice := front.InputMenuGen("Action:", []string{"Change device ID", "Change master passphrase/Optimize entries"})
				switch choice {
				case 1:
					oldDeviceID, err := global.GetCurrentDeviceID()
					if err != nil {
						other.PrintError("Failed to get current device ID: "+err.Error(), back.ErrorRead)
					}
					_, _, err = synccycles.DeviceIDGen(oldDeviceID)
					if err != nil {
						other.PrintError("Failed to change device ID: "+err.Error(), global.ErrorSyncProcess)
					}
					fmt.Println("\nDevice ID changed successfully.")
				case 2:
					oldPassphrase := confirmRCWPassphrase("old")
					newPassphrase := confirmRCWPassphrase("new")
					fmt.Print("\nRe-encrypting entries. Please wait; do not force close this process.\n")
					err := core.EntryRefresh(oldPassphrase, newPassphrase, false)
					if err != nil {
						other.PrintError("Re-encryption failed: "+err.Error(), global.ErrorEncryption)
					}
					fmt.Println("\nRe-encryption complete.")
				}
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

// confirmRCWPassphrase prompts the user for a master passphrase and confirms it
// is typed correctly. Returns the passphrase as a byte slice.
func confirmRCWPassphrase(lowercasePrefix string) []byte {
	for {
		// get master passphrase
		rcwPass := front.InputHidden(strings.ToUpper(string(lowercasePrefix[0])) + lowercasePrefix[1:] + " master passphrase:")
		if !bytes.Equal(rcwPass, front.InputHidden("Confirm "+lowercasePrefix+" master passphrase:")) || len(rcwPass) == 0 {
			fmt.Println(back.AnsiError + "Passphrases do not match or passphrase is invalid" + back.AnsiReset)
			continue
		}
		return rcwPass
	}
}
