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
	AnsiError        = "\033[38;5;9m"
	AnsiReset        = "\033[0m"
	LibmuttonVersion = "0.2.A" // untagged releases feature a letter suffix corresponding to the eventual release version, e.g "0.2.A" -> "0.2.0", "0.2.B" -> "0.2.1"
)
