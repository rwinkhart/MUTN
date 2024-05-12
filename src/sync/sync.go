package sync

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"os"
)

// returns lists of local entries and mod times (two separate lists)
func getLocalData() {
	fileList, _ := WalkEntryDir()
	fmt.Println(fileList) // TODO placeholder
}

// RunJob runs the SSH sync job
func RunJob(manualSync bool) {
	// get SSH config info, exit if not configured (displaying an error if the sync job was called manually)
	var sshUserIPPortIdentity []string
	if manualSync {
		sshUserIPPortIdentity = backend.ReadConfig([]string{"sshUser", "sshIP", "sshPort", "sshIdentity"}, "SSH settings not configured - run \"mutn init\" to configure")
	} else {
		sshUserIPPortIdentity = backend.ReadConfig([]string{"sshUser", "sshIP", "sshPort", "sshIdentity"}, "0")
	}

	var sshUser, sshIP, sshPort, sshIdentity string
	for i, key := range sshUserIPPortIdentity {
		switch i {
		case 0:
			sshUser = key
		case 1:
			sshIP = key
		case 2:
			sshPort = key
		case 3:
			sshIdentity = key
		}
	}

	fmt.Println(sshUser, sshIP, sshPort, sshIdentity) // TODO placeholder

	os.Exit(0)
}
