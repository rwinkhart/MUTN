package cli

import (
	"os"

	"golang.org/x/term"
)

// global variables used across multiple files
var (
	width, _, _ = term.GetSize(int(os.Stdout.Fd()))
)

// global constants used across multiple files
const (
	AnsiBold         = "\033[1m"
	ansiBlackOnWhite = "\033[38;5;0;48;5;15m"
)
