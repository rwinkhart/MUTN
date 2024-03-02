package cli

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/offline"
	"os"
)

const ansiShownPassword = "\033[38;5;10m"

// EntryReader prints the decrypted contents of a libmutton entry in a human-readable format
func EntryReader(decryptedEntry []string, hidePassword bool) {
	fmt.Println()

	// track if extended notes have been printed (to avoid printing an extra newline)
	notesFlag := false

	for i := range decryptedEntry {

		// if the entry only contains a password, skip the username field to avoid an index out of range error
		if len(decryptedEntry) == 1 {
			i = 1
		}

		switch i {
		case 0:
			// if the second field (username) is not empty, print it
			if decryptedEntry[1] != "" {
				fmt.Print(ansiDirectoryHeader + "Username:" + offline.AnsiReset + "\n" + decryptedEntry[1] + "\n\n")
			}
		case 1:
			// if the first field (password) is not empty, print it
			if decryptedEntry[0] != "" {
				if !hidePassword {
					fmt.Print(ansiDirectoryHeader + "Password:" + offline.AnsiReset + "\n" + ansiShownPassword + decryptedEntry[0] + offline.AnsiReset + "\n\n")
				} else {
					fmt.Print(ansiDirectoryHeader + "Password:" + offline.AnsiReset + "\n" + ansiEmptyDirectoryWarning + "End command in \"show\" or \"-s\" to view" + offline.AnsiReset + "\n\n")
				}
			}
		case 2:
			// if the third field (url) is not empty, print it
			if decryptedEntry[2] != "" {
				fmt.Print(ansiDirectoryHeader + "URL:" + offline.AnsiReset + "\n" + decryptedEntry[2] + "\n\n")
			}
		case 3:
			// if the fourth field (notes begin) is not empty, print it
			if decryptedEntry[3] != "" {
				fmt.Println(ansiDirectoryHeader + "Notes:" + offline.AnsiReset + "\n" + decryptedEntry[3])
			}
		default:
			// print extended notes line
			fmt.Println(decryptedEntry[i])

			// indicate that extended notes have been printed
			if !notesFlag {
				notesFlag = true
			}
		}
	}
	// print trailing newline if extended notes were printed
	if notesFlag {
		fmt.Println()
	}
	os.Exit(0)
}

func EntryReaderShortcut(targetLocation string, hidePassword bool) {
	if isFile, _ := offline.TargetIsFile(targetLocation, true); isFile {
		EntryReader(offline.DecryptGPG(targetLocation), hidePassword)
	} else {
		fmt.Println(offline.AnsiError + "Failed to read \"" + targetLocation + "\" - it is a directory" + offline.AnsiReset)
	}
}
