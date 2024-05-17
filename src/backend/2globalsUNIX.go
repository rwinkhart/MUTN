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
