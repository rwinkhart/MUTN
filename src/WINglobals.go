//go:build windows

package main

// EntryRoot path to libmutton entry directory
var EntryRoot = home + "\\AppData\\Local\\libmutton\\entries"

// exported constants
const (
	PathSeparator = "\\"
	Windows       = true
)
