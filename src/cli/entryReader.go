package cli

import (
	"fmt"
	"os"

	"github.com/rwinkhart/libmutton/core"
	"github.com/rwinkhart/libmutton/sync"
)

const ansiShownPassword = "\033[38;5;10m"

// EntryReader prints the decrypted contents of a libmutton entry in a human-readable format.
func EntryReader(decryptedEntry []string, hideSecrets, syncEnabled bool) {
	fmt.Println()

	for field := range decryptedEntry {
		switch field {
		case 0:
			// if the first field (password) is not empty, print it
			if decryptedEntry[0] != "" {
				if !hideSecrets {
					fmt.Print(ansiDirectoryHeader + "Password:" + core.AnsiReset + "\n" + ansiShownPassword + decryptedEntry[0] + core.AnsiReset + "\n\n")
				} else {
					fmt.Print(ansiDirectoryHeader + "Password:" + core.AnsiReset + "\n" + ansiEmptyDirectoryWarning + "End command in \"show\" or \"-s\" to view" + core.AnsiReset + "\n\n")
				}
			}
		case 1:
			// if the second field (username) is not empty, print it
			if decryptedEntry[1] != "" {
				fmt.Print(ansiDirectoryHeader + "Username:" + core.AnsiReset + "\n" + decryptedEntry[1] + "\n\n")
			}
		case 2:
			// if the third field (TOTP secret) is not empty, print it
			if decryptedEntry[2] != "" {
				if !hideSecrets {
					fmt.Print(ansiDirectoryHeader + "TOTP Secret:" + core.AnsiReset + "\n" + ansiShownPassword + decryptedEntry[2] + core.AnsiReset + "\n\n")
				} else {
					fmt.Print(ansiDirectoryHeader + "TOTP Secret:" + core.AnsiReset + "\n" + ansiEmptyDirectoryWarning + "End command in \"show\" or \"-s\" to view" + core.AnsiReset + "\n\n")
				}
			}
		case 3:
			// if the fourth field (url) is not empty, print it
			if decryptedEntry[3] != "" {
				fmt.Print(ansiDirectoryHeader + "URL:" + core.AnsiReset + "\n" + decryptedEntry[3] + "\n\n")
			}
		case 4:
			// print the notes header
			fmt.Println(ansiDirectoryHeader + "Notes:" + core.AnsiReset)

			// combine remaining fields into a single string (to support Markdown rendering)
			var noteLines []string
			for ; field < len(decryptedEntry); field++ {
				noteLines = append(noteLines, decryptedEntry[field])
			}

			// print notes to stdout
			renderNote(noteLines)

			// break after notes have been printed
			break
		}
	}

	if syncEnabled {
		sync.RunJob(false)
	}

	os.Exit(0)
}

// EntryReaderDecrypt is a wrapper for EntryReader that first decrypts a GPG-encrypted file before sending it to EntryReader.
func EntryReaderDecrypt(targetLocation string, hideSecrets bool) {
	if isFile, _ := core.TargetIsFile(targetLocation, true, 2); isFile {
		EntryReader(core.DecryptGPG(targetLocation), hideSecrets, false) // never sync if decrypting straight to EntryReader, as this means the entry could not have been modified
	}
	// do not exit, as this is the job of EntryReader
}
