//go:build !windows

package offline

// TargetLocationFormat returns the target location of an entry formatted for the current platform
func TargetLocationFormat(entryName string) string {
	return EntryRoot + PathSeparator + entryName
}
