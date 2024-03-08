package cli

import (
	"bufio"
	"fmt"
	"github.com/rwinkhart/MUTN/src/offline"
	"os"
	"os/exec"
	"reflect"
)

// RenameCli renames an entry at oldLocation to a new location (user input)
func RenameCli(oldLocation string) {
	// ensure targetLocation exists
	offline.TargetIsFile(oldLocation, true, 0)

	// prompt user for new location and rename
	newLocation := offline.EntryRoot + offline.PathSeparator + input("New location:")
	offline.Rename(oldLocation, newLocation)

	// exit is done from offline.Rename
}

// EditEntry edits a field of an entry at targetLocation (user input), does not allow for editing notes
func EditEntry(targetLocation string, hidePassword bool, field int) {
	// ensure targetLocation exists
	offline.TargetIsFile(targetLocation, true, 2)

	// read old entry data
	unencryptedEntry := offline.DecryptGPG(targetLocation)

	// ensure slice is long enough for field
	for len(unencryptedEntry) <= field {
		unencryptedEntry = append(unencryptedEntry, "")
	}

	// edit the field
	switch field {
	case 0:
		unencryptedEntry[field] = inputHidden("Password:")
	case 1:
		unencryptedEntry[field] = input("Username:")
	case 2:
		unencryptedEntry[field] = input("URL:")
	}

	// write and preview the modified entry
	writeEntryShortcut(targetLocation, unencryptedEntry, hidePassword)
}

func EditEntryNote(targetLocation string, hidePassword bool) {
	// ensure targetLocation exists
	offline.TargetIsFile(targetLocation, true, 2)

	// read old entry data
	unencryptedEntry := offline.DecryptGPG(targetLocation)

	// store non-note data separately
	nonNoteData := unencryptedEntry[:3]

	// store note data separately
	noteData := unencryptedEntry[3:]

	// edit the note
	editedNote, noteEdited := editNote(noteData)
	if !noteEdited { // exit early if the note was not edited
		fmt.Println(offline.AnsiError + "Entry is unchanged" + offline.AnsiReset)
		os.Exit(1)
	}
	unencryptedEntry = append(nonNoteData, editedNote...)

	// write and preview the modified entry
	writeEntryShortcut(targetLocation, unencryptedEntry, hidePassword)
}

// GenUpdate generates a new password for an entry at targetLocation (user input)
func GenUpdate(targetLocation string, hidePassword bool) {
	// ensure targetLocation exists
	offline.TargetIsFile(targetLocation, true, 2)

	// read old entry data
	unencryptedEntry := offline.DecryptGPG(targetLocation)

	// generate a new password
	unencryptedEntry[0] = offline.StringGen(inputInt("Password length:", -1), inputBinary("Generate a complex (special characters) password?"), 0.2)

	// write and preview the modified entry
	writeEntryShortcut(targetLocation, unencryptedEntry, hidePassword)
}

// editNote uses the user-specified text editor to edit an existing note (or create a new one if baseNote is empty)
// returns the edited note and a boolean indicating whether the note was edited
func editNote(baseNote []string) ([]string, bool) {
	tempFile := offline.CreateTempFile()
	defer os.Remove(tempFile.Name())
	editor := offline.ReadConfig([]string{"textEditor"})[0]

	// write baseNote to tempFile (if it is not empty)
	if len(baseNote) > 0 {
		for _, line := range baseNote {
			_, _ = tempFile.WriteString(line + "\n")
		}
	}

	// edit the tempFile (note) with the user's text editor
	cmd := exec.Command(editor, tempFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(offline.AnsiError + "Failed to write note with " + editor + offline.AnsiReset)
		os.Exit(1)
	}

	// open the tempFile for reading
	file, err := os.Open(tempFile.Name())
	if err != nil {
		fmt.Println(offline.AnsiError + "Failed to open temporary file (\"" + tempFile.Name() + "\") " + err.Error() + offline.AnsiReset)
		os.Exit(1)
	}
	defer file.Close()

	// read the edited note from the tempFile
	var note []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		note = append(note, scanner.Text())
	}

	// return the edited note if it is different from baseNote
	if !reflect.DeepEqual(offline.RemoveTrailingEmptyStrings(note), baseNote) {
		return note, true
	} else {
		return note, false
	}
}
