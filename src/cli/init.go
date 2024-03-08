package cli

import (
	"github.com/rwinkhart/MUTN/src/offline"
	"os"
	"os/exec"
)

func TempInitCli() {
	// gpgID
	cmd := exec.Command("gpg", "-k")
	cmd.Stdout = os.Stdout
	cmd.Run()
	gpgID := input("GPG key ID:")

	// textEditor
	textEditor := input("Text editor (leave blank for $EDITOR, falls back to \"vi\"):")

	offline.TempInit(map[string]string{"textEditor": textEditor, "gpgID": gpgID})
}
