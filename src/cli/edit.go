package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/rwinkhart/MUTN/src/backend"
	"github.com/rwinkhart/MUTN/src/sync"
)

// RenameCli renames an entry at oldLocationIncomplete to a new location (user input) on both the client and the server
func RenameCli(oldLocationIncomplete string) {
	// prompt user for new location and rename
	newLocationIncomplete := input("New location:")
	sync.RenameRemoteFromClient(oldLocationIncomplete, newLocationIncomplete)

	// exit is done from sync.RenameRemoteFromClient
}

// EditEntryField edits a field of an entry at targetLocation (user input)
func EditEntryField(targetLocation string, hideSecrets bool, field int) {
	// fetch old entry data (with all required lines present)
	unencryptedEntry := backend.GetOldEntryData(targetLocation, field)

	// edit the field
	switch field {
	case 0:
		unencryptedEntry[field] = inputHidden("Password:")
	case 1:
		unencryptedEntry[field] = input("Username:")
	case 2:
		unencryptedEntry[field] = inputHidden("TOTP secret:")
	case 3:
		unencryptedEntry[field] = input("URL:")
	case 4: // edit notes fields
		// store note and non-note data separately
		nonNoteData := unencryptedEntry[:4]
		noteData := unencryptedEntry[4:]

		// edit the note
		editedNote, noteEdited := editNote(noteData)
		if !noteEdited { // exit early if the note was not edited
			fmt.Println(backend.AnsiError + "Entry is unchanged" + backend.AnsiReset)
			os.Exit(1)
		}
		unencryptedEntry = append(nonNoteData, editedNote...)
	}

	// write and preview the modified entry
	writeEntryCLI(targetLocation, unencryptedEntry, hideSecrets, false)
}

// GenUpdate generates a new password for an entry at targetLocation (user input)
func GenUpdate(targetLocation string, hideSecrets bool) {
	// fetch old entry data
	unencryptedEntry := backend.GetOldEntryData(targetLocation, 0)

	// generate a new password
	unencryptedEntry[0] = backend.StringGen(inputInt("Password length:", -1), inputBinary("Generate a complex (special characters) password?"), 0.2, false)

	// write and preview the modified entry
	writeEntryCLI(targetLocation, unencryptedEntry, hideSecrets, false)
}

// editNote uses the user-specified text editor to edit an existing note (or create a new one if baseNote is empty)
// returns the edited note and a boolean indicating whether the note was edited
func editNote(baseNote []string) ([]string, bool) {
	tempFile := backend.CreateTempFile()
	defer func(name string) {
		_ = os.Remove(name) // error ignored; if the file could be created, it can probably be removed
	}(tempFile.Name())

	// fetch the user's text editor
	editor := backend.ParseConfig([][2]string{{"MUTN", "textEditor"}}, "")[0]

	// write baseNote to tempFile (if it is not empty)
	if len(baseNote) > 0 {
		for _, line := range baseNote {
			_, _ = tempFile.WriteString(line + "\n")
		}
	}

	// close tempFile to allow it to be modified by the user's text editor
	_ = tempFile.Close() // error ignored; if the file could be created, it can probably be closed

	// edit the tempFile (note) with the user's text editor
	cmd := exec.Command(editor, tempFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(backend.AnsiError + "Failed to write note with " + editor + backend.AnsiReset) // panic is used to ensure the tempFile is removed, as per the defer statement
	}

	// open the tempFile for reading
	tempFile, err = os.Open(tempFile.Name())
	if err != nil {
		panic(backend.AnsiError + "Failed to read note written with " + editor + backend.AnsiReset) // panic is used to ensure the tempFile is removed, as per the defer statement
	}

	// read the edited note from the tempFile
	var note []string
	scanner := bufio.NewScanner(tempFile)
	for scanner.Scan() {
		note = append(note, scanner.Text())
	}

	// close tempFile
	_ = tempFile.Close() // error ignored; if the file could be opened, it can probably be closed

	// remove trailing empty strings from the edited note
	note = backend.RemoveTrailingEmptyStrings(note)

	// trim trailing whitespace from each note line
	for i, line := range note {
		note[i] = strings.TrimRight(line, " \t\r\n")
	}

	// return the edited note if it is different from baseNote
	if !reflect.DeepEqual(note, baseNote) {
		return note, true
	} else {
		return []string{}, false
	}
}
