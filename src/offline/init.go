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
	err := os.Remove(ConfigPath)
	if err != nil {
		// ignore error if file does not exist
		if !os.IsNotExist(err) {
			fmt.Println(AnsiError + "Failed to remove existing libmutton.ini:" + err.Error() + AnsiReset)
		}
	}

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
