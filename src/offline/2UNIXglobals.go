//go:build !windows

package offline

// EntryRoot path to libmutton entry directory
var EntryRoot = home + "/.local/share/libmutton"

// exported constants
const (
	PathSeparator = "/"
	Windows       = false
)
