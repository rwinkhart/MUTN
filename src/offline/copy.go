package offline

import (
	"bufio"
	"fmt"
	"github.com/pquerna/otp/totp"
	"os"
	"time"
)

// CopyArgument copies a field from an entry to the clipboard
func CopyArgument(targetLocation string, field int, executableName string) {
	if isFile, _ := TargetIsFile(targetLocation, true, 2); isFile {

		decryptedEntry := DecryptGPG(targetLocation)
		var copySubject string // will store data to be copied

		// ensure field exists in entry
		if len(decryptedEntry) > field {

			// ensure field is not empty
			if decryptedEntry[field] == "" {
				fmt.Println(AnsiError + "Field is empty" + AnsiReset)
				os.Exit(1)
			}

			if field != 5 { // TODO Update field after removed from notes (breaking sshyp entry compatibility)
				copySubject = decryptedEntry[field]
			} else { // TOTP mode
				var err error
				var currentSecond int // track current second for TOTP code refresh
				var copied bool       // track whether the TOTP code has been copied to the clipboard for the first time
				for {                 // keep field copied to clipboard, refresh on 30-second intervals
					currentSecond = time.Now().Second()
					if currentSecond == 0 || currentSecond == 30 || !copied {
						copySubject, err = totp.GenerateCode(decryptedEntry[5], time.Now())
						if err != nil {
							fmt.Println(AnsiError + "Error generating TOTP code" + AnsiReset)
							os.Exit(1)
						}
						copyField(copySubject, "")
					}
					time.Sleep(1 * time.Second)
				}
			}
		} else {
			fmt.Println(AnsiError + "Field does not exist in entry" + AnsiReset)
			os.Exit(1)
		}

		// copy field to clipboard, launch clipboard clearing process
		copyField(copySubject, executableName)
	}
}

// ClipClearArgument is called to clear the clipboard after 30 seconds if the contents have not been modified
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
