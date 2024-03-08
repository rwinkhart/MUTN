package cli

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/offline"
	"os"
)

// TempInitCli initializes the MUTN environment based on user input
func TempInitCli() {
	// gpgID
	var gpgID string
	if inputBinary("Auto-generate GPG key?") {
		gpgID = offline.GpgKeyGen()
	} else {
		// select GPG key from menu
		uidSlice := offline.GpgUIDListGen()
		gpgIDInt := inputMenuGen("Select GPG key:", uidSlice)
		if gpgIDInt == 0 {
			fmt.Println(offline.AnsiError + "No GPG keys found - please generate one" + offline.AnsiReset)
			os.Exit(1)
		}
		gpgID = uidSlice[gpgIDInt-1]
	}

	// textEditor
	textEditor := input("Text editor (leave blank for $EDITOR, falls back to \"" + offline.FallbackEditor + "\"):")

	offline.TempInit(map[string]string{"textEditor": textEditor, "gpgID": gpgID})
}
