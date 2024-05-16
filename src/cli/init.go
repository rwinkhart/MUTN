package cli

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"math/rand"
	"os"
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
		// client device ID
		deviceIDPrefix, _ := os.Hostname()
		deviceIDSuffix := backend.StringGen(rand.Intn(48)+48, true, 0.2)
		deviceID := deviceIDPrefix + "-" + deviceIDSuffix
		os.Create(backend.ConfigDir + backend.PathSeparator + "devices" + backend.PathSeparator + deviceID) // TODO remove existing device ID file if it exists, copy device ID to server

		// necessary SSH info
		sshUser := input("Remote SSH username:")
		sshIP := input("Remote SSH IP address:")
		sshPort := input("Remote SSH port:")
		sshIdentity := input("SSH private identity file path:") // TODO implement generator and selector

		// write config file
		backend.TempInit(map[string]string{"textEditor": textEditor, "gpgID": gpgID, "sshUser": sshUser, "sshIP": sshIP, "sshPort": sshPort, "sshIdentity": sshIdentity})
	} else {
		// write config file
		backend.TempInit(map[string]string{"textEditor": textEditor, "gpgID": gpgID})
	}

}
