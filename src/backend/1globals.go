package backend

import (
	"os"
)

// global variables used across multiple files
var (
	Home, _ = os.UserHomeDir()
)

// global constants used across multiple files
const (
	AnsiError = "\033[38;5;9m"
	AnsiReset = "\033[0m"
)
