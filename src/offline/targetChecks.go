package offline

import (
	"fmt"
	"os"
)

// TargetIsFile TargetStatusCheck checks if the targetLocation is a file, directory, or is inaccessible
// returns: isFile, isAccessible
func TargetIsFile(targetLocation string, errorOnFail bool) (bool, bool) {
	targetInfo, err := os.Stat(targetLocation)
	if err != nil {
		if errorOnFail {
			fmt.Println(AnsiError + "Failed to access " + targetLocation + " - ensure it exists and has the correct permissions" + AnsiReset)
			os.Exit(1)
		} else {
			return false, false
		}
	}
	if targetInfo.IsDir() {
		return false, true
	} else {
		return true, true
	}
}
