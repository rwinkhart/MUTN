package cli

import (
	"bufio"
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

// RenameCli renames an entry at oldLocation to a new location (user input)
func RenameCli(oldLocation string) {
	// ensure targetLocation exists
	backend.TargetIsFile(oldLocation, true, 0)

	// prompt user for new location and rename
	newLocation := backend.EntryRoot + backend.PathSeparator + input("New location:")
	backend.Rename(oldLocation, newLocation)

	// exit is done from backend.Rename
}

// EditEntryField edits a field of an entry at targetLocation (user input)
func EditEntryField(targetLocation string, hidePassword bool, field int) {
	// fetch old entry data (with all required lines present)
	unencryptedEntry := backend.GetOldEntryData(targetLocation, field)

	// edit the field
	switch field {
	case 0:
		unencryptedEntry[field] = inputHidden("Password:")
	case 1:
		unencryptedEntry[field] = input("Username:")
	case 2:
		unencryptedEntry[field] = input("URL:")
	case 3: // edit notes fields
		// store note and non-note data separately
		nonNoteData := unencryptedEntry[:3]
		noteData := unencryptedEntry[3:]

		// edit the note
		editedNote, noteEdited := editNote(noteData)
		if !noteEdited { // exit early if the note was not edited
			fmt.Println(backend.AnsiError + "Entry is unchanged" + backend.AnsiReset)
			os.Exit(1)
		}
		unencryptedEntry = append(nonNoteData, editedNote...)
	}

	// write and preview the modified entry
	writeEntryCLI(targetLocation, unencryptedEntry, hidePassword)
}

// GenUpdate generates a new password for an entry at targetLocation (user input)
func GenUpdate(targetLocation string, hidePassword bool) {
	// fetch old entry data
	unencryptedEntry := backend.GetOldEntryData(targetLocation, 0)

	// generate a new password
	unencryptedEntry[0] = backend.StringGen(inputInt("Password length:", -1), inputBinary("Generate a complex (special characters) password?"), 0.2)

	// write and preview the modified entry
	writeEntryCLI(targetLocation, unencryptedEntry, hidePassword)
}

// editNote uses the user-specified text editor to edit an existing note (or create a new one if baseNote is empty)
// returns the edited note and a boolean indicating whether the note was edited
func editNote(baseNote []string) ([]string, bool) {
	tempFile := backend.CreateTempFile()
	defer os.Remove(tempFile.Name())
	editor := backend.ReadConfig([]string{"textEditor"})[0]

	// write baseNote to tempFile (if it is not empty)
	if len(baseNote) > 0 {
		for _, line := range baseNote {
			_, _ = tempFile.WriteString(line + "\n")
		}
	}

	// close tempFile
	tempFile.Close()

	// edit the tempFile (note) with the user's text editor
	cmd := exec.Command(editor, tempFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(backend.AnsiError + "Failed to write note with " + editor + backend.AnsiReset)
	}

	// open the tempFile for reading
	tempFile, err = os.Open(tempFile.Name())
	if err != nil {
		panic(backend.AnsiError + "Failed to write note with " + editor + backend.AnsiReset)
	}

	// read the edited note from the tempFile
	var note []string
	scanner := bufio.NewScanner(tempFile)
	for scanner.Scan() {
		note = append(note, scanner.Text())
	}

	// close tempFile
	tempFile.Close()

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
