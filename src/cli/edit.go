package cli

import (
	"bufio"
	"os"
	"os/exec"
	"reflect"

	"github.com/rwinkhart/go-boilerplate/back"
	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/libmutton/core"
	"github.com/rwinkhart/libmutton/syncclient"
)

// RenameCli renames an entry at oldLocationIncomplete to a new location (user input) on both the client and the server.
func RenameCli(oldLocationIncomplete string) {
	// prompt user for new location and rename
	newLocationIncomplete := front.Input("New location:")
	syncclient.RenameRemoteFromClient(oldLocationIncomplete, newLocationIncomplete, false)

	// exit is done from sync.RenameRemoteFromClient
}

// EditEntryField edits a field of an entry at targetLocation (user input).
func EditEntryField(targetLocation string, hideSecrets bool, field int) {
	// fetch old entry data (with all required lines present)
	unencryptedEntry := core.GetOldEntryData(targetLocation, field)

	// edit the field
	switch field {
	case 0:
		unencryptedEntry[field] = string(front.InputHidden("Password:"))
	case 1:
		unencryptedEntry[field] = front.Input("Username:")
	case 2:
		unencryptedEntry[field] = string(front.InputHidden("TOTP secret:"))
	case 3:
		unencryptedEntry[field] = front.Input("URL:")
	case 4: // edit notes fields
		// store note and non-note data separately
		nonNoteData := unencryptedEntry[:4]
		noteData := unencryptedEntry[4:]

		// edit the note
		editedNote, noteEdited := editNote(noteData)
		if !noteEdited { // exit early if the note was not edited
			back.PrintError("Entry is unchanged", 0, true)
		}
		unencryptedEntry = append(nonNoteData, editedNote...)
	}

	// write and preview the modified entry
	writeEntryCLI(targetLocation, unencryptedEntry, hideSecrets)
}

// GenUpdate generates a new password for an entry at targetLocation (user input).
func GenUpdate(targetLocation string, hideSecrets bool) {
	// fetch old entry data
	unencryptedEntry := core.GetOldEntryData(targetLocation, 0)

	// generate a new password
	unencryptedEntry[0] = inputPasswordGen()

	// write and preview the modified entry
	writeEntryCLI(targetLocation, unencryptedEntry, hideSecrets)
}

// editNote uses the user-specified text editor to edit an existing note (or create a new one if baseNote is empty).
// Returns the edited note and a boolean indicating whether the note was edited.
func editNote(baseNote []string) ([]string, bool) {
	tempFile := back.CreateTempFile()
	defer func(name string) {
		_ = os.Remove(name) // error ignored; if the file could be created, it can probably be removed
	}(tempFile.Name())

	// fetch the user's text editor
	editorCfg, _ := core.ParseConfig([][2]string{{"MUTN", "textEditor"}}, "")
	editor := editorCfg[0]

	// write baseNote to tempFile (if it is not empty)
	baseNoteLen := len(baseNote)
	switch baseNoteLen {
	case 0: // do not write empty baseNote to file
	case 1: // truncate baseNote if its only length is an empty string (do not write to file)
		if baseNote[0] == "" {
			baseNote = []string{}
			baseNoteLen = 0
			break
		}
		fallthrough // if baseNote contained real data, write it to tempFile
	default: // write baseNote to tempFile for external editing
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
		panic(back.AnsiError + "Failed to write note with " + editor + back.AnsiReset) // panic is used to ensure the tempFile is removed, as per the defer statement
	}

	// open the tempFile for reading
	tempFile, err = os.Open(tempFile.Name())
	if err != nil {
		panic(back.AnsiError + "Failed to read note written with " + editor + back.AnsiReset) // panic is used to ensure the tempFile is removed, as per the defer statement
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
	note = back.RemoveTrailingEmptyStrings(note)

	// clamp trailing whitespace in each note line
	core.ClampTrailingWhitespace(note)

	// return the edited note if:
	// it is different from baseNote AND
	// [it is not empty (prevent needless writes when adding blank notes) OR
	// it has a different length from baseNote (allow removing notes from entries)]
	noteLen := len(note)
	if !reflect.DeepEqual(note, baseNote) && (noteLen > 0 || baseNoteLen != noteLen) {
		return note, true
	} else {
		return nil, false
	}
}
