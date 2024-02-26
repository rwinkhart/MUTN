package main

import (
	"github.com/rwinkhart/MUTN/src/cli"
	"os"
)

func main() {

	args := os.Args[1:]
	argsCount := len(args)

	if argsCount == 0 {

		cli.EntryListGen()

	} else if argsCount == 1 {

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
	}
}
