//go:build windows

package backend

import "strings"

// TargetLocationFormat returns the target location of an entry formatted for the current platform
func TargetLocationFormat(entryName string) string {
	return EntryRoot + PathSeparator + strings.ReplaceAll(entryName, "/", PathSeparator)
}
