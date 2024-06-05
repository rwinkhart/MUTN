package sync

import (
	"fmt"
	"github.com/pkg/sftp"
	"github.com/rwinkhart/MUTN/src/backend"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"os"
	"strconv"
	"strings"
	"time"
)

// global constants used only in this file
const (
	ansiDelete   = "\033[38;5;1m"
	ansiDownload = "\033[38;5;2m"
	ansiUpload   = "\033[38;5;4m"
)

// getSSHClient returns an SSH client connection to the server (also returns the remote EntryRoot as a string and the server's OS as a bool - IsWindows)
// only supports key-based authentication (passphrases are supported for CLI-based implementations)
func getSSHClient(manualSync bool) (*ssh.Client, string, bool) {
	// get SSH config info, exit if not configured (displaying an error if the sync job was called manually)
	var sshUserConfig []string
	var missingValueError string
	if manualSync {
		missingValueError = "SSH settings not configured - run \"mutn init\" to configure"
	} else {
		missingValueError = "0"
	}
	sshUserConfig = backend.ParseConfig([]string{"sshUser", "sshIP", "sshPort", "sshKey", "sshKeyProtected", "sshEntryRoot", "sshIsWindows"}, missingValueError)

	var user, ip, port, keyFile, keyFileProtected, entryRoot string
	var isWindows bool
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
		case 5:
			entryRoot = key
		case 6:
			isWindows, _ = strconv.ParseBool(key)
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
		parsedKey, err = ssh.ParsePrivateKeyWithPassphrase(key, inputKeyFilePassphrase())
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

	return sshClient, entryRoot, isWindows
}

// GetSSHOutput runs a command over SSH and returns the output as a string
// TODO run getSSHClient() only ONCE (from RunJob) - this saves re-creating the client AND prevents prompting for keyfile passphrase multiple times
func GetSSHOutput(cmd, stdin string, manualSync bool) string {
	sshClient, _, _ := getSSHClient(manualSync)
	defer sshClient.Close()

	// create a session
	sshSession, err := sshClient.NewSession()
	if err != nil {
		fmt.Println(backend.AnsiError+"Sync failed - Unable to establish SSH session:", err.Error()+backend.AnsiReset)
		os.Exit(1)
	}

	// provide stdin data for session
	sshSession.Stdin = strings.NewReader(stdin)

	// run the provided command
	var output []byte
	output, err = sshSession.CombinedOutput(cmd)
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
	if len(clientDeviceID) == 0 {
		if manualSync {
			fmt.Println(backend.AnsiError + "Sync failed - No device ID found; run \"mutn init\" to generate a device ID" + backend.AnsiReset)
			os.Exit(1)
		} else {
			os.Exit(0) // exit silently if the sync job was called automatically, as the user may just be in offline mode
		}
	}
	output := GetSSHOutput("libmuttonserver fetch", clientDeviceID[0].Name(), manualSync)

	// split output into slice based on occurrences of "\x1e"
	outputSlice := strings.Split(output, "\x1e")

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

// targetLocationFormatSFTP formats the target location to match the remote server's entry directory and path separator
func targetLocationFormatSFTP(targetName, serverEntryRoot string, serverIsWindows bool) string {
	if !serverIsWindows {
		return serverEntryRoot + targetName
	} else {
		return serverEntryRoot + strings.ReplaceAll(targetName, "/", "\\")
	}
}

// sftpSync takes two slices of entries (one for downloads and one for uploads) and syncs them between the client and server using SFTP
// TODO test Windows server hosting support
func sftpSync(downloadList, uploadList []string, manualSync bool) {
	// establish an SSH connection for transfers
	sshClient, sshEntryRoot, sshIsWindows := getSSHClient(manualSync)
	defer sshClient.Close()

	// create an SFTP client
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		fmt.Println(backend.AnsiError+"Sync failed - Unable to establish SFTP session:", err.Error()+backend.AnsiReset)
		os.Exit(1)
	}
	defer sftpClient.Close()

	// iterate over the download list
	var filesTransfered bool
	for _, entryName := range downloadList {
		filesTransfered = true // set a flag to indicate that files have been downloaded (used to determine whether to print a gap between download and upload messages)

		fmt.Println("Downloading " + ansiDownload + entryName + backend.AnsiReset)

		// store path to remote entry
		remoteEntryFullPath := targetLocationFormatSFTP(entryName, sshEntryRoot, sshIsWindows)

		// save modification time of remote file
		var fileInfo os.FileInfo
		fileInfo, err = sftpClient.Stat(remoteEntryFullPath)
		if err != nil {
			fmt.Println(backend.AnsiError+"Sync failed - Unable to get remote file info (modtime):", err.Error()+backend.AnsiReset)
			os.Exit(1)
		}
		modTime := fileInfo.ModTime()

		// open remote file
		var remoteFile *sftp.File
		remoteFile, err = sftpClient.Open(remoteEntryFullPath)
		if err != nil {
			fmt.Println(backend.AnsiError+"Sync failed - Unable to open remote file:", err.Error()+backend.AnsiReset)
			os.Exit(1)
		}

		// store path to local entry
		localEntryFullPath := backend.TargetLocationFormat(entryName)

		// create local file
		var localFile *os.File
		localFile, err = os.Create(localEntryFullPath)
		if err != nil {
			fmt.Println(backend.AnsiError+"Sync failed - Unable to create local file:", err.Error()+backend.AnsiReset)
			os.Exit(1)
		}

		// download the file
		_, err = remoteFile.WriteTo(localFile)
		if err != nil {
			fmt.Println(backend.AnsiError+"Sync failed - Unable to download remote file:", err.Error()+backend.AnsiReset)
			os.Exit(1)
		}

		// close the files
		remoteFile.Close()
		localFile.Close()

		// set the modification time of the local file to match the value saved from the remote file (from before the download)
		err = os.Chtimes(localEntryFullPath, time.Now(), modTime)
	}

	if filesTransfered {
		fmt.Println() // add a gap between download and upload messages
	}

	// iterate over the upload list
	filesTransfered = false
	for _, entryName := range uploadList {
		filesTransfered = true // set a flag to indicate that files have been uploaded (used to determine whether to print a gap between upload and sync complete messages)

		fmt.Println("Uploading " + ansiUpload + entryName + backend.AnsiReset)

		// store path to local entry
		localEntryFullPath := backend.TargetLocationFormat(entryName)

		// save modification time of local file
		var fileInfo os.FileInfo
		fileInfo, err = os.Stat(localEntryFullPath)
		if err != nil {
			fmt.Println(backend.AnsiError+"Sync failed - Unable to get local file info (modtime):", err.Error()+backend.AnsiReset)
			os.Exit(1)
		}
		modTime := fileInfo.ModTime()

		// open local file
		var localFile *os.File
		localFile, err = os.Open(localEntryFullPath)
		if err != nil {
			fmt.Println(backend.AnsiError+"Sync failed - Unable to open local file:", err.Error()+backend.AnsiReset)
			os.Exit(1)
		}

		// store path to remote entry
		remoteEntryFullPath := targetLocationFormatSFTP(entryName, sshEntryRoot, sshIsWindows)

		// create remote file
		var remoteFile *sftp.File
		remoteFile, err = sftpClient.Create(remoteEntryFullPath)
		if err != nil {
			fmt.Println(backend.AnsiError+"Sync failed - Unable to create remote file ("+remoteEntryFullPath+"):", err.Error()+backend.AnsiReset)
			os.Exit(1)
		}

		// upload the file
		_, err = localFile.WriteTo(remoteFile)
		if err != nil {
			fmt.Println(backend.AnsiError+"Sync failed - Unable to upload local file:", err.Error()+backend.AnsiReset)
			os.Exit(1)
		}

		// close the files
		localFile.Close()
		remoteFile.Close()

		// set the modification time of the remote file to match the value saved from the local file (from before the upload)
		err = sftpClient.Chtimes(remoteEntryFullPath, time.Now(), modTime)
	}

	if filesTransfered {
		fmt.Println() // add a gap between upload and sync complete messages
	}
}

// syncLists determines which entries need to be downloaded and uploaded for synchronizations and calls sftpSync with this information
// using maps means that syncing will be done in an arbitrary order, but it is a worthy tradeoff for speed and simplicity
func syncLists(localEntryModMap, remoteEntryModMap map[string]int64, manualSync bool) {
	// initialize slices to store entries that need to be downloaded or uploaded
	var downloadList, uploadList []string

	// iterate over client entries
	for entry, localModTime := range localEntryModMap {
		// check if the entry is present in the server map
		if remoteModTime, present := remoteEntryModMap[entry]; present {
			// entry exists on both client and server, compare mod times
			if remoteModTime > localModTime {
				fmt.Println(ansiDownload+entry+backend.AnsiReset, "is newer on server, adding to download list")
				downloadList = append(downloadList, entry)
			} else if remoteModTime < localModTime {
				fmt.Println(ansiUpload+entry+backend.AnsiReset, "is newer on client, adding to upload list")
				uploadList = append(uploadList, entry)
			}
			// remove entry from remoteEntryModMap (process of elimination)
			delete(remoteEntryModMap, entry)
		} else {
			fmt.Println(ansiUpload+entry+backend.AnsiReset, "does not exist on server, adding to upload list")
			uploadList = append(uploadList, entry)
		}
	}

	// iterate over remaining entries in remoteEntryModMap
	for entry := range remoteEntryModMap {
		fmt.Println(ansiDownload+entry+backend.AnsiReset, "does not exist on client, adding to download list")
		downloadList = append(downloadList, entry)
	}

	// call sftpSync with the download and upload lists
	if max(len(downloadList), len(uploadList)) > 0 { // only call sftpSync if there are entries to download or upload
		fmt.Println() // add a gap between list-add messages and the actual sync messages from sftpSync
		sftpSync(downloadList, uploadList, manualSync)
	}

	fmt.Println("Client is synchronized with server")
}

// ShearRemoteFromClient removes the target file or directory from the local system and calls the server to remove it remotely and add it to the deletions list
// can safely be called in offline mode, as well, so this is the intended interface for shearing (ShearLocal should only be used directly in the server binary)
func ShearRemoteFromClient(targetLocationIncomplete string) {
	deviceID := ShearLocal(targetLocationIncomplete, "") // remove the target from the local system and get the device ID of the client

	// call the server to remotely shear the target and add it to the deletions list
	GetSSHOutput("libmuttonserver shear", deviceID+"\n"+strings.ReplaceAll(targetLocationIncomplete, backend.PathSeparator, "\x1d"), false)

	os.Exit(0) // sync is not required after shearing since the target has already been removed from the local system
}

// deletionSync removes entries from the client that have been deleted on the server (multi-client deletion)
func deletionSync(deletions []string) {
	var filesDeleted bool
	for _, deletion := range deletions {
		filesDeleted = true // set a flag to indicate that files have been deleted (used to determine whether to print a gap between deletion and other messages)
		fmt.Println(ansiDelete+deletion+backend.AnsiReset, "has been sheared, removing locally (if it exists)")
		os.RemoveAll(backend.TargetLocationFormat(deletion))
	}

	if filesDeleted {
		fmt.Println() // add a gap between deletion and other messages
	}
}

// AddFolderRemoteFromClient creates a new entry-containing directory on the local system and calls the server to create the folder remotely
func AddFolderRemoteFromClient(targetLocationIncomplete string) {
	AddFolderLocal(targetLocationIncomplete)                                                                                      // add the folder on the local system
	GetSSHOutput("libmuttonserver addfolder", strings.ReplaceAll(targetLocationIncomplete, backend.PathSeparator, "\x1d"), false) // call the server to create the folder remotely

	os.Exit(0)
}

// folderSync creates folders on the client (from the given list of folder names)
func folderSync(folders []string) {
	for _, folder := range folders {
		// store the full local path of the folder
		folderFullPath := backend.TargetLocationFormat(folder)

		// check if folder already exists
		isFile, isAccessible := backend.TargetIsFile(folderFullPath, false, 1)

		if !isFile && !isAccessible {
			os.MkdirAll(folderFullPath, 0700)
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
