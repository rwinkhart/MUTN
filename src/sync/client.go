package sync

import (
	"context"
	"fmt"
	"github.com/bramvdbogaerde/go-scp"
	"github.com/rwinkhart/MUTN/src/backend"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"os"
	"strconv"
	"strings"
)

// TODO it is likely that all of this breaks on Windows, since the server returns unpredictable paths
// See what libmuttonserver fetch returns on Windows
// If the path separators are different, then we'll have to save the server OS during init and adjust accordingly
// I would prefer not to unnecessarily replace the path separators if syncing from Unix->Unix or Windows->Windows

// global constants used only in this file
const (
	ansiDelete   = "\033[38;5;1m"
	ansiDownload = "\033[38;5;2m"
	ansiUpload   = "\033[38;5;4m"
)

// getSSHClient returns an SSH client connection to the server (also returns the remote username as a string)
// only supports key-based authentication (passphrases are supported for CLI-based implementations)
func getSSHClient(manualSync bool) (*ssh.Client, string) {
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
		parsedKey, err = ssh.ParsePrivateKeyWithPassphrase(key, inputKeyFilePassphrase()) // TODO test passphrase-protected keys
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

	return sshClient, user
}

// GetSSHOutput runs a command over SSH and returns the output as a string
func GetSSHOutput(cmd string, manualSync bool) string {
	sshClient, _ := getSSHClient(manualSync)
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

// scpTransfer uploads or downloads an entry over SCP // TODO consider using goroutines for multithreaded syncs
// WARNING: does not close sshClient; it is left open for further operations
// TODO does not yet support modification time preservation, needs https://github.com/bramvdbogaerde/go-scp/pull/81/commits
func scpTransfer(sshClient *ssh.Client, entryName, sshUser string, download bool) {
	// create an SCP client TODO see if one can be passed in and reused
	scpClient, err := scp.NewClientBySSH(sshClient)

	if err != nil {
		fmt.Println(backend.AnsiError+"Sync failed - Unable to establish SCP session:", err.Error()+backend.AnsiReset)
		os.Exit(1)
	}

	// upload or download the entry
	if download {
		f, _ := os.Create(backend.EntryRoot + entryName)
		defer f.Close()
		err = scpClient.CopyFromRemote(context.Background(), f, backend.PathSeparator+"home"+backend.PathSeparator+sshUser+bareEntryRoot+entryName) // TODO support remote home directories other than "/home/$USER" (maybe fetch home directory form server during init, transparently to the user)
		if err != nil {
			fmt.Println(backend.AnsiError+"Sync failed - Unable to download entry:", entryName, err.Error()+backend.AnsiReset)
			os.Exit(1)

		}
	} else {
		f, _ := os.Open(backend.EntryRoot + entryName)
		defer f.Close()
		err = scpClient.CopyFromFile(context.Background(), *f, backend.PathSeparator+"home"+backend.PathSeparator+sshUser+bareEntryRoot+entryName, "0600")
	}
}

// getRemoteDataFromClient returns a map of remote entries to their modification times, a list of remote folders, and a list of queued deletions
func getRemoteDataFromClient(manualSync bool) (map[string]int64, []string, []string) {
	// get remote output over SSH
	clientDeviceID, _ := os.ReadDir(backend.ConfigDir + backend.PathSeparator + "devices")
	if len(clientDeviceID) == 0 {
		fmt.Println(backend.AnsiError + "Sync failed - No device ID found; run \"mutn init\" to generate a device ID" + backend.AnsiReset)
		os.Exit(1)
	}
	output := GetSSHOutput("libmuttonserver fetch "+clientDeviceID[0].Name(), manualSync)

	// split output into slice based on occurrences of "\x1d"
	outputSlice := strings.Split(output, "\x1d")

	// re-form the lists
	if len(outputSlice) != 4 { // ensure information from server is complete
		fmt.Println(backend.AnsiError + "Sync failed - Unable to fetch remote data; server returned an unexpected response" + backend.AnsiReset)
		os.Exit(1)
	}
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
func syncLists(localEntryModMap, remoteEntryModMap map[string]int64, manualSync bool) {
	// establish an SSH connection for transfers TODO only establish if needed
	sshClient, sshUser := getSSHClient(manualSync)

	// iterate over client entries
	for entry, localModTime := range localEntryModMap {
		// check if the entry is present in the server map
		if remoteModTime, present := remoteEntryModMap[entry]; present {
			// entry exists on both client and server, compare mod times
			if remoteModTime > localModTime {
				fmt.Println(ansiDownload+entry+backend.AnsiReset, "is newer on server, downloading...")
				scpTransfer(sshClient, entry, sshUser, true)
			} else if remoteModTime < localModTime {
				fmt.Println(ansiUpload+entry+backend.AnsiReset, "is newer on client, uploading...")
				scpTransfer(sshClient, entry, sshUser, false)
			}
			// remove entry from remoteEntryModMap (process of elimination)
			delete(remoteEntryModMap, entry)
		} else {
			fmt.Println(ansiUpload+entry+backend.AnsiReset, "does not exist on server, uploading...")
			scpTransfer(sshClient, entry, sshUser, false)
		}
	}

	// iterate over remaining entries in remoteEntryModMap
	for entry := range remoteEntryModMap {
		fmt.Println(ansiDownload+entry+backend.AnsiReset, "does not exist on client, downloading...")
		scpTransfer(sshClient, entry, sshUser, true)
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
	syncLists(localEntryModMap, remoteEntryModMap, manualSync)

	// exit program after successful sync
	os.Exit(0)
}