//go:build windows

package backend

// EntryRoot path to libmutton entry directory
var EntryRoot = Home + "\\AppData\\Local\\libmutton\\entries"
var ConfigDir = Home + "\\AppData\\Local\\libmutton\\config"
var ConfigPath = ConfigDir + "\\libmutton.ini"

// PathSeparator defines the character used to separate directories in a path (platform-specific)
const (
	PathSeparator = "\\"
	IsWindows     = true
)
