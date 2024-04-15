package backend

import (
	"bufio"
	"fmt"
	steamtotp "github.com/fortis/go-steam-totp"
	"github.com/pquerna/otp/totp"
	"os"
	"strings"
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
				var secret string // stores secret for TOTP generation
				var forSteam bool // indicates whether to generate TOTP in Steam format

				if strings.HasPrefix(decryptedEntry[5], "steam@") {
					secret = decryptedEntry[5][6:]
					forSteam = true
				} else {
					secret = decryptedEntry[5]
				}

				for { // keep field copied to clipboard, refresh on 30-second intervals
					currentTime := time.Now()
					copyField(GenTOTP(secret, currentTime, forSteam), "")
					// sleep until next 30-second interval
					time.Sleep(time.Duration(30-(currentTime.Second()%30)) * time.Second)
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

// GenTOTP generates a TOTP token from a secret (supports standard and Steam TOTP)
func GenTOTP(secret string, time time.Time, forSteam bool) string {
	var totpToken string
	var err error

	if forSteam {
		totpToken, err = steamtotp.GenerateAuthCode(secret, time)
	} else {
		totpToken, err = totp.GenerateCode(secret, time)
	}

	if err != nil {
		fmt.Println(AnsiError + "Error generating TOTP code" + AnsiReset)
		os.Exit(1)
	}

	return totpToken
}
