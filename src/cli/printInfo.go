package cli

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/offline"
	"os"
)

// global constants used only in this file
const (
	ansiGoFuchsia = "\033[38;2;206;48;98m"
	ansiGoGopher  = "\033[38;2;1;173;216m"
)

func HelpMain() {
	fmt.Print(ansiBold + "\nMUTN | Copyright (c) 2024 Randall Winkhart\n" + offline.AnsiReset + `
This software exists under the MIT license; you may redistribute it under certain conditions.
This program comes with absolutely no warranty; type "mutn version" for details.

` + ansiBold + "Usage:" + offline.AnsiReset + ` mutn [/<entry name> [argument] [option]] | [argument]

` + ansiBold + "Arguments:" + offline.AnsiReset + `
 help|` + OptionFlag + `h                 Bring up this menu
 version|` + OptionFlag + `v              Display version and license information
 init                    Set up MUTN (generates libmutton.ini)
 tweak                   Change configuration options
 add                     Add an entry
 gen                     Generate a new password
 edit                    Edit an existing entry
 copy                    Copy details of an entry to your clipboard
 shear                   Delete an existing entry
 sync                    Manually sync the entry directory

` + ansiBold + "Options:" + offline.AnsiReset + `
 add:
  password|` + OptionFlag + `p            Add a password entry
  note|` + OptionFlag + `n                Add a note entry
  folder|` + OptionFlag + `f              Add a new folder for entries
 edit:
  rename|` + OptionFlag + `r              Rename or relocate an entry
  username|` + OptionFlag + `u            Change the username of an entry
  password|` + OptionFlag + `p            Change the password of an entry
  url|` + OptionFlag + `l                 Change the url attached to an entry
  note|` + OptionFlag + `n                Change the note attached to an entry
 copy:
  username|` + OptionFlag + `u            Copy the username of an entry to your clipboard
  password|` + OptionFlag + `p            Copy the password of an entry to your clipboard
  url|` + OptionFlag + `l                 Copy the url of an entry to your clipboard
  note|` + OptionFlag + `n                Copy the note of an entry to your clipboard
 gen:
  update|` + OptionFlag + `u              Generate a password for an existing entry

` + ansiBold + "Tip 1:" + offline.AnsiReset + ` You can quickly read an entry with "mutn /<entry name>"
` + ansiBold + "Tip 2:" + offline.AnsiReset + ` Type "mutn" (no arguments/options) to view a list of saved entries
` + ansiBold + "Tip 3:" + offline.AnsiReset + " Provide \"add\", \"edit\", \"copy\", or \"gen\" as the only argument to receive more specific help\n\n")
	os.Exit(0)
}

func HelpAdd() {
	fmt.Print(ansiBold + "\nUsage:" + offline.AnsiReset + ` mutn /<entry name> add <option>

` + ansiBold + "Options:" + offline.AnsiReset + `
 add:
  password|` + OptionFlag + `p            Add a password entry
  note|` + OptionFlag + `n                Add a note entry
  folder|` + OptionFlag + `f              Add a new folder for entries` + "\n\n")
	os.Exit(0)
}

func HelpEdit() {
	fmt.Print(ansiBold + "\nUsage:" + offline.AnsiReset + ` mutn /<entry name> edit <option>

` + ansiBold + "Options:" + offline.AnsiReset + `
 edit:
  rename|` + OptionFlag + `r              Rename or relocate an entry
  username|` + OptionFlag + `u            Change the username of an entry
  password|` + OptionFlag + `p            Change the password of an entry
  url|` + OptionFlag + `l                 Change the url attached to an entry
  note|` + OptionFlag + `n                Change the note attached to an entry` + "\n\n")
	os.Exit(0)
}

func HelpCopy() {
	fmt.Print(ansiBold + "\nUsage:" + offline.AnsiReset + ` mutn /<entry name> copy <option>

` + ansiBold + "Options:" + offline.AnsiReset + `
 copy:
  username|` + OptionFlag + `u            Copy the username in an entry to your clipboard
  password|` + OptionFlag + `p            Copy the password in an entry to your clipboard
  url|` + OptionFlag + `l                 Copy the url in an entry to your clipboard
  note|` + OptionFlag + `n                Copy the first note line in an entry to your clipboard` + "\n\n")
	os.Exit(0)
}

func HelpGen() {
	fmt.Print(ansiBold + "\nUsage:" + offline.AnsiReset + ` mutn /<entry name> gen [option]

` + ansiBold + "Options:" + offline.AnsiReset + `
 gen:
  update|` + OptionFlag + `u              Generate a password for an existing entry

` + ansiBold + "Tip:" + offline.AnsiReset + " If no options are provided, a new password entry is generated\n\n")
	os.Exit(0)
}

func Version() {
	fmt.Print(ansiBold + "\n                    MIT License" + offline.AnsiReset + `

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
           OR OTHER DEALINGS IN THE SOFTWARE.` +
		"\n\n---------------------------------------------------------" +
		"\n\n             MUTN is a simple, self-hosted,\n  SSH-synchronized password manager based on libmutton\n\n" +
		"         ..                                     ..\n" +
		"        /()\\''.''.    " + ansiGoFuchsia + "♥♥♥♥" + offline.AnsiReset + "               .''.''/()\\   _)\n" +
		"     _.   :       *  " + ansiGoFuchsia + "♥♥♥♥♥♥   ♥♥♥♥♥♥♥♥" + offline.AnsiReset + "  *       :   <[◎]|_|=\n" +
		" }-}-*]    `..'..'    " + ansiGoFuchsia + "♥♥♥♥♥♥♥♥♥♥♥♥♥" + offline.AnsiReset + "      `..'..'      |\n" +
		"    ◎-◎    //   \\\\     " + ansiGoFuchsia + "♥♥♥♥♥♥♥♥♥" + offline.AnsiReset + "         //   \\\\     /|\\\n" +
		ansiGoGopher + "<><><><><><><><><><><><><><>-<><><><><><><><><><><><><><>\n" +
		"\\" + ansiBlackOnWhite + "                                                       " + offline.AnsiReset + ansiGoGopher + "/\n" +
		"\\" + ansiBlackOnWhite + "                  MUTN Version 0.0.1                   " + offline.AnsiReset + ansiGoGopher + "/\n" +
		"\\" + ansiBlackOnWhite + "                 The Butchered Update                  " + offline.AnsiReset + ansiGoGopher + "/\n" +
		"\\" + ansiBlackOnWhite + "                                                       " + offline.AnsiReset + ansiGoGopher + "/\n" +
		"\\" + ansiBlackOnWhite + "          Copyright (c) 2024 Randall Winkhart          " + offline.AnsiReset + ansiGoGopher + "/\n" +
		"\\" + ansiBlackOnWhite + "                                                       " + offline.AnsiReset + ansiGoGopher + "/\n" +
		"<><><><><><><><><><><><><><>-<><><><><><><><><><><><><><>\n" + offline.AnsiReset +
		"\n               For more information, see:\n\n" +
		"           https://github.com/rwinkhart/MUTN\n" +
		"         https://github.com/rwinkhart/libmutton\n\n")
	os.Exit(0)
}
