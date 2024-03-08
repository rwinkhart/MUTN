package cli

import (
	"github.com/rwinkhart/MUTN/src/offline"
)

// TempInitCli initializes the MUTN environment based on user input
func TempInitCli() {
	// gpgID
	var gpgID string
	if inputBinary("Auto-generate GPG key?") {
		gpgID = offline.GpgKeyGen()
	} else {
		uidSlice := offline.GpgUIDListGen()
		gpgIDInt := inputMenuGen("Select GPG key:", uidSlice)
		gpgID = uidSlice[gpgIDInt-1]
	}

	// textEditor
	textEditor := input("Text editor (leave blank for $EDITOR, falls back to \"" + offline.FallbackEditor + "\"):")

	offline.TempInit(map[string]string{"textEditor": textEditor, "gpgID": gpgID})
}
