//go:build windows

package backend

import (
	"os"
	"syscall"
)

// EntryRoot path to libmutton entry directory
var EntryRoot = Home + "\\AppData\\Local\\libmutton\\entries"
var ConfigDir = Home + "\\AppData\\Local\\libmutton\\config"
var ConfigPath = ConfigDir + "\\libmutton.ini"

// PathSeparator defines the character used to separate directories in a path (platform-specific)
const (
	PathSeparator = "\\"
	IsWindows     = true
)

// enableVirtualTerminalProcessing ensures ANSI escape sequences are interpreted properly on Windows
// TODO remove after migration off of GPG, as pinentry is responsible for disabling ANSI escape sequence interpretation
func enableVirtualTerminalProcessing() {
	stdout := syscall.Handle(os.Stdout.Fd())

	var originalMode uint32
	syscall.GetConsoleMode(stdout, &originalMode)
	originalMode |= 0x0004

	syscall.MustLoadDLL("kernel32").MustFindProc("SetConsoleMode").Call(uintptr(stdout), uintptr(originalMode))
}
