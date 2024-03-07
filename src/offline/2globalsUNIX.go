//go:build !windows

package offline

// EntryRoot path to libmutton entry directory
var EntryRoot = home + "/.local/share/libmutton"
var ConfigPath = home + "/.config/libmutton/libmutton.ini"

// PathSeparator defines the character used to separate directories in a path (platform-specific)
const PathSeparator = "/"
