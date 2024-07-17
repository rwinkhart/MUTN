//go:build !windows

package backend

// EntryRoot path to libmutton entry directory
var EntryRoot = Home + "/.local/share/libmutton"
var ConfigDir = Home + "/.config/libmutton"
var ConfigPath = ConfigDir + "/libmutton.ini"

// PathSeparator defines the character used to separate directories in a path (platform-specific)
const (
	PathSeparator = "/"
	IsWindows     = false
)

// enableVirtualTerminalProcessing is a dummy function on UNIX-like systems (only needed on Windows)
// TODO remove after migration off of GPG, as pinentry is responsible for disabling ANSI escape sequence interpretation
func enableVirtualTerminalProcessing() {
	return
}
