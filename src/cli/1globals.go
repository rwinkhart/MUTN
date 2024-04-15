package cli

import (
	"github.com/rwinkhart/MUTN/src/backend"
	termy "golang.org/x/crypto/ssh/terminal"
	"os"
)

// global variables used across multiple files
var (
	rootLength  = len(backend.EntryRoot)
	width, _, _ = termy.GetSize(int(os.Stdout.Fd()))
)

// global constants used across multiple files
const (
	ansiBold         = "\033[1m"
	ansiBlackOnWhite = "\033[38;5;0;48;5;15m"
)
