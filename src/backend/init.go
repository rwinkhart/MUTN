package backend

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// GpgUIDListGen generates a list of all GPG key IDs on the system and returns them as a slice of strings
func GpgUIDListGen() []string {
	cmd := exec.Command("gpg", "-k", "--with-colons")
	gpgOutputBytes, _ := cmd.Output()
	gpgOutputLines := strings.Split(string(gpgOutputBytes), "\n")
	var uidSlice []string
	for _, line := range gpgOutputLines {
		if strings.HasPrefix(line, "uid") {
			uid := strings.Split(line, ":")[9]
			uidSlice = append(uidSlice, uid)
		}
	}
	return uidSlice
}

// GpgKeyGen generates a new GPG key and returns the key ID
func GpgKeyGen() string {
	gpgGenTempFile := CreateTempFile()
	defer os.Remove(gpgGenTempFile.Name())

	// create and write gpg-gen file
	unixTime := strconv.FormatInt(time.Now().Unix(), 10)
	gpgGenTempFile.WriteString(strings.Join([]string{"Key-Type: eddsa", "Key-Curve: ed25519", "Key-Usage: sign", "Subkey-Type: ecdh", "Subkey-Curve: cv25519", "Subkey-Usage: encrypt", "Name-Real: libmutton-" + unixTime, "Name-Comment: gpg-libmutton", "Name-Email: github.com/rwinkhart/libmutton", "Expire-Date: 0"}, "\n"))

	// close gpg-gen file
	gpgGenTempFile.Close()

	// generate GPG key based on gpg-gen file
	cmd := exec.Command("gpg", "-q", "--batch", "--generate-key", gpgGenTempFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()

	return "libmutton-" + unixTime + " (gpg-libmutton) <github.com/rwinkhart/libmutton>"
}

// DirInit creates the libmutton directories
func DirInit(preserveOldConfigDir bool) {
	// create EntryRoot
	err := os.MkdirAll(EntryRoot, 0700)
	if err != nil {
		fmt.Println(AnsiError + "Failed to create \"" + EntryRoot + "\":" + err.Error() + AnsiReset)
		os.Exit(1)
	}

	// remove existing config directory (if it exists and not in append mode)
	if !preserveOldConfigDir {
		_, isAccessible := TargetIsFile(ConfigDir, false, 1)
		if isAccessible {
			err = os.RemoveAll(ConfigDir)
			if err != nil {
				fmt.Println(AnsiError + "Failed to remove existing config directory: " + err.Error() + AnsiReset)
				os.Exit(1)

			}
		}
	}

	// create config directory w/devices subdirectory
	err = os.MkdirAll(ConfigDir+PathSeparator+"devices", 0700)
	if err != nil {
		fmt.Println(AnsiError + "Failed to create \"" + ConfigDir + "\":" + err.Error() + AnsiReset)
		os.Exit(1)
	}
}
