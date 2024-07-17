package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rwinkhart/MUTN/src/backend"
	"github.com/rwinkhart/MUTN/src/sync"
)

// TempInitCli initializes the MUTN environment based on user input (will be replaced with a TUI menu)
func TempInitCli() {
	// gpgID
	var gpgID string
	if inputBinary("Auto-generate GPG key?") {
		gpgID = backend.GpgKeyGen()
	} else {
		// select GPG key from menu
		uidSlice := backend.GpgUIDListGen()
		gpgIDInt := inputMenuGen("Select GPG key:", uidSlice)
		if gpgIDInt == 0 {
			fmt.Println(backend.AnsiError + "No GPG keys found - please generate one" + backend.AnsiReset)
			os.Exit(1)
		}
		gpgID = uidSlice[gpgIDInt-1]
	}

	// textEditor
	textEditor := input("Text editor (leave blank for $EDITOR, falls back to \"" + fallbackEditor + "\"):")
	if textEditor == "" {
		textEditor = textEditorFallback()
	}

	// SSH info
	configSSH := inputBinary("Configure SSH settings (for synchronization)?")
	if configSSH {
		// necessary SSH info
		fmt.Println(AnsiBold + "\nNote:" + backend.AnsiReset + " Only key-based authentication is supported (keys may optionally be passphrase-protected).\nThe remote server must already be in your ~/.ssh/known_hosts file.")
		sshUser := input("Remote SSH username:")
		sshPort := input("Remote SSH port:")
		sshIP := input("Remote SSH IP/domain:")
		sshKey := expandPathWithHome(input("SSH private identity file path:")) // TODO implement generator and selector
		sshKeyProtected := inputBinary("Is the identity file password-protected?")

		// initialize libmutton directories
		backend.DirInit(false)

		// write config file (temporarily assigns sshEntryRoot and sshIsWindows to null to pass initial device ID registration)
		backend.WriteConfig(map[string]string{"textEditor": textEditor, "gpgID": gpgID, "sshUser": sshUser, "sshIP": sshIP, "sshPort": sshPort, "sshKey": sshKey, "sshKeyProtected": strconv.FormatBool(sshKeyProtected), "sshEntryRoot": "null", "sshIsWindows": "null"}, false)

		// generate and register device ID
		sshEntryRoot, sshIsWindows := sync.DeviceIDGen()

		// update config file with sshEntryRoot and sshIsWindows
		backend.WriteConfig(map[string]string{"sshEntryRoot": sshEntryRoot, "sshIsWindows": sshIsWindows}, true)
	} else {
		// initialize libmutton directories
		backend.DirInit(false)

		// write config file
		backend.WriteConfig(map[string]string{"textEditor": textEditor, "gpgID": gpgID}, false)
	}
}
