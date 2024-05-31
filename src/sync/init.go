package sync

import (
	"github.com/rwinkhart/MUTN/src/backend"
	"math/rand"
	"os"
	"strings"
)

// DeviceIDGen generates a new client device ID and registers it with the server
// device IDs are only needed for online synchronization
// returns the remote EntryRoot and OS type (OS type is a bool: backend.IsWindows)
func DeviceIDGen() (string, string) {
	deviceIDPrefix, _ := os.Hostname()
	deviceIDSuffix := backend.StringGen(rand.Intn(48)+48, false, 0) // TODO consider using complex string generator and removing unsafe characters manually
	deviceID := deviceIDPrefix + "-" + deviceIDSuffix
	os.Create(backend.ConfigDir + backend.PathSeparator + "devices" + backend.PathSeparator + deviceID)         // TODO remove existing device ID file if it exists (from both client and server)
	sshEntryRootSSHIsWindows := strings.Split(GetSSHOutput("libmuttonserver register "+deviceID, true), "\x1d") // register device ID with server and fetch remote EntryRoot and OS type; manualSync is true so the user is alerted if device ID registration fails

	return sshEntryRootSSHIsWindows[0], sshEntryRootSSHIsWindows[1]
}
