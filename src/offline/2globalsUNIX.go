//go:build !windows

package offline

// EntryRoot path to libmutton entry directory
var EntryRoot = home + "/.local/share/libmutton"

// PathSeparator defines the character used to separate directories in a path (platform-specific)
const PathSeparator = "/"
