//go:build !windows

package sync

import (
	"fmt"
)

// printToStdout prints a string to stdout with UNIX path separators
func printToStdout(targetLocationIncomplete string) {
	fmt.Print("\x1f" + targetLocationIncomplete)
}
