//go:build windows

package offline

// EntryRoot path to libmutton entry directory
var EntryRoot = home + "\\AppData\\Local\\libmutton\\entries"

// exported constants
const (
	PathSeparator = "\\"
	Windows       = true
)
