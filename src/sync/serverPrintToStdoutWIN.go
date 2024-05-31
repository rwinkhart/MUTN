//go:build windows

package sync

import (
	"fmt"
	"strings"
)

// printToStdout prints a string to stdout with UNIX path separators
func printToStdout(targetLocationIncomplete string) {
	fmt.Print("\x1f" + strings.ReplaceAll(targetLocationIncomplete, "\\", "/"))
}
