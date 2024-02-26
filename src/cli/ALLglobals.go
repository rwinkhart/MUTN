package cli

import (
	"github.com/rwinkhart/MUTN/src/offline"
	termy "golang.org/x/crypto/ssh/terminal"
	"os"
)

// global variables used across multiple files
var (
	rootLength  = len(offline.EntryRoot)
	width, _, _ = termy.GetSize(int(os.Stdout.Fd()))
)

// global constants used across multiple files
const (
	ansiReset        = "\033[0m"
	ansiBold         = "\033[1m"
	ansiBlackOnWhite = "\033[38;5;0;48;5;15m"
)
