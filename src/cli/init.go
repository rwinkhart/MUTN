package cli

import (
	"bytes"
	"cmp"
	"fmt"
	"os"
	"strconv"

	"github.com/rwinkhart/go-boilerplate/back"
	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/libmutton/core"
	"github.com/rwinkhart/libmutton/sync"
)

// TempInitCli initializes the MUTN environment based on user input.
func TempInitCli() {
	// text editor
	textEditor := cmp.Or(front.Input("Text editor (leave blank for $EDITOR, falls back to \""+fallbackEditor+"\"):"), os.Getenv("EDITOR"), fallbackEditor)

	// SSH info
	configSSH := front.InputBinary("Configure SSH settings (for synchronization)?")
	if configSSH {
		// necessary SSH info
		fmt.Println(AnsiBold + "\nNote:" + back.AnsiReset + " Only key-based authentication is supported (keys may optionally be passphrase-protected).\nThe remote server must already be in your ~" + core.PathSeparator + ".ssh" + core.PathSeparator + "known_hosts file.")
		sshUser := front.Input("Remote SSH username:")
		sshPort := front.Input("Remote SSH port:")
		sshIP := front.Input("Remote SSH IP/domain:")

		// prompt for and ensure existence of SSH identity file
		var sshKey string
		var sshKeyIsFile bool
		for !sshKeyIsFile {
			fallbackSSHKey := back.Home + core.PathSeparator + ".ssh" + core.PathSeparator + "id_ed25519"
			sshKey = cmp.Or(back.ExpandPathWithHome(front.Input("SSH private identity file path (falls back to \""+fallbackSSHKey+"\"):")), fallbackSSHKey)
			sshKeyIsFile, _ = back.TargetIsFile(sshKey, false, 0)
			if !sshKeyIsFile {
				fmt.Println(back.AnsiError+"SSH identity file not found:", sshKey+back.AnsiReset) // do not exit after error (allow user to retry)
			}
		}

		sshKeyProtected := front.InputBinary("Is the identity file password-protected?")

		// initialize libmutton directories
		oldDeviceID := core.DirInit(false)

		// write config file (temporarily assigns sshEntryRoot and sshIsWindows to null to pass initial device ID registration)
		core.WriteConfig([][3]string{{"MUTN", "textEditor", textEditor}, {"LIBMUTTON", "sshUser", sshUser}, {"LIBMUTTON", "sshIP", sshIP}, {"LIBMUTTON", "sshPort", sshPort}, {"LIBMUTTON", "sshKey", sshKey}, {"LIBMUTTON", "sshKeyProtected", strconv.FormatBool(sshKeyProtected)}, {"LIBMUTTON", "sshEntryRoot", "null"}, {"LIBMUTTON", "sshIsWindows", "false"}}, nil, false)

		// generate and register device ID
		sshEntryRoot, sshIsWindows := sync.DeviceIDGen(oldDeviceID)

		// update config file with sshEntryRoot and sshIsWindows
		core.WriteConfig([][3]string{{"LIBMUTTON", "sshEntryRoot", sshEntryRoot}, {"LIBMUTTON", "sshIsWindows", sshIsWindows}}, nil, true)
	} else {
		// initialize libmutton directories
		core.DirInit(false)

		// write config file
		core.WriteConfig([][3]string{{"MUTN", "textEditor", textEditor}}, nil, false)
	}
	// RCW sanity check file
	var passphrase []byte
	for {
		passphrase = front.InputHidden("Master passphrase:")
		if !bytes.Equal(passphrase, front.InputHidden("Confirm master passphrase:")) {
			fmt.Println(back.AnsiError + "Passphrases do not match" + back.AnsiReset)
			continue
		}
		core.RCWSanityCheckGen(passphrase)
		break
	}
}
