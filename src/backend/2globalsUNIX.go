//go:build !windows

package backend

// EntryRoot path to libmutton entry directory
var EntryRoot = home + "/.local/share/libmutton"
var ConfigDir = home + "/.config/libmutton"
var ConfigPath = ConfigDir + "/libmutton.ini"

// PathSeparator defines the character used to separate directories in a path (platform-specific)
const (
	PathSeparator = "/"
	Windows       = false // TODO temporary, remove after native sync is implemented
)
