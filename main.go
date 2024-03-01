package main

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/cli"
	"github.com/rwinkhart/MUTN/src/offline"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	argsCount := len(args)

	if argsCount == 0 {

		cli.EntryListGen()

	} else if argsCount == 1 {

		// initial entry reader shortcut
		if strings.HasPrefix(args[0], "/") {
			targetLocation := offline.EntryRoot + offline.PathSeparator + args[0][1:]
			if isFile, _ := offline.TargetIsFile(targetLocation, true); isFile {
				// TODO decrypt and display the entry
				fmt.Println("It's a file!")
			} else {
				fmt.Println(offline.AnsiError + "Failed to read \"" + targetLocation + "\" - it is a directory" + offline.AnsiReset)
			}
			os.Exit(0)
		}

		switch args[0] {

		case "help", "--help", "-h":
			cli.HelpMain()
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
	} else if argsCount > 1 {
		// load config (libmutton.ini)
		offline.ReadConfig()
	}
}
