package offline

import (
	"fmt"
	"os"
)

func CopyField(targetLocation string, field rune) {
	if isFile, _ := TargetIsFile(targetLocation, true); isFile {
		//decryptedEntry := DecryptGPG(targetLocation)
		// TODO implement clipboard support (maybe use third-party library?)
		fmt.Println("In the future, this will copy " + string(field))
	} else {
		fmt.Println(AnsiError + "Failed to read \"" + targetLocation + "\" - it is a directory" + AnsiReset)
	}
	os.Exit(0)
}

func clipClear() {

}
