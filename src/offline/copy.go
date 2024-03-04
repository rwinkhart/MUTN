package offline

import (
	"bufio"
	"fmt"
	"os"
)

func CopyArgument(targetLocation string, field int, executableName string) {
	if isFile, _ := TargetIsFile(targetLocation, true, 2); isFile {

		decryptedEntry := DecryptGPG(targetLocation)
		var copySubject string // will store data to be copied

		// ensure field exists in entry
		if len(decryptedEntry) > field {
			copySubject = decryptedEntry[field]
		} else {
			fmt.Println(AnsiError + "Field does not exist in entry" + AnsiReset)
			os.Exit(1)
		}

		// ensure field is not blank
		if copySubject == "" {
			fmt.Println(AnsiError + "Field is empty" + AnsiReset)
			os.Exit(1)
		}

		// copy field to clipboard, launch clipboard clearing process
		copyField(copySubject, executableName)
	}
}

func ClipClearArgument() {
	// read previous clipboard contents from stdin
	clipScanner := bufio.NewScanner(os.Stdin)
	if clipScanner.Scan() {
		oldContents := clipScanner.Text()
		clipClear(oldContents)
	} else {
		os.Exit(0)
	}
}
