package cli

import (
	"fmt"

	"github.com/rwinkhart/go-boilerplate/back"
	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
	"github.com/rwinkhart/go-boilerplate/security"
	"github.com/rwinkhart/libmutton/core"
)

// inputPasswordGen prompts the user for password generation parameters and returns a generated password as a string.
func inputPasswordGen() []byte {
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
	return security.BytesGen(passLength, complexity, passCharset)
}

// writeEntryCLI writes decSlice to realPath and previews it (errors if no data is supplied).
func writeEntryCLI(realPath string, decSlice []string, passwordIsNew bool, oldPassword []byte) {
	if core.EntryIsNotEmpty(decSlice) {
		if err := core.WriteEntry(realPath, decSlice, passwordIsNew, nil); err != nil {
			other.PrintError("Failed to write entry: "+err.Error(), back.ErrorWrite)
		}
		// preview the entry
		CopyMenu("", decSlice, oldPassword)
	} else {
		other.PrintError("No data supplied for entry", back.ErrorTargetNotFound)
	}
}
