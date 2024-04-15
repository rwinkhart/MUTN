package backend

import (
	"os"
)

// global variables used across multiple files
var (
	home, _ = os.UserHomeDir()
)

// global constants used across multiple files
const (
	AnsiError = "\033[38;5;9m"
	AnsiReset = "\033[0m"
)
