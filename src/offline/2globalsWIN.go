//go:build windows

package offline

// EntryRoot path to libmutton entry directory
var EntryRoot = home + "\\AppData\\Local\\libmutton\\entries"
var ConfigPath = home + "\\AppData\\Local\\libmutton\\libmutton.ini"

// PathSeparator defines the character used to separate directories in a path (platform-specific)
const PathSeparator = "\\"
