//go:build windows

package backend

// EntryRoot path to libmutton entry directory
var EntryRoot = home + "\\AppData\\Local\\libmutton\\entries"
var ConfigDir = home + "\\AppData\\Local\\libmutton\\config"
var ConfigPath = ConfigDir + "\\libmutton.ini"

// PathSeparator defines the character used to separate directories in a path (platform-specific)
const (
	PathSeparator = "\\"
	Windows       = true // TODO temporary, remove after native sync is implemented
)
