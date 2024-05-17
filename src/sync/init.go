package sync

import (
	"github.com/rwinkhart/MUTN/src/backend"
	"math/rand"
	"os"
)

// DeviceIDGen generates a new client device ID and registers it with the server
// device IDs are only needed for online synchronization
func DeviceIDGen() {
	deviceIDPrefix, _ := os.Hostname()
	deviceIDSuffix := backend.StringGen(rand.Intn(48)+48, false, 0) // TODO consider using complex string generator and removing unsafe characters manually
	deviceID := deviceIDPrefix + "-" + deviceIDSuffix
	os.Create(backend.ConfigDir + backend.PathSeparator + "devices" + backend.PathSeparator + deviceID) // TODO remove existing device ID file if it exists (from both client and server)
	GetSSHOutput("libmuttonserver register "+deviceID, false)                                           // register device ID with server
}
