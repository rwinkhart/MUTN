package offline

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func TempInit(configFileMap map[string]string) {
	// create EntryRoot and ConfigDir
	dirInit()

	// remove existing config file
	removeFile(ConfigPath)

	// ensure textEditor is set
	if configFileMap["textEditor"] == "" {
		textEditor := os.Getenv("EDITOR")
		if textEditor == "" {
			textEditor = fallbackEditor
		}
		configFileMap["textEditor"] = textEditor
	}

	// create and write config file
	configFile, _ := os.OpenFile(ConfigPath, os.O_CREATE|os.O_WRONLY, 0600)
	defer configFile.Close()
	configFile.WriteString("[LIBMUTTON]\n")
	for key, value := range configFileMap {
		configFile.WriteString(key + " = " + value + "\n")
	}

	os.Exit(0)
}

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

// GpgKeyGen generates a new GPG key and returns the key ID as a string
func GpgKeyGen() string {
	createFile(ConfigDir+"/gpg-gen", []string{"Key-Type: eddsa", "Key-Curve: ed25519", "Key-Usage: sign", "Subkey-Type: ecdh", "Subkey-Curve: cv25519", "Subkey-Usage: encrypt", "Name-Real: libmutton", "Name-Comment: gpg-libmutton", "Name-Email: github.com/rwinkhart/libmutton", "Expire-Date: 0"})
	cmd := exec.Command("gpg", "-q", "--batch", "--generate-key", ConfigDir+"/gpg-gen")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()

	cmd = exec.Command("gpg", "-k", "--with-colons")
	gpgOutputBytes, _ := cmd.Output()
	gpgOutputLines := strings.Split(string(gpgOutputBytes), "\n")
	uid := strings.Split(gpgOutputLines[len(gpgOutputLines)-4], ":")[9]
	fmt.Println(uid)
	return uid
}
