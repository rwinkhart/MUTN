package cli

import (
	"bufio"
	"os"
	"os/exec"
	"reflect"

	"github.com/rwinkhart/go-boilerplate/back"
	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
	"github.com/rwinkhart/libmutton/cfg"
	"github.com/rwinkhart/libmutton/core"
	"github.com/rwinkhart/libmutton/syncclient"
)

// RenameCli renames an entry at oldLocationIncomplete to a new location (user input) on both the client and the server.
func RenameCli(oldVanityPath string) {
	// prompt user for new location and rename
	newVanityPath := front.Input("New location:")
	err := syncclient.RenameRemote(oldVanityPath, newVanityPath)
	if err != nil {
		other.PrintError("Failed to rename entry: "+err.Error(), back.ErrorWrite)
	}

	// exit is done from sync.RenameRemote
}

// EditEntryField edits a field of an entry at realPath (user input).
func EditEntryField(realPath string, field int) {
	// fetch old entry data (with all required lines present)
	decSlice, err := core.GetOldEntryData(realPath, field)
	if err != nil {
		other.PrintError("Failed to fetch entry data: "+err.Error(), back.ErrorRead)
	}

	// edit the field
	var oldPassword string
	switch field {
	case 0:
		oldPassword = decSlice[field]
		decSlice[field] = string(front.InputHidden("Password:"))
	case 1:
		decSlice[field] = front.Input("Username:")
	case 2:
		decSlice[field] = string(front.InputHidden("TOTP secret:"))
	case 3:
		decSlice[field] = front.Input("URL:")
	case 4: // edit notes fields
		// store note and non-note data separately
		fieldsMain := decSlice[:4]
		fieldsNote := decSlice[4:]

		// edit the note
		editedNote, noteEdited := editNote(fieldsNote)
		if !noteEdited { // exit early if the note was not edited
			other.PrintError("Entry is unchanged", 0)
		}
		decSlice = append(fieldsMain, editedNote...)
	}

	// write and preview the modified entry
	if field == 0 {
		writeEntryCLI(realPath, decSlice, true, oldPassword)
	} else {
		writeEntryCLI(realPath, decSlice, false, oldPassword)
	}
}

// GenUpdate generates a new password for an entry at realPath (user input).
func GenUpdate(realPath string) {
	// fetch old entry data
	decSlice, err := core.GetOldEntryData(realPath, 0)
	if err != nil {
		other.PrintError("Failed to fetch entry data: "+err.Error(), back.ErrorRead)
	}

	// generate a new password
	oldPassword := decSlice[0]
	decSlice[0] = inputPasswordGen()

	// write and preview the modified entry
	writeEntryCLI(realPath, decSlice, true, oldPassword)
}

// editNote uses the user-specified text editor to edit an existing note (or create a new one if baseNote is empty).
// Returns the edited note and a boolean indicating whether the note was edited.
func editNote(baseNote []string) ([]string, bool) {
	tempFile, err := back.CreateTempFile()
	if err != nil {
		other.PrintError("Failed to create temporary note file: "+err.Error(), back.ErrorWrite)
	}
	defer func(name string) {
		_ = os.Remove(name) // error ignored; if the file could be created, it can probably be removed
	}(tempFile.Name())

	// fetch the user's text editor
	config, err := cfg.LoadConfig()
	if err != nil {
		other.PrintError("Failed to load libmuttoncfg.json: "+err.Error(), back.ErrorRead)
	}
	var editor string
	if value, exists := (*config.ThirdParty)["mutnTextEditor"]; !exists {
		other.PrintError("Failed to retrieve text editor from libmuttoncfg.json\n\nPlease specify one with \"mutn tweak\"", back.ErrorRead)
	} else {
		editor = value.(string)
	}

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
	err = cmd.Run()
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
	}
	return nil, false
}
