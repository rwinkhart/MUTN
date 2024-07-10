package backend

// GetOldEntryData decrypts and returns old entry data (with all required lines present)
func GetOldEntryData(targetLocation string, field int) []string {
	// ensure targetLocation exists
	TargetIsFile(targetLocation, true, 2)

	// read old entry data
	unencryptedEntry := DecryptGPG(targetLocation)

	// return the old entry data with all required lines present
	if field > 0 {
		return EnsureSliceLength(unencryptedEntry, field)
	} else {
		return unencryptedEntry
	}
}

// EnsureSliceLength ensures slice is long enough to contain the specified index
func EnsureSliceLength(slice []string, index int) []string {
	for len(slice) <= index {
		slice = append(slice, "")
	}
	return slice
}
