package backend

import (
	"fmt"
	"os"
)

func Shear(targetLocation string) {
	// TODO If in online mode, remove from server and add to shear list
	TargetIsFile(targetLocation, true, 0) // needed because os.RemoveAll does not return an error if target does not exist
	err := os.RemoveAll(targetLocation)
	if err != nil {
		fmt.Println(AnsiError + "Failed to remove target: " + err.Error() + AnsiReset)
		os.Exit(1)
	}
	os.Exit(0)
}
