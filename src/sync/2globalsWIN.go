//go:build windows

package sync

import (
	"os"
	"strings"
)

var vanityEXE = os.Args[0][strings.LastIndex(os.Args[0], "\\")+1:]
