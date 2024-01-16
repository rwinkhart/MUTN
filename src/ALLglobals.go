package main

import (
	termy "golang.org/x/crypto/ssh/terminal"
	"os"
)

// invisible variables
var home, _ = os.UserHomeDir()

// exported variables
var (
	RootLength  = len(EntryRoot)
	Width, _, _ = termy.GetSize(int(os.Stdout.Fd()))
)
