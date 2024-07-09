package sync

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// DeviceIDGen generates a new client device ID and registers it with the server
// device IDs are only needed for online synchronization
// device IDs are guaranteed unique as the current UNIX time is appended to them
// returns the remote EntryRoot and OS type (OS type is a bool: backend.IsWindows)
func DeviceIDGen() (string, string) {
	deviceIDPrefix, _ := os.Hostname()
	deviceIDSuffix := backend.StringGen(rand.Intn(32)+48, false, 0) + "-" + strconv.FormatInt(time.Now().Unix(), 10) // TODO consider using complex string generator and removing unsafe characters manually
	deviceID := deviceIDPrefix + "-" + deviceIDSuffix
	_, err := os.Create(backend.ConfigDir + backend.PathSeparator + "devices" + backend.PathSeparator + deviceID) // TODO remove existing device ID file if it exists (from both client and server)
	if err != nil {
		fmt.Println(backend.AnsiError + "Failed to create local device ID file: " + err.Error() + backend.AnsiReset)
	}

	// register device ID with server and fetch remote EntryRoot and OS type
	//manualSync is true so the user is alerted if device ID registration fails
	sshEntryRootSSHIsWindows := strings.Split(GetSSHOutput("libmuttonserver register", deviceID, true), "\x1e")

	return sshEntryRootSSHIsWindows[0], sshEntryRootSSHIsWindows[1]
}
