package cli

import (
	"fmt"
	"os"

	"github.com/rwinkhart/MUTN/src/backend"
)

// global constants used only in this file
const (
	ansiVersionMeat    = "\033[38;2;157;0;6m"
	ansiVersionOutline = "\033[38;2;131;165;152m"
)

func HelpMain() {
	fmt.Print(AnsiBold + "\nMUTN | Copyright (c) 2024 Randall Winkhart\n" + backend.AnsiReset + `
This software exists under the MIT license; you may redistribute it under certain conditions.
This program comes with absolutely no warranty; type "mutn version" for details.

` + AnsiBold + "Usage:" + backend.AnsiReset + ` mutn [/<entry name> [argument] [option]] | [argument]

` + AnsiBold + "Arguments:" + backend.AnsiReset + `
 help                    Bring up this menu
 version                 Display version and license information
 init                    Set up MUTN (generates libmutton.ini)
 tweak                   Change configuration options
 copy                    Copy details of an entry to your clipboard
 edit                    Edit an existing entry
 gen                     Generate a new password
 add                     Add an entry
 shear                   Delete an existing entry
 sync                    Manually sync the entry directory

` + AnsiBold + "Options:" + backend.AnsiReset + `
 copy:
  password|-pw|<blank>   Copy the password of an entry to your clipboard
  username|-u            Copy the username of an entry to your clipboard
  totp|-t                Copy the TOTP code of an entry to your clipboard
  url|-l                 Copy the URL of an entry to your clipboard
  note|-n                Copy the note of an entry to your clipboard
 edit:
  password|-pw|<blank>   Change the password of an entry
  username|-u            Change the username of an entry
  totp|-t                Change the TOTP secret of an entry
  url|-l                 Change the URL attached to an entry
  note|-n                Change the note attached to an entry
  rename|-r              Rename or relocate an entry
 gen:
  update|-u              Generate a password for an existing entry
 add:
  password|-pw|<blank>   Add a password entry
  note|-n                Add a note entry
  folder|-f              Add a new folder for entries

` + AnsiBold + "Tip 1:" + backend.AnsiReset + ` You can quickly read an entry with "mutn /<entry name>"
` + AnsiBold + "Tip 2:" + backend.AnsiReset + ` Type "mutn" (no arguments/options) to view a list of saved entries
` + AnsiBold + "Tip 3:" + backend.AnsiReset + ` Provide "add", "edit", "copy", or "gen" as the only argument to receive more specific help
` + AnsiBold + "Tip 4:" + backend.AnsiReset + " Using \"add\", \"edit\", or \"copy\" without specifying an option (field) will default to \"password\"\n\n")
	os.Exit(0)
}

func HelpAdd() {
	fmt.Print(AnsiBold + "\nUsage:" + backend.AnsiReset + ` mutn /<entry name> add <option>

` + AnsiBold + "Options:" + backend.AnsiReset + `
 add:
  password|-pw|<blank>   Add a password entry
  note|-n                Add a note entry
  folder|-f              Add a new folder for entries` + "\n\n")
	os.Exit(0)
}

func HelpEdit() {
	fmt.Print(AnsiBold + "\nUsage:" + backend.AnsiReset + ` mutn /<entry name> edit <option>

` + AnsiBold + "Options:" + backend.AnsiReset + `
 edit:
  password|-pw|<blank>   Change the password of an entry
  username|-u            Change the username of an entry
  totp|-t                Change the TOTP secret of an entry
  url|-l                 Change the URL attached to an entry
  note|-n                Change the note attached to an entry
  rename|-r              Rename or relocate an entry` + "\n\n")
	os.Exit(0)
}

func HelpCopy() {
	fmt.Print(AnsiBold + "\nUsage:" + backend.AnsiReset + ` mutn /<entry name> copy <option>

` + AnsiBold + "Options:" + backend.AnsiReset + `
 copy:
  password|-pw|<blank>   Copy the password in an entry to your clipboard
  username|-u            Copy the username in an entry to your clipboard
  totp|-t                Copy the TOTP code of an entry to your clipboard
  url|-l                 Copy the URL in an entry to your clipboard
  note|-n                Copy the first note line in an entry to your clipboard` + "\n\n")
	os.Exit(0)
}

func HelpGen() {
	fmt.Print(AnsiBold + "\nUsage:" + backend.AnsiReset + ` mutn /<entry name> gen [option]

` + AnsiBold + "Options:" + backend.AnsiReset + `
 gen:
  update|-u              Generate a password for an existing entry

` + AnsiBold + "Tip:" + backend.AnsiReset + " If no options are provided, a new password entry is generated\n\n")
	os.Exit(0)
}

func MITLicense() {
	fmt.Print(AnsiBold + "\n                    MIT License" + backend.AnsiReset + `

  Permission is hereby granted, free of charge, to any
person obtaining a copy of this software and associated
  documentation files (the "Software"), to deal in the
    Software without restriction, including without
   limitation the rights to use, copy, modify, merge,
 publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software
    is furnished to do so, subject to the following
                      conditions:

 The above copyright notice and this permission notice
shall be included in all copies or substantial portions
                   of the Software.

 THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF
ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED
  TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A
  PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT
 SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR
 ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
 ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE
           OR OTHER DEALINGS IN THE SOFTWARE.` + "\n\n---------------------------------------------------------")
	// do not exit, as this is meant to be followed by other information
}

func Version() {
	MITLicense()
	fmt.Print("\n\n             MUTN is a simple, self-hosted,\n  SSH-synchronized password manager based on libmutton\n\n" +
		"         ..                                     ..\n" +
		"        /()\\''.''.    " + ansiVersionMeat + "♥♥♥♥" + backend.AnsiReset + "               .''.''/()\\   _)\n" +
		"     _.   :       *  " + ansiVersionMeat + "♥♥♥♥♥♥   ♥♥♥♥♥♥♥♥" + backend.AnsiReset + "  *       :   <[◎]|_|=\n" +
		" }-}-*]    `..'..'    " + ansiVersionMeat + "♥♥♥♥♥♥♥♥♥♥♥♥♥" + backend.AnsiReset + "      `..'..'      |\n" +
		"    ◎-◎    //   \\\\     " + ansiVersionMeat + "♥♥♥♥♥♥♥♥♥" + backend.AnsiReset + "         //   \\\\     /|\\\n" +
		ansiVersionOutline + "<><><><><><><><><><><><><><>-<><><><><><><><><><><><><><>\n" +
		"\\" + ansiBlackOnWhite + "                                                       " + backend.AnsiReset + ansiVersionOutline + "/\n" +
		"\\" + ansiBlackOnWhite + "                  MUTN Version " + backend.LibmuttonVersion + "                   " + backend.AnsiReset + ansiVersionOutline + "/\n" +
		"\\" + ansiBlackOnWhite + "                The Placeholder Update                 " + backend.AnsiReset + ansiVersionOutline + "/\n" +
		"\\" + ansiBlackOnWhite + "                                                       " + backend.AnsiReset + ansiVersionOutline + "/\n" +
		"\\" + ansiBlackOnWhite + "          Copyright (c) 2024 Randall Winkhart          " + backend.AnsiReset + ansiVersionOutline + "/\n" +
		"\\" + ansiBlackOnWhite + "                                                       " + backend.AnsiReset + ansiVersionOutline + "/\n" +
		"<><><><><><><><><><><><><><>-<><><><><><><><><><><><><><>\n" + backend.AnsiReset +
		"\n               For more information, see:\n\n" +
		"           https://github.com/rwinkhart/MUTN\n" +
		"         https://github.com/rwinkhart/libmutton\n\n")
	os.Exit(0)
}
