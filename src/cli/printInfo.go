package cli

import "fmt"

// global constants used only in this file
const (
	ansiGoFuchsia = "\033[38;2;206;48;98m"
	ansiGoGopher  = "\033[38;2;1;173;216m"
)

func HelpMain() {
	fmt.Print(ansiBold + "\nMUTN | Copyright (c) 2024 Randall Winkhart\n" + ansiReset + `
This software exists under the MIT license; you may redistribute it under certain conditions.
This program comes with absolutely no warranty; type "mutn version" for details.

` + ansiBold + "Usage:" + ansiReset + ` mutn [/<entry name> [argument] [option]] | [argument]

` + ansiBold + "Arguments:" + ansiReset + `
 help/--help/-h          Bring up this menu
 version/-v              Display version and license information
 init                    Set up MUTN
 tweak                   Change configuration options
 add                     Add an entry
 gen                     Generate a new password
 edit                    Edit an existing entry
 copy                    Copy details of an entry to your clipboard
 shear                   Delete an existing entry
 sync                    Manually sync the entry directory

` + ansiBold + "Options:" + ansiReset + `
 add:
  password/-p            Add a password entry
  note/-n                Add a note entry
  folder/-f              Add a new folder for entries
 edit:
  rename/relocate/-r     Rename or relocate an entry
  username/-u            Change the username of an entry
  password/-p            Change the password of an entry
  url/-l                 Change the url attached to an entry
  note/-n                Change the note attached to an entry
 copy:
  username/-u            Copy the username of an entry to your clipboard
  password/-p            Copy the password of an entry to your clipboard
  url/-l                 Copy the url of an entry to your clipboard
  note/-n                Copy the note of an entry to your clipboard
 gen:
  update/-u              Generate a password for an existing entry

` + ansiBold + "Tip 1:" + ansiReset + ` You can quickly read an entry with "mutn /<entry name>"
` + ansiBold + "Tip 2:" + ansiReset + ` Type "mutn" (no arguments/options) to view a list of saved entries
` + ansiBold + "Tip 3:" + ansiReset + " Provide \"add\", \"edit\", \"copy\", or \"gen\" as the only argument to receive more specific help\n\n")
}

func HelpAdd() {
	fmt.Print(ansiBold + "\nUsage:" + ansiReset + ` mutn /<entry name> add <option>

` + ansiBold + "Options:" + ansiReset + `
 add:
  password/-p            Add a password entry
  note/-n                Add a note entry
  folder/-f              Add a new folder for entries` + "\n\n")
}

func HelpEdit() {
	fmt.Print(ansiBold + "\nUsage:" + ansiReset + ` mutn /<entry name> edit <option>

` + ansiBold + "Options:" + ansiReset + `
 edit:
  rename/-r              Rename or relocate an entry
  username/-u            Change the username of an entry
  password/-p            Change the password of an entry
  url/-l                 Change the url attached to an entry
  note/-n                Change the note attached to an entry` + "\n\n")
}

func HelpCopy() {
	fmt.Print(ansiBold + "\nUsage:" + ansiReset + ` mutn /<entry name> copy <option>

` + ansiBold + "Options:" + ansiReset + `
 copy:
  username/-u            Copy the username in an entry to your clipboard
  password/-p            Copy the password in an entry to your clipboard
  url/-l                 Copy the url in an entry to your clipboard
  note/-n                Copy the first note line in an entry to your clipboard` + "\n\n")
}

func HelpGen() {
	fmt.Print(ansiBold + "\nUsage:" + ansiReset + ` mutn /<entry name> gen [option]

` + ansiBold + "Options:" + ansiReset + `
 gen:
  update/-u              Generate a password for an existing entry

` + ansiBold + "Tip:" + ansiReset + " If no options are provided, a new password entry is generated\n\n")
}

func Version() {
	fmt.Print(ansiBold + "\n                    MIT License" + ansiReset + `

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
		"        /()\\''.''.    " + ansiGoFuchsia + "♥♥♥♥" + ansiReset + "               .''.''/()\\   _)\n" +
		"     _.   :       *  " + ansiGoFuchsia + "♥♥♥♥♥♥   ♥♥♥♥♥♥♥♥" + ansiReset + "  *       :   <[◎]|_|=\n" +
		" }-}-*]    `..'..'    " + ansiGoFuchsia + "♥♥♥♥♥♥♥♥♥♥♥♥♥" + ansiReset + "      `..'..'      |\n" +
		"    ◎-◎    //   \\\\     " + ansiGoFuchsia + "♥♥♥♥♥♥♥♥♥" + ansiReset + "         //   \\\\     /|\\\n" +
		ansiGoGopher + "<><><><><><><><><><><><><><>-<><><><><><><><><><><><><><>\n" +
		"\\" + ansiBlackOnWhite + "                                                       " + ansiReset + ansiGoGopher + "/\n" +
		"\\" + ansiBlackOnWhite + "                  MUTN Version 0.0.1                   " + ansiReset + ansiGoGopher + "/\n" +
		"\\" + ansiBlackOnWhite + "                 The Butchered Update                  " + ansiReset + ansiGoGopher + "/\n" +
		"\\" + ansiBlackOnWhite + "                                                       " + ansiReset + ansiGoGopher + "/\n" +
		"\\" + ansiBlackOnWhite + "          Copyright (c) 2024 Randall Winkhart          " + ansiReset + ansiGoGopher + "/\n" +
		"\\" + ansiBlackOnWhite + "                                                       " + ansiReset + ansiGoGopher + "/\n" +
		"<><><><><><><><><><><><><><>-<><><><><><><><><><><><><><>\n" + ansiReset +
		"\n               For more information, see:\n\n" +
		"           https://github.com/rwinkhart/MUTN\n" +
		"         https://github.com/rwinkhart/libmutton\n\n")
}
