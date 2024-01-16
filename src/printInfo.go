package main

import "fmt"

func HelpMain() {
	fmt.Print("\n\033[1mMUTN | Copyright (c) 2024 Randall Winkhart\033[0m\n" + `
This software exists under the MIT license; you may redistribute it under certain conditions.
This program comes with absolutely no warranty; type "mutn version" for details.

` + "\033[1mUsage:\033[0m" + ` mutn [/<entry name> [argument] [option]] | [argument]

` + "\033[1mArguments:\033[0m" + `
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

` + "\033[1mOptions:\033[0m" + `
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

` + "\033[1mTip 1:\033[0m" + ` You can quickly read an entry with "mutn /<entry name>"
` + "\033[1mTip 2:\033[0m" + ` Type "mutn" to view a list of saved entries` + "\n\n")
}

func Version() {
	fmt.Print("\n\033[1m                    MIT License\033[0m" + `

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
		"         ..               \033[38;2;206;48;98m♥♥ ♥♥\033[0m               ..\n" +
		"        /()\\''.''.       \033[38;2;206;48;98m♥♥♥♥♥♥♥\033[0m       .''.''/()\\   _)\n" +
		"     _.   :       *       \033[38;2;206;48;98m♥♥♥♥♥\033[0m       *       :   <[◎]|_|=\n" +
		" }-}-*]    `..'..'         \033[38;2;206;48;98m♥♥♥\033[0m         `..'..'      |\n" +
		"    ◎-◎    //   \\\\          \033[38;2;206;48;98m♥\033[0m          //   \\\\     /|\\\n" +
		"\033[38;2;1;173;216m<><><><><><><><><><><><><><>-<><><><><><><><><><><><><><>\n" +
		"\\\033[38;5;15;48;5;15m                                                       \033[0m\033[38;2;1;173;216m/\n" +
		"\\\033[38;5;15;48;5;15m                  \033[0mMUTN Version 0.0.1\033[38;5;15;48;5;15m                   \033[0m\033[38;2;1;173;216m/\n" +
		"\\\033[38;5;15;48;5;15m                 \033[0mThe Butchered Update\033[38;5;15;48;5;15m                  \033[0m\033[38;2;1;173;216m/\n" +
		"\\\033[38;5;15;48;5;15m                                                       \033[0m\033[38;2;1;173;216m/\n" +
		"\\\033[38;5;15;48;5;15m          \033[0mCopyright (c) 2024 Randall Winkhart\033[38;5;15;48;5;15m          \033[0m\033[38;2;1;173;216m/\n" +
		"\\\033[38;5;15;48;5;15m                                                       \033[0m\033[38;2;1;173;216m/\n" +
		"<><><><><><><><><><><><><><>-<><><><><><><><><><><><><><>\033[0m\n" +
		"\nFor more information, see:\n" +
		"https://github.com/rwinkhart/MUTN\n" +
		"https://github.com/rwinkhart/libmutton\n\n")
}
