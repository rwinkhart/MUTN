package cli

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"github.com/rwinkhart/MUTN/src/sync"
	"os"
	"strconv"
)

// TempInitCli initializes the MUTN environment based on user input
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
	textEditor := input("Text editor (leave blank for $EDITOR, falls back to \"" + backend.FallbackEditor + "\"):")

	// SSH info
	configSSH := inputBinary("Configure SSH settings (for synchronization)?")
	if configSSH {
		// necessary SSH info
		fmt.Println(AnsiBold + "\nNote:" + backend.AnsiReset + " Only key-based authentication is supported (keys may optionally be passphrase-protected).\nThe remote server must already be in your ~/.ssh/known_hosts file.")
		sshUser := input("Remote SSH username:")
		sshIP := input("Remote SSH IP address:")
		sshPort := input("Remote SSH port:")
		sshKey := input("SSH private identity file path:") // TODO implement generator and selector
		sshKeyProtected := inputBinary("Is the identity file password-protected?")

		// write config file
		backend.TempInit(map[string]string{"textEditor": textEditor, "gpgID": gpgID, "sshUser": sshUser, "sshIP": sshIP, "sshPort": sshPort, "sshKey": sshKey, "sshKeyProtected": strconv.FormatBool(sshKeyProtected)}, false)

		// generate device ID
		sshEntryRoot, sshIsWindows := sync.DeviceIDGen()

		// update config file with sshEntryRoot and sshIsWindows TODO append to existing config file
		backend.TempInit(map[string]string{"textEditor": textEditor, "gpgID": gpgID, "sshUser": sshUser, "sshIP": sshIP, "sshPort": sshPort, "sshKey": sshKey, "sshKeyProtected": strconv.FormatBool(sshKeyProtected), "sshEntryRoot": sshEntryRoot, "sshIsWindows": sshIsWindows}, true)
	} else {
		// write config file
		backend.TempInit(map[string]string{"textEditor": textEditor, "gpgID": gpgID}, false)
	}

}
