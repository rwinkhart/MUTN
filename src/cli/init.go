package cli

import (
	"github.com/rwinkhart/MUTN/src/offline"
)

func TempInitCli() {
	// gpgID
	//generateGPG := inputBinary("Auto-generate GPG key?")

	uidSlice := offline.GpgUIDListGen()
	gpgIDInt := inputMenuGen("Select GPG key:", uidSlice)
	gpgID := uidSlice[gpgIDInt-1]

	// textEditor
	textEditor := input("Text editor (leave blank for $EDITOR, falls back to \"vi\"):")

	offline.TempInit(map[string]string{"textEditor": textEditor, "gpgID": gpgID})
}