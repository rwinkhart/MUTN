package offline

import (
	"fmt"
	"os"
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
