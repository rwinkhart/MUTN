//go:build windows

package main

// EntryRoot path to libmutton entry directory TODO to be moved to libmutton
var EntryRoot = home + "\\AppData\\Local\\libmutton\\entries"

// exported constants TODO to be moved to libmutton
const (
	PathSeparator = "\\"
	Windows       = true
)
