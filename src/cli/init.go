package cli

import (
	"cmp"
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
	textEditor := cmp.Or(input("Text editor (leave blank for $EDITOR, falls back to \""+fallbackEditor+"\"):"), os.Getenv("EDITOR"), fallbackEditor)

	// SSH info
	configSSH := inputBinary("Configure SSH settings (for synchronization)?")
	if configSSH {
		// necessary SSH info
		fmt.Println(AnsiBold + "\nNote:" + backend.AnsiReset + " Only key-based authentication is supported (keys may optionally be passphrase-protected).\nThe remote server must already be in your ~" + backend.PathSeparator + ".ssh" + backend.PathSeparator + "known_hosts file.")
		sshUser := input("Remote SSH username:")
		sshPort := input("Remote SSH port:")
		sshIP := input("Remote SSH IP/domain:")

		// prompt for and ensure existence of SSH identity file
		var sshKey string
		var sshKeyIsFile bool
		for !sshKeyIsFile {
			fallbackSSHKey := backend.Home + backend.PathSeparator + ".ssh" + backend.PathSeparator + "id_ed25519"
			sshKey = cmp.Or(expandPathWithHome(input("SSH private identity file path (falls back to \""+fallbackSSHKey+"\"):")), fallbackSSHKey)
			sshKeyIsFile, _ = backend.TargetIsFile(sshKey, false, 0)
			if !sshKeyIsFile {
				fmt.Println(backend.AnsiError + "SSH identity file not found: " + sshKey + backend.AnsiReset)
			}
		}

		sshKeyProtected := inputBinary("Is the identity file password-protected?")

		// initialize libmutton directories
		backend.DirInit(false)

		// write config file (temporarily assigns sshEntryRoot and sshIsWindows to null to pass initial device ID registration)
		backend.WriteConfig([][3]string{{"MUTN", "textEditor", textEditor}, {"LIBMUTTON", "gpgID", gpgID}, {"LIBMUTTON", "sshUser", sshUser}, {"LIBMUTTON", "sshIP", sshIP}, {"LIBMUTTON", "sshPort", sshPort}, {"LIBMUTTON", "sshKey", sshKey}, {"LIBMUTTON", "sshKeyProtected", strconv.FormatBool(sshKeyProtected)}, {"LIBMUTTON", "sshEntryRoot", "null"}, {"LIBMUTTON", "sshIsWindows", "null"}}, false)

		// generate and register device ID
		sshEntryRoot, sshIsWindows := sync.DeviceIDGen()

		// update config file with sshEntryRoot and sshIsWindows
		backend.WriteConfig([][3]string{{"LIBMUTTON", "sshEntryRoot", sshEntryRoot}, {"LIBMUTTON", "sshIsWindows", sshIsWindows}}, true)
	} else {
		// initialize libmutton directories
		backend.DirInit(false)

		// write config file
		backend.WriteConfig([][3]string{{"MUTN", "textEditor", textEditor}, {"LIBMUTTON", "gpgID", gpgID}}, false)
	}
}
