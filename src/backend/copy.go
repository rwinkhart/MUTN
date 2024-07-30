package backend

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	steamtotp "github.com/fortis/go-steam-totp"
	"github.com/pquerna/otp/totp"
)

// CopyArgument copies a field from an entry to the clipboard
func CopyArgument(executableName, targetLocation string, field int) {
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

			if field != 2 {
				copySubject = decryptedEntry[field]
			} else { // TOTP mode
				var secret string // stores secret for TOTP generation
				var forSteam bool // indicates whether to generate TOTP in Steam format

				if strings.HasPrefix(decryptedEntry[2], "steam@") {
					secret = decryptedEntry[2][6:]
					forSteam = true
				} else {
					secret = decryptedEntry[2]
				}

				fmt.Println("TOTP code has been copied to the clipboard - your clipboard will be kept up to date with the current code until this process is closed")

				for { // keep field copied to clipboard, refresh on 30-second intervals
					currentTime := time.Now()
					copyField("", GenTOTP(secret, currentTime, forSteam))
					// sleep until next 30-second interval
					time.Sleep(time.Duration(30-(currentTime.Second()%30)) * time.Second)
				}
			}
		} else {
			fmt.Println(AnsiError + "Field does not exist in entry" + AnsiReset)
			os.Exit(1)
		}

		// copy field to clipboard, launch clipboard clearing process
		copyField(executableName, copySubject)
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
		os.Exit(0) // use os.Exit instead of backend.Exit, as this function runs out of a background subprocess that is invisible to the user (will never appear in GUI/TUI environment)
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
