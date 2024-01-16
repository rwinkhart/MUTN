package main

import (
	"os"
)

func main() {

	args := os.Args[1:]
	argsCount := len(args)

	if argsCount == 0 {

		EntryListGen()

	} else if argsCount == 1 {

		switch args[0] {

		case "help", "--help", "-h":
			HelpMain()
		case "version", "-v":
			Version()
		default:
			HelpMain()

		}
	}
}
