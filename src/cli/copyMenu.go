package cli

import (
	"fmt"

	"github.com/rwinkhart/go-boilerplate/back"
	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
	"github.com/rwinkhart/libmutton/clip"
	"github.com/rwinkhart/libmutton/crypt"
	"github.com/rwinkhart/libmutton/global"
	"github.com/rwinkhart/libmutton/synccommon"
)

// CopyMenu decrypts an entry and allows the user to
// interactively copy fields without having to re-decrypt each time.
func CopyMenu(targetLocation string) {
	// decrypt entry
	decSlice, err := crypt.DecryptFileToSlice(targetLocation)
	if err != nil {
		other.PrintError("Failed to decrypt entry: "+err.Error(), global.ErrorDecryption)
	}

	// determine populated fields in entry
	fieldIndexToString := map[int]string{
		0: "Password",
		1: "Username",
		2: "TOTP Code",
		3: "URL",
		4: "Note (first line)",
	}
	var fields []string
	for i, _ := range decSlice[:5] {
		if decSlice[i] != "" {
			fields = append(fields, fieldIndexToString[i])
		}
	}

	// copy selected field to clipboard
	for {
		choice := front.InputMenuGen("Field to copy:", fields)
		switch fields[choice-1] {
		case "Password":
			choice = 0
		case "Username":
			choice = 1
		case "TOTP Code":
			choice = 2
		case "URL":
			choice = 3
		case "Note (first line)":
			choice = 4
		}
		if choice == 2 {
			fmt.Println(ansiEmptyDirectoryWarning + "[Starting]" + back.AnsiReset + " TOTP clipboard refresher")
			errorChan := make(chan error)
			done := make(chan bool)
			go clip.TOTPCopier(decSlice[2], 0, errorChan, done)
			// block until first successful copy
			err = <-errorChan
			if err != nil {
				other.PrintError("Failed to copy TOTP token: "+err.Error(), global.ErrorClipboard)
			}
			fmt.Println(synccommon.AnsiDownload + "[Started]" + back.AnsiReset + " TOTP clipboard refresher")
			// block until the user sends "q"
			for {
				input := front.Input("Service will run until 'q' is entered:")
				if input == "q" {
					break
				}
			}
			close(done)
			fmt.Print(synccommon.AnsiUpload + "\n[Stopped]" + back.AnsiReset + " TOTP clipboard refresher\n\n")
		} else {
			err := clip.CopyString(true, decSlice[choice])
			if err != nil {
				other.PrintError("Failed to copy field to clipboard: "+err.Error(), global.ErrorClipboard)
			}
		}
	}
	// TODO clear clipboard on exit!
}
