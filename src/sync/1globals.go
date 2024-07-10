package sync

import "github.com/rwinkhart/MUTN/src/backend"

// define field separator constants
const (
	FSSpace = "\u259d" // ▝ space/list separator
	FSPath  = "\u259e" // ▞ path separator
	FSMisc  = "\u259f" // ▟ misc. field separator (if \u259d is already used)
)

// rootLength stores length of backend.EntryRoot string
var rootLength = len(backend.EntryRoot)
