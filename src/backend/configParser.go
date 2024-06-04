package backend

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

// ReadConfig reads the libmutton.ini file and returns a slice of values for the specified keys
// requires readKeys: a slice of key names (indicates requested values)
// requires missingValueError: an error message to display if a key is missing a value, set to "" for auto-generated or "0" to exit silently with status 0
// returns config: a slice of values for the specified keys
func ReadConfig(readKeys []string, missingValueError string) []string {
	cfg, err := ini.Load(ConfigPath)
	if err != nil {
		fmt.Println(AnsiError + "Failed to load libmutton.ini: " + err.Error() + AnsiReset)
		os.Exit(1)
	}

	var config []string

	for _, key := range readKeys {
		keyConfig := cfg.Section("LIBMUTTON").Key(key).String()

		// ensure specified key has a value
		if keyConfig == "" {
			switch missingValueError {
			case "":
				fmt.Println(AnsiError + "Failed to find value for key \"" + key + "\" in section \"[LIBMUTTON]\" in libmutton.ini" + AnsiReset)
			case "0":
				os.Exit(0)
			default:
				fmt.Println(AnsiError + missingValueError + AnsiReset)
			}
			os.Exit(1)
		}

		config = append(config, keyConfig)
	}

	return config
}

// libmuttn.ini layout
// [LIBMUTTON]
// gpgID = <gpg key id>
// textEditor = <editor command>
// sshUser = <remote user>
// sshIP = <remote ip>
// sshPort = <remote ssh port>
// sshKey = <ssh private key identity file path>
// sshKeyProtected = <true/false>
// netPinEnabled = <true/false>
