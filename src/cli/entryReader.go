package cli

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/offline"
	"os"
)

const ansiShownPassword = "\033[38;5;10m"

// EntryReader prints the decrypted contents of a libmutton entry in a human-readable format
func EntryReader(decryptedEntry []string, hidePassword bool, sync bool) {
	fmt.Println()

	// track if extended notes have been printed (to avoid printing an extra newline)
	var notesFlag bool

	for i := range decryptedEntry {
		switch i {
		case 0:
			// if the first field (password) is not empty, print it
			if decryptedEntry[0] != "" {
				if !hidePassword {
					fmt.Print(ansiDirectoryHeader + "Password:" + offline.AnsiReset + "\n" + ansiShownPassword + decryptedEntry[0] + offline.AnsiReset + "\n\n")
				} else {
					fmt.Print(ansiDirectoryHeader + "Password:" + offline.AnsiReset + "\n" + ansiEmptyDirectoryWarning + "End command in \"show\" or \"-s\" to view" + offline.AnsiReset + "\n\n")
				}
			}
		case 1:
			// if the second field (username) is not empty, print it
			if decryptedEntry[1] != "" {
				fmt.Print(ansiDirectoryHeader + "Username:" + offline.AnsiReset + "\n" + decryptedEntry[1] + "\n\n")
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
				// indicate that the first notes line was printed (not empty)
				notesFlag = true
			}
		default:
			// print notes header if current line is not empty and notes have not been printed
			if !notesFlag && decryptedEntry[i] != "" {
				fmt.Println(ansiDirectoryHeader + "Notes:" + offline.AnsiReset)
				// indicate that notes have been printed
				notesFlag = true
			}

			// print extended notes line if header was printed
			if notesFlag {
				fmt.Println(decryptedEntry[i])
			}
		}
	}
	// print trailing newline if notes were printed
	if notesFlag {
		fmt.Println()
	}

	if sync && !offline.Windows {
		SshypSync() // TODO Remove after native sync is implemented
	}

	os.Exit(0)
}

// EntryReaderShortcut is a shortcut for EntryReader that decrypts a GPG-encrypted file and prints the contents
func EntryReaderShortcut(targetLocation string, hidePassword bool, sync bool) {
	if isFile, _ := offline.TargetIsFile(targetLocation, true, 2); isFile {
		EntryReader(offline.DecryptGPG(targetLocation), hidePassword, sync)
	}
	// do not exit, as this is the job of EntryReader
}
