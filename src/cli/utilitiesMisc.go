package cli

import (
	"fmt"
	"strings"

	"github.com/rwinkhart/go-boilerplate/back"
	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
	"github.com/rwinkhart/libmutton/core"
	"github.com/rwinkhart/libmutton/synccycles"
)

// inputPasswordGen prompts the user for password generation parameters and returns a generated password as a string.
func inputPasswordGen() string {
	passLength := front.InputInt("Password length:", 1, -1)
	fmt.Println()
	passCharset := uint8(front.InputMenuGen("Password complexity:", []string{"Simple", "Complex", "Ultra Complex (not compatible with many services)"}))
	var complexity float64
	switch passCharset {
	case 1:
		complexity = 0 // simple
		// 1 indicates string gen for filenames, but since complexity is 0, only the base charset is used
	default:
		complexity = 0.2 // (ultra) complex
		// 2 and 3 indicate complex and ultra complex charsets, respectively
	}
	return synccycles.StringGen(passLength, complexity, passCharset)
}

// writeEntryCLI writes an entry to targetLocation and previews it (errors if no data is supplied).
func writeEntryCLI(targetLocation string, decryptedEntry []string, hideSecrets bool) {
	if core.EntryIsNotEmpty(decryptedEntry) {
		// write the entry to the target location
		err := core.WriteEntry(targetLocation, []byte(strings.Join(decryptedEntry, "\n")))
		if err != nil {
			other.PrintError("Failed to write entry: "+err.Error(), back.ErrorWrite, true)
		}
		// preview the entry
		fmt.Println(back.AnsiBold + "\nEntry Preview:" + back.AnsiReset)
		EntryReader(decryptedEntry, hideSecrets, true)
	} else {
		other.PrintError("No data supplied for entry", back.ErrorTargetNotFound, true)
	}
}
