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
	"github.com/rwinkhart/libmutton/age"
	"github.com/rwinkhart/libmutton/cfg"
	"github.com/rwinkhart/libmutton/clip"
	"github.com/rwinkhart/libmutton/core"
	"github.com/rwinkhart/libmutton/crypt"
	"github.com/rwinkhart/libmutton/global"
	"github.com/rwinkhart/libmutton/syncclient"
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

			// store realPath of target entry
			realPath := global.GetRealPath(args[1])

			// entry reader shortcut (if no other arguments are supplied)
			if argsCount == 2 {
				cli.EntryReaderDecrypt(realPath, true)

				// perform other operations on the entry (if other arguments are supplied)
			} else if argsCount == 3 || (argsCount == 4 && (args[3] == "show" || args[3] == "-s")) {
				if argsCount == 3 { // default to "password" if no field is specified (for copy, edit, and add)
					switch args[2] {
					case "show", "-s":
						cli.EntryReaderDecrypt(realPath, false)
					case "copy":
						err := clip.CopyShortcut(realPath, 0)
						if err != nil {
							other.PrintError("Failed to copy password to clipboard: "+err.Error(), global.ErrorClipboard)
						}
					case "edit":
						cli.EditEntryField(realPath, 0)
					case "gen":
						cli.AddEntry(realPath, 1)
					case "add":
						cli.AddEntry(realPath, 0)
					case "shear":
						err := syncclient.ShearRemote(args[1], false) // pass the incomplete path as the server and all clients (reading from the deletions directory) will have a different home directory
						if err != nil {
							other.PrintError("Failed to shear target: "+err.Error(), back.ErrorWrite)
						}
					default:
						cli.HelpMain()
					}
				} else { // handle "show" or "-s" argument for gen, edit, and add
					switch args[2] {
					case "edit":
						cli.EditEntryField(realPath, 0)
					case "gen":
						cli.AddEntry(realPath, 1)
					case "add":
						cli.AddEntry(realPath, 0)
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
					case "menu", "-m":
						cli.CopyMenu(args[1], nil, "")
					default:
						cli.HelpCopy()
					}
					err := clip.CopyShortcut(realPath, field)
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
						isAccessible, _ := back.TargetIsFile(realPath, true) // error is ignored because dir/file status is irrelevant
						if !isAccessible {
							other.PrintError("Failed to access location ("+realPath+")", back.ErrorTargetNotFound)
						}
						cli.RenameCli(args[1]) // pass the incomplete path as the server and all clients (reading from the deletions directory) will have a different home directory
					default:
						cli.HelpEdit()
					}
					if argsCount == 4 {
						cli.EditEntryField(realPath, field)
					} else {
						switch args[4] {
						case "show", "-s":
							cli.EditEntryField(realPath, field)
						default:
							cli.EditEntryField(realPath, field)
						}
					}
				case "gen":
					if argsCount == 4 {
						switch args[3] {
						case "update", "-u":
							cli.GenUpdate(realPath)
						default:
							cli.HelpGen()
						}
					} else if args[3] == "update" || args[3] == "-u" {
						switch args[4] {
						case "show", "-s":
							cli.GenUpdate(realPath)
						default:
							cli.GenUpdate(realPath)
						}
					}
					cli.HelpGen()
				case "add":
					switch args[3] {
					case "password", "-pw":
						if argsCount == 4 {
							cli.AddEntry(realPath, 0)
						} else {
							switch args[4] {
							case "show", "-s":
								cli.AddEntry(realPath, 0)
							default:
								cli.AddEntry(realPath, 0)
							}
						}
					case "note", "-n":
						cli.AddEntry(realPath, 2)
					case "folder", "-f":
						err := syncclient.AddFolderRemote(args[1]) // pass the incomplete path as the server will have a different home directory
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
				err := clip.ClearArgument()
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
					map[string]any{"mutnTextEditor": getTextEditorInput()},
					confirmRCWPassword("new"), false, false)
				if err != nil {
					other.PrintError("Initialization failed: "+err.Error(), 0)
				}
			case "tweak":
				choice := front.InputMenuGen("Action:", []string{"Change device ID", "Change master password/Optimize entries", "Set text editor", "Age all entries"})
				switch choice {
				case 1:
					oldDeviceID, err := global.GetCurrentDeviceID()
					if err != nil {
						other.PrintError("Failed to get current device ID: "+err.Error(), back.ErrorRead)
					}
					_, _, _, err = syncclient.GenDeviceID(oldDeviceID, "")
					if err != nil {
						other.PrintError("Failed to change device ID: "+err.Error(), global.ErrorSyncProcess)
					}
					fmt.Println("\nDevice ID changed successfully.")
				case 2:
					oldPassword := confirmRCWPassword("old")
					newPassword := confirmRCWPassword("new")
					fmt.Print("\nRe-encrypting entries. Please wait; do not force close this process.\n")
					err := core.EntryRefresh(oldPassword, newPassword, false)
					if err != nil {
						other.PrintError("Re-encryption failed: "+err.Error(), global.ErrorEncryption)
					}
					fmt.Println("\nRe-encryption complete.")
				case 3:
					newCfg := &cfg.ConfigT{}
					newThirdPartyCfg := map[string]any{"mutnTextEditor": getTextEditorInput()}
					newCfg.ThirdParty = &newThirdPartyCfg
					err := cfg.WriteConfig(newCfg, true)
					if err != nil {
						other.PrintError("Failed to set text editor: "+err.Error(), back.ErrorWrite)
					}
				case 4:
					forceReage := front.InputBinary("Re-age aged entries?")
					fmt.Println("Aging entries; this may take awhile - do not terminate this process")
					err := age.AllPasswordEntries(forceReage)
					if err != nil {
						other.PrintError("Failed to age entries: "+err.Error(), 1)
					}
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

// confirmRCWPassword prompts the user for a master password and confirms it
// is typed correctly. Returns the password as a byte slice.
func confirmRCWPassword(lowercasePrefix string) []byte {
	for {
		// get master password
		rcwPass := front.InputHidden(strings.ToUpper(string(lowercasePrefix[0])) + lowercasePrefix[1:] + " master password:")
		if !bytes.Equal(rcwPass, front.InputHidden("Confirm "+lowercasePrefix+" master password:")) || len(rcwPass) == 0 {
			fmt.Println(back.AnsiError + "Passwords do not match or password is invalid" + back.AnsiReset)
			continue
		}
		return rcwPass
	}
}

// getTextEditorInput prompts the user for a text editor input, first falling back to "$EDITOR", then to a default value
func getTextEditorInput() string {
	return cmp.Or(front.Input("Text editor (leave blank for $EDITOR, falls back to \""+cli.FallbackEditor+"\"):"), os.Getenv("EDITOR"), cli.FallbackEditor)
}
