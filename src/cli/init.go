package cli

import (
	"cmp"
	"fmt"
	"os"
	"strconv"

	"github.com/rwinkhart/libmutton/core"
	"github.com/rwinkhart/libmutton/sync"
)

// TempInitCli initializes the MUTN environment based on user input.
func TempInitCli() {
	// gpgID
	var gpgID string
	if inputBinary("Auto-generate GPG key?") {
		gpgID = core.GpgKeyGen()
	} else {
		// select GPG key from menu
		uidSlice := core.GpgUIDListGen()
		gpgIDInt := inputMenuGen("Select GPG key:", uidSlice)
		if gpgIDInt == 0 {
			fmt.Println(core.AnsiError + "No GPG keys found - please generate one" + core.AnsiReset)
			os.Exit(core.ErrorTargetNotFound)
		}
		gpgID = uidSlice[gpgIDInt-1]
	}

	// textEditor
	textEditor := cmp.Or(input("Text editor (leave blank for $EDITOR, falls back to \""+fallbackEditor+"\"):"), os.Getenv("EDITOR"), fallbackEditor)

	// SSH info
	configSSH := inputBinary("Configure SSH settings (for synchronization)?")
	if configSSH {
		// necessary SSH info
		fmt.Println(AnsiBold + "\nNote:" + core.AnsiReset + " Only key-based authentication is supported (keys may optionally be passphrase-protected).\nThe remote server must already be in your ~" + core.PathSeparator + ".ssh" + core.PathSeparator + "known_hosts file.")
		sshUser := input("Remote SSH username:")
		sshPort := input("Remote SSH port:")
		sshIP := input("Remote SSH IP/domain:")

		// prompt for and ensure existence of SSH identity file
		var sshKey string
		var sshKeyIsFile bool
		for !sshKeyIsFile {
			fallbackSSHKey := core.Home + core.PathSeparator + ".ssh" + core.PathSeparator + "id_ed25519"
			sshKey = cmp.Or(core.ExpandPathWithHome(input("SSH private identity file path (falls back to \""+fallbackSSHKey+"\"):")), fallbackSSHKey)
			sshKeyIsFile, _ = core.TargetIsFile(sshKey, false, 0)
			if !sshKeyIsFile {
				fmt.Println(core.AnsiError+"SSH identity file not found:", sshKey+core.AnsiReset)
			}
		}

		sshKeyProtected := inputBinary("Is the identity file password-protected?")

		// initialize libmutton directories
		oldDeviceID := core.DirInit(false)

		// write config file (temporarily assigns sshEntryRoot and sshIsWindows to null to pass initial device ID registration)
		core.WriteConfig([][3]string{{"MUTN", "textEditor", textEditor}, {"LIBMUTTON", "gpgID", gpgID}, {"LIBMUTTON", "sshUser", sshUser}, {"LIBMUTTON", "sshIP", sshIP}, {"LIBMUTTON", "sshPort", sshPort}, {"LIBMUTTON", "sshKey", sshKey}, {"LIBMUTTON", "sshKeyProtected", strconv.FormatBool(sshKeyProtected)}, {"LIBMUTTON", "sshEntryRoot", "null"}, {"LIBMUTTON", "sshIsWindows", "false"}}, nil, false)

		// generate and register device ID
		sshEntryRoot, sshIsWindows := sync.DeviceIDGen(oldDeviceID)

		// update config file with sshEntryRoot and sshIsWindows
		core.WriteConfig([][3]string{{"LIBMUTTON", "sshEntryRoot", sshEntryRoot}, {"LIBMUTTON", "sshIsWindows", sshIsWindows}}, nil, true)
	} else {
		// initialize libmutton directories
		core.DirInit(false)

		// write config file
		core.WriteConfig([][3]string{{"MUTN", "textEditor", textEditor}, {"LIBMUTTON", "gpgID", gpgID}}, nil, false)
	}
}
