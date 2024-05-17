package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/rwinkhart/MUTN/src/backend"

	"github.com/charmbracelet/glamour"
)

const ansiShownPassword = "\033[38;5;10m"

// EntryReader prints the decrypted contents of a libmutton entry in a human-readable format
func EntryReader(decryptedEntry []string, hidePassword bool, sync bool) {
	fmt.Println()

	for i := range decryptedEntry {
		switch i {
		case 0:
			// if the first field (password) is not empty, print it
			if decryptedEntry[0] != "" {
				if !hidePassword {
					fmt.Print(ansiDirectoryHeader + "Password:" + backend.AnsiReset + "\n" + ansiShownPassword + decryptedEntry[0] + backend.AnsiReset + "\n\n")
				} else {
					fmt.Print(ansiDirectoryHeader + "Password:" + backend.AnsiReset + "\n" + ansiEmptyDirectoryWarning + "End command in \"show\" or \"-s\" to view" + backend.AnsiReset + "\n\n")
				}
			}
		case 1:
			// if the second field (username) is not empty, print it
			if decryptedEntry[1] != "" {
				fmt.Print(ansiDirectoryHeader + "Username:" + backend.AnsiReset + "\n" + decryptedEntry[1] + "\n\n")
			}
		case 2:
			// if the third field (url) is not empty, print it
			if decryptedEntry[2] != "" {
				fmt.Print(ansiDirectoryHeader + "URL:" + backend.AnsiReset + "\n" + decryptedEntry[2] + "\n\n")
			}
		case 3:
			// print the notes header
			fmt.Println(ansiDirectoryHeader + "Notes:" + backend.AnsiReset)

			// combine remaining fields into a single string (for markdown rendering)
			var markdownNotes []string
			for field := 3; field < len(decryptedEntry); field++ {
				markdownNotes = append(markdownNotes, decryptedEntry[field])
			}
			r, _ := glamour.NewTermRenderer(glamour.WithStylesFromJSONBytes(glamourStyle()), glamour.WithPreservedNewLines(), glamour.WithWordWrap(width))
			markdownNotesString, _ := r.Render(strings.Join(markdownNotes, "\n"))

			// print markdown-rendered notes
			fmt.Print(markdownNotesString)

			// break after all lines have been printed
			break
		}
	}

	if sync && !backend.IsWindows {
		SshypSync() // TODO Remove after native sync is implemented
	}

	os.Exit(0)
}

// EntryReaderDecrypt is a wrapper for EntryReader that first decrypts a GPG-encrypted file before sending it to EntryReader
func EntryReaderDecrypt(targetLocation string, hidePassword bool, sync bool) {
	if isFile, _ := backend.TargetIsFile(targetLocation, true, 2); isFile {
		EntryReader(backend.DecryptGPG(targetLocation), hidePassword, sync)
	}
	// do not exit, as this is the job of EntryReader
}
