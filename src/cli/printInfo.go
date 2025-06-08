package cli

import (
	"fmt"
	"os"

	"github.com/rwinkhart/go-boilerplate/back"
	"github.com/rwinkhart/libmutton/global"
)

// ANSI color constants used only in this file
const (
	ansiVersionMeat    = "\033[38;2;157;0;6m"
	ansiVersionOutline = "\033[38;2;131;165;152m"
)

func HelpMain() {
	fmt.Print(back.AnsiBold + "\nMUTN | Copyright (c) 2024-2025 Randall Winkhart\n" + back.AnsiReset + `
This software exists under the MIT license; you may redistribute it under certain conditions.
This program comes with absolutely no warranty; type "mutn version" for details.

` + back.AnsiBold + "Usage:" + back.AnsiReset + ` mutn [/<entry name> [argument] [option]] | [argument]

` + back.AnsiBold + "Arguments:" + back.AnsiReset + `
 help                    Bring up this menu
 version                 Display version and license information
 init                    Set up MUTN (generates libmutton.ini)
 tweak                   Make changes to the libmutton configuration
 copy                    Copy details of an entry to your clipboard
 edit                    Edit an existing entry
 gen                     Generate a new password
 add                     Add a new entry
 shear                   Delete an existing entry
 sync                    Manually sync the entry directory

` + back.AnsiBold + "Options:" + back.AnsiReset + `
 copy:
  password|-pw|<blank>   Copy the password in an entry to your clipboard
  username|-u            Copy the username in an entry to your clipboard
  totp|-t                Generate and copy the TOTP token for an entry to your clipboard
  url|-l                 Copy the URL in an entry to your clipboard
  note|-n                Copy the first note line in an entry to your clipboard
 edit:
  password|-pw|<blank>   Change the password in an entry
  username|-u            Change the username in an entry
  totp|-t                Change the TOTP secret in an entry
  url|-l                 Change the URL in an entry
  note|-n                Change the note in an entry
  rename|-r              Rename or relocate an entry
 gen:
  update|-u              Generate a password for an existing entry
 add:
  password|-pw|<blank>   Add a password entry
  note|-n                Add a note entry
  folder|-f              Add a new folder for entries

` + back.AnsiBold + "Tip 1:" + back.AnsiReset + ` You can quickly read an entry with "mutn /<entry name>"
` + back.AnsiBold + "Tip 2:" + back.AnsiReset + ` Type "mutn" (no arguments/options) to view a list of saved entries
` + back.AnsiBold + "Tip 3:" + back.AnsiReset + ` Provide "add", "edit", "copy", or "gen" as the only argument to receive more specific help
` + back.AnsiBold + "Tip 4:" + back.AnsiReset + " Using \"add\", \"edit\", or \"copy\" without specifying an option (field) will default to \"password\"\n\n")
	os.Exit(0)
}

func HelpAdd() {
	fmt.Print(back.AnsiBold + "\nUsage:" + back.AnsiReset + ` mutn /<entry name> add <option>

` + back.AnsiBold + "Options:" + back.AnsiReset + `
 add:
  password|-pw|<blank>   Add a password entry
  note|-n                Add a note entry
  folder|-f              Add a new folder for entries` + "\n\n")
	os.Exit(0)
}

func HelpEdit() {
	fmt.Print(back.AnsiBold + "\nUsage:" + back.AnsiReset + ` mutn /<entry name> edit <option>

` + back.AnsiBold + "Options:" + back.AnsiReset + `
 edit:
  password|-pw|<blank>   Change the password in an entry
  username|-u            Change the username in an entry
  totp|-t                Change the TOTP secret in an entry
  url|-l                 Change the URL in an entry
  note|-n                Change the note in an entry
  rename|-r              Rename or relocate an entry` + "\n\n")
	os.Exit(0)
}

func HelpCopy() {
	fmt.Print(back.AnsiBold + "\nUsage:" + back.AnsiReset + ` mutn /<entry name> copy <option>

` + back.AnsiBold + "Options:" + back.AnsiReset + `
 copy:
  password|-pw|<blank>   Copy the password in an entry to your clipboard
  username|-u            Copy the username in an entry to your clipboard
  totp|-t                Generate and copy the TOTP token for an entry to your clipboard
  url|-l                 Copy the URL in an entry to your clipboard
  note|-n                Copy the first note line in an entry to your clipboard` + "\n\n")
	os.Exit(0)
}

func HelpGen() {
	fmt.Print(back.AnsiBold + "\nUsage:" + back.AnsiReset + ` mutn /<entry name> gen [option]

` + back.AnsiBold + "Options:" + back.AnsiReset + `
 gen:
  update|-u              Generate a password for an existing entry

` + back.AnsiBold + "Tip:" + back.AnsiReset + " If no options are provided, a new password entry is generated\n\n")
	os.Exit(0)
}

func MITLicense() {
	fmt.Print(back.AnsiBold + "\n                    MIT License" + back.AnsiReset + `

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
		"        /()\\''.''.    " + ansiVersionMeat + "♥♥♥♥" + back.AnsiReset + "               .''.''/()\\   _)\n" +
		"     _.   :       *  " + ansiVersionMeat + "♥♥♥♥♥♥   ♥♥♥♥♥♥♥♥" + back.AnsiReset + "  *       :   <[◎]|_|=\n" +
		" }-}-*]    `..'..'    " + ansiVersionMeat + "♥♥♥♥♥♥♥♥♥♥♥♥♥" + back.AnsiReset + "      `..'..'      |\n" +
		"    ◎-◎    //   \\\\     " + ansiVersionMeat + "♥♥♥♥♥♥♥♥♥" + back.AnsiReset + "         //   \\\\     /|\\\n" +
		ansiVersionOutline + "<><><><><><><><><><><><><><>-<><><><><><><><><><><><><><>\n" +
		"\\" + ansiBlackOnWhite + "                                                       " + back.AnsiReset + ansiVersionOutline + "/\n" +
		"\\" + ansiBlackOnWhite + "                      MUTN v" + MUTNVersion + "                      " + back.AnsiReset + ansiVersionOutline + "/\n" +
		"\\" + ansiBlackOnWhite + "                The Cryptic Cuts Update                " + back.AnsiReset + ansiVersionOutline + "/\n" +
		"\\" + ansiBlackOnWhite + "                                                       " + back.AnsiReset + ansiVersionOutline + "/\n" +
		"\\" + ansiBlackOnWhite + "              Built with libmutton v" + global.LibmuttonVersion + "              " + back.AnsiReset + ansiVersionOutline + "/\n" +
		"\\" + ansiBlackOnWhite + "                                                       " + back.AnsiReset + ansiVersionOutline + "/\n" +
		"\\" + ansiBlackOnWhite + "       Copyright (c) 2024-2025: Randall Winkhart       " + back.AnsiReset + ansiVersionOutline + "/\n" +
		"\\" + ansiBlackOnWhite + "                                                       " + back.AnsiReset + ansiVersionOutline + "/\n" +
		"<><><><><><><><><><><><><><>-<><><><><><><><><><><><><><>\n" + back.AnsiReset +
		"\n               For more information, see:\n\n" +
		"           https://github.com/rwinkhart/MUTN\n" +
		"         https://github.com/rwinkhart/libmutton\n\n")
	os.Exit(0)
}
