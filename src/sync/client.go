package sync

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"github.com/rwinkhart/MUTN/src/cli"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"os"
	"strconv"
	"strings"
)

// global constants used only in this file
const (
	ansiDelete   = "\033[38;5;1m"
	ansiDownload = "\033[38;5;2m"
	ansiUpload   = "\033[38;5;4m"
)

// GetSSHOutput runs a command over SSH and returns the output
// only supports key-based authentication (passphrase-protected keys are supported in a CLI environment) TODO create a more generic interface for passphrase input
func GetSSHOutput(cmd string, manualSync bool) string {
	// get SSH config info, exit if not configured (displaying an error if the sync job was called manually)
	var sshUserConfig []string
	if manualSync {
		sshUserConfig = backend.ReadConfig([]string{"sshUser", "sshIP", "sshPort", "sshKey", "sshKeyProtected"}, "SSH settings not configured - run \"mutn init\" to configure")
	} else {
		sshUserConfig = backend.ReadConfig([]string{"sshUser", "sshIP", "sshPort", "sshKey", "sshKeyProtected"}, "0")
	}

	var user, ip, port, keyFile, keyFileProtected string
	for i, key := range sshUserConfig {
		switch i {
		case 0:
			user = key
		case 1:
			ip = key
		case 2:
			port = key
		case 3:
			keyFile = key
		case 4:
			keyFileProtected = key
		}
	}

	// read private key
	key, err := os.ReadFile(keyFile)
	if err != nil {
		fmt.Println(backend.AnsiError+"Sync failed - unable to read private key file:", keyFile+backend.AnsiReset)
		os.Exit(1)
	}

	// parse private key
	var parsedKey ssh.Signer
	if keyFileProtected != "true" {
		parsedKey, err = ssh.ParsePrivateKey(key)
	} else {
		parsedKey, err = ssh.ParsePrivateKeyWithPassphrase(key, []byte(cli.InputHidden("Enter passphrase for \""+keyFile+"\":"))) // TODO test passphrase-protected keys
	}
	if err != nil {
		fmt.Println(backend.AnsiError+"Sync failed - Unable to parse private key:", keyFile+backend.AnsiReset)
		os.Exit(1)
	}

	// read known hosts file
	hostKeyCallback, err := knownhosts.New(backend.Home + backend.PathSeparator + ".ssh" + backend.PathSeparator + "known_hosts")
	if err != nil {
		fmt.Println(backend.AnsiError + "Sync failed - Unable to read known hosts file:" + err.Error() + backend.AnsiReset)
		os.Exit(1)
	}

	// configure SSH client
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(parsedKey),
		},
		HostKeyCallback: hostKeyCallback,
	}

	// connect to SSH server
	sshClient, err := ssh.Dial("tcp", ip+":"+port, sshConfig)
	if err != nil {
		fmt.Println(backend.AnsiError+"Sync failed - Unable to connect to remote server:", err.Error()+backend.AnsiReset)
		os.Exit(1)
	}
	defer sshClient.Close()

	// create a session
	sshSession, err := sshClient.NewSession()
	if err != nil {
		fmt.Println(backend.AnsiError+"Sync failed - Unable to establish SSH session:", err.Error()+backend.AnsiReset)
		os.Exit(1)
	}

	// run the provided command
	output, err := sshSession.CombinedOutput(cmd)
	if err != nil {
		fmt.Println(backend.AnsiError+"Sync failed - Unable to run SSH command:", err.Error()+backend.AnsiReset)
		os.Exit(1)
	}

	// convert the output to a string and remove leading/trailing whitespace
	outputString := string(output)
	outputString = strings.TrimSpace(outputString)

	return outputString
}

// getRemoteDataFromClient returns a map of remote entries to their modification times, a list of remote folders, and a list of queued deletions
func getRemoteDataFromClient(manualSync bool) (map[string]int64, []string, []string) {
	// get remote output over SSH
	clientDeviceID, _ := os.ReadDir(backend.ConfigDir + backend.PathSeparator + "devices")
	output := GetSSHOutput("libmuttonserver fetch "+clientDeviceID[0].Name(), manualSync)

	// split output into slice based on occurrences of "\x1d"
	outputSlice := strings.Split(output, "\x1d")

	// re-form the lists TODO handle error for index out of bounds (occurs if reading deletions directory on server fails)
	entries := strings.Split(outputSlice[0], "\x1f")[1:]
	modsStrings := strings.Split(outputSlice[1], "\x1f")[1:]
	folders := strings.Split(outputSlice[2], "\x1f")[1:]
	deletions := strings.Split(outputSlice[3], "\x1f")[1:]

	// convert the mod times to int64
	var mods []int64
	for _, modString := range modsStrings {
		mod, _ := strconv.ParseInt(modString, 10, 64)
		mods = append(mods, mod)
	}

	// map remote entries to their modification times
	entryModMap := make(map[string]int64)
	for i, entry := range entries {
		entryModMap[entry] = mods[i]
	}

	return entryModMap, folders, deletions
}

// getLocalData returns a map of local entries to their modification times
func getLocalData() map[string]int64 {
	// get a list of all entries
	entries, _ := WalkEntryDir()

	// get a list of all entry modification times
	modList := getModTimes(entries)

	// map the entries to their modification times
	entryModMap := make(map[string]int64)
	for i, entry := range entries {
		entryModMap[entry] = modList[i]
	}

	// return the lists
	return entryModMap
}

// syncLists syncs entries between the client and server based on modification times
// using maps means that syncing will be done in an arbitrary order, but it is a worthy tradeoff for speed and simplicity
func syncLists(localEntryModMap, remoteEntryModMap map[string]int64) {
	// iterate over client entries
	for entry, localModTime := range localEntryModMap {
		// check if the entry is present in the server map
		if remoteModTime, present := remoteEntryModMap[entry]; present {
			// entry exists on both client and server, compare mod times
			if remoteModTime > localModTime {
				fmt.Println(ansiDownload+entry+backend.AnsiReset, "is newer on server, downloading...")
				// TODO entry is newer on server, download
			} else if remoteModTime < localModTime {
				fmt.Println(ansiUpload+entry+backend.AnsiReset, "is newer on client, uploading...")
				// TODO entry is newer on client, upload
			}
			// remove entry from remoteEntryModMap (process of elimination)
			delete(remoteEntryModMap, entry)
		} else {
			fmt.Println(ansiUpload+entry+backend.AnsiReset, "does not exist on server, uploading...")
			// TODO entry does not exist on server, upload
		}
	}

	// iterate over remaining entries in remoteEntryModMap
	for entry := range remoteEntryModMap {
		fmt.Println(ansiDownload+entry+backend.AnsiReset, "does not exist on client, downloading...")
		// TODO entry does not exist on client, download
	}
}

// deletionSync removes entries from the client that have been deleted on the server (multi-client deletion)
func deletionSync(deletions []string) {
	for _, deletion := range deletions {
		fmt.Println(ansiDelete+deletion+backend.AnsiReset, "has been sheared, removing...")
		os.RemoveAll(backend.EntryRoot + deletion)
	}
}

// folderSync creates folders on the client (from the given list of folder names)
func folderSync(folders []string) {
	for _, folder := range folders {
		// check if folder already exists
		isFile, isAccessible := backend.TargetIsFile(backend.EntryRoot+folder, false, 1)

		if !isFile && !isAccessible {
			os.MkdirAll(backend.EntryRoot+folder, 0700)
		} else if isFile {
			fmt.Println(backend.AnsiError + "Sync failed - Failed to create folder \"" + folder + "\" - a file with the same name already exists" + backend.AnsiReset)
			os.Exit(1)
		}
	}
}

// RunJob runs the SSH sync job
func RunJob(manualSync bool) {
	// fetch remote lists
	remoteEntryModMap, remoteFolders, deletions := getRemoteDataFromClient(manualSync)

	// sync folders
	folderSync(remoteFolders)

	// sync deletions
	deletionSync(deletions)

	// fetch local lists
	localEntryModMap := getLocalData()

	// sync new and updated entries
	syncLists(localEntryModMap, remoteEntryModMap)

	// exit program after successful sync
	os.Exit(0)
}
