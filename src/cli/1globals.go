package cli

import (
	"os"

	"golang.org/x/term"
)

var (
	width, _, _ = term.GetSize(int(os.Stdout.Fd()))
)

const (
	AnsiBold         = "\033[1m"
	ansiBlackOnWhite = "\033[38;5;0;48;5;15m"
	MUTNVersion      = "0.2.B" // untagged releases feature a letter suffix corresponding to the eventual release version, e.g "0.2.A" -> "0.2.0", "0.2.B" -> "0.2.1"
)
