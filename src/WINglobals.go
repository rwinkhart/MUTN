//go:build windows

package main

// exported variables
var EntryRoot = home + "\\AppData\\Local\\libmutton\\entries"

// exported constants
const (
	PathSeparator = "\\"
	Windows       = true
)
