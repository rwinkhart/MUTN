//go:build !windows

package main

// EntryRoot path to libmutton entry directory
var EntryRoot = home + "/.local/share/fake"

// exported constants
const (
	PathSeparator = "/"
	Windows       = false
)
