//go:build !windows

package sync

import (
	"os"
)

var vanityEXE = os.Args[0]
