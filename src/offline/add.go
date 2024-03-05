package offline

import (
	"fmt"
	"os"
)

// AddFolder creates a new directory at targetLocation
func AddFolder(targetLocation string) {
	// create the directory specified by targetLocation
	err := os.Mkdir(targetLocation, 0700)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println(AnsiError + "Directory already exists" + AnsiReset)
			os.Exit(1)
		} else {
			fmt.Println(AnsiError + "Failed to create directory: " + err.Error() + AnsiReset)
			os.Exit(1)
		}
	}
	// TODO If in online mode, create the directory on the server
	os.Exit(0)
}

// WriteEntry writes entryData to an encrypted file at targetLocation
func WriteEntry(targetLocation string, entryData []string) {
	encryptedBytes := EncryptGPG(entryData)
	err := os.WriteFile(targetLocation, encryptedBytes, 0600)
	if err != nil {
		fmt.Println(AnsiError + "Failed to write to file: " + err.Error() + AnsiReset)
		os.Exit(1)
	}
}
