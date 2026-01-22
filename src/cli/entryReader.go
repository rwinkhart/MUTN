package cli

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/rwinkhart/go-boilerplate/back"
	"github.com/rwinkhart/go-boilerplate/other"
	"github.com/rwinkhart/libmutton/crypt"
	"github.com/rwinkhart/libmutton/global"
)

const ansiShownPassword = "\033[38;5;10m"

// EntryReader prints the decrypted contents of a libmutton entry in a human-readable format.
func EntryReader(vanityPath string, decSlice []string, hideSecrets bool) {
	fmt.Print("\n" + back.AnsiBold + vanityPath + back.AnsiReset + "\n\n")

fieldLoop:
	for field := range decSlice {
		switch field {
		case 0:
			// if the first field (password) is not empty, print it
			if decSlice[0] != "" {
				if !hideSecrets {
					fmt.Print(ansiDirectoryHeader + "Password:" + back.AnsiReset + "\n" + ansiShownPassword + decSlice[0] + back.AnsiReset + "\n\n")
				} else {
					fmt.Print(ansiDirectoryHeader + "Password:" + back.AnsiReset + "\n" + back.AnsiWarning + "End command in \"show\" or \"-s\" to view" + back.AnsiReset + "\n\n")
				}
			}
		case 1:
			// if the second field (username) is not empty, print it
			if decSlice[1] != "" {
				fmt.Print(ansiDirectoryHeader + "Username:" + back.AnsiReset + "\n" + decSlice[1] + "\n\n")
			}
		case 2:
			// if the third field (TOTP secret) is not empty, print it
			if decSlice[2] != "" {
				if !hideSecrets {
					fmt.Print(ansiDirectoryHeader + "TOTP Secret:" + back.AnsiReset + "\n" + ansiShownPassword + decSlice[2] + back.AnsiReset + "\n\n")
				} else {
					fmt.Print(ansiDirectoryHeader + "TOTP Secret:" + back.AnsiReset + "\n" + back.AnsiWarning + "End command in \"show\" or \"-s\" to view" + back.AnsiReset + "\n\n")
				}
			}
		case 3:
			// if the fourth field (url) is not empty, print it
			if decSlice[3] != "" {
				fmt.Print(ansiDirectoryHeader + "URL:" + back.AnsiReset + "\n" + decSlice[3] + "\n\n")
			}
		case 4:
			// print the notes header
			fmt.Println(ansiDirectoryHeader + "Notes:" + back.AnsiReset)

			// render notes as Markdown
			r, _ := glamour.NewTermRenderer(glamour.WithStylesFromJSONBytes(glamourStyle()), glamour.WithPreservedNewLines(), glamour.WithWordWrap(width))
			markdownNotesString, _ := r.Render(strings.Join(decSlice[4:], "\n"))
			_ = r.Close()

			// print the rendered Markdown notes
			fmt.Print(markdownNotesString)

			// break after notes have been printed
			break fieldLoop
		}
	}
}

// EntryReaderDecrypt is a wrapper for EntryReader that first decrypts an RCW-wrapped file before sending it to EntryReader.
func EntryReaderDecrypt(realPath string, hideSecrets bool) {
	_, err := back.TargetIsFile(realPath, true)
	if err != nil { // if the location does not exist or is a directory...
		other.PrintError("Failed to verify target location: "+err.Error(), back.ErrorTargetNotFound)
	}
	decSlice, err := crypt.DecryptFileToSlice(realPath, nil)
	if err != nil {
		other.PrintError("Failed to decrypt entry: "+err.Error(), global.ErrorDecryption)
	}
	EntryReader(global.GetVanityPath(realPath), decSlice, hideSecrets)
}
