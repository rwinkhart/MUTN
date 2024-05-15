package sync

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"golang.org/x/crypto/ssh"
	"os"
	"strconv"
	"strings"
)

// getSSHOutput runs a command over SSH and returns the output
// currently only supports password-less key-based authentication TODO add password support, still require key
func getSSHOutput(cmd string, manualSync bool) string {
	// get SSH config info, exit if not configured (displaying an error if the sync job was called manually)
	var sshUserIPPortIdentity []string
	if manualSync {
		sshUserIPPortIdentity = backend.ReadConfig([]string{"sshUser", "sshIP", "sshPort", "sshIdentity"}, "SSH settings not configured - run \"mutn init\" to configure")
	} else {
		sshUserIPPortIdentity = backend.ReadConfig([]string{"sshUser", "sshIP", "sshPort", "sshIdentity"}, "0")
	}

	var user, ip, port, identity string
	for i, key := range sshUserIPPortIdentity {
		switch i {
		case 0:
			user = key
		case 1:
			ip = key
		case 2:
			port = key
		case 3:
			identity = key
		}
	}

	// read and parse private key
	key, err := os.ReadFile(identity)
	if err != nil {
		fmt.Println(backend.AnsiError+"Sync failed - unable to read private key file:", identity+backend.AnsiReset)
		os.Exit(1)
	}
	parsedKey, err := ssh.ParsePrivateKey(key)
	if err != nil {
		fmt.Println(backend.AnsiError+"Sync failed - Unable to parse private key:", identity+backend.AnsiReset)
		os.Exit(1)
	}

	// configure SSH client
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(parsedKey),
		},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey() // TODO remove

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

// getRemoteData returns lists of remote entries, mod times, folders, and deletions (four separate lists)
func getRemoteData(manualSync bool) ([]string, []int64, []string, []string) {
	// get remote output over SSH
	output := getSSHOutput("ls && uname -r", manualSync) // TODO replace cmd with remote server command (e.g. "libmuttonserver fetch")

	// split output into slice based on occurrences of "\x1f"
	outputSlice := strings.Split(output, "\x1f")

	// re-form the lists
	entries := strings.Split(outputSlice[0], "\n")
	modsStrings := strings.Split(outputSlice[1], "\n")
	folders := strings.Split(outputSlice[2], "\n")
	deletions := strings.Split(outputSlice[3], "\n")

	// convert the mod times to int64
	var mods []int64
	for _, modString := range modsStrings {
		mod, _ := strconv.ParseInt(modString, 10, 64)
		mods = append(mods, mod)
	}

	return entries, mods, folders, deletions
}

// getLocalData returns lists of local entries and mod times (two separate lists)
func getLocalData() ([]string, []int64) {
	// get a list of all entries
	fileList, _ := WalkEntryDir()

	// get a list of all entry modification times
	var modList []int64
	for _, file := range fileList {
		modTime, _ := os.Stat(backend.EntryRoot + file)
		modList = append(modList, modTime.ModTime().Unix())
	}

	// return the lists
	return fileList, modList
}

// RunJob runs the SSH sync job
func RunJob(manualSync bool) {
	// TODO fetch remote lists
	remoteEntries, remoteMods, remoteFolders, remoteDeletions := getRemoteData(manualSync)
	fmt.Println(remoteEntries, remoteMods, remoteFolders, remoteDeletions) // TODO placeholder

	// TODO sync deletions and folders

	// fetch local lists
	localEntries, localMods := getLocalData()
	fmt.Println(localEntries, localMods) // TODO placeholder

	// TODO sort both local and remote entry/mod lists to ensure common entries are listed first

	// TODO sync new and updated entries

	// exit program after successful sync
	os.Exit(0)
}
