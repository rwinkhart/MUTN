package backend

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

// loadConfig loads the libmutton.ini file and returns the configuration
// utility function for ParseConfig and WriteConfig, do not call directly
func loadConfig() *ini.File {
	cfg, err := ini.Load(ConfigPath)
	if err != nil {
		fmt.Println(AnsiError + "Failed to load libmutton.ini: " + err.Error() + AnsiReset)
		os.Exit(1)
	}
	return cfg
}

// ParseConfig reads the libmutton.ini file and returns a slice of values for the specified keys
// requires readKeys: a slice of key names (indicates requested values)
// requires missingValueError: an error message to display if a key is missing a value, set to "" for auto-generated or "0" to exit silently with status 0
// returns config: a slice of values for the specified keys
func ParseConfig(readKeys []string, missingValueError string) []string {
	cfg := loadConfig()

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

// WriteConfig writes the provided key-value pairs to the libmutton.ini file
func WriteConfig(configFileMap map[string]string, append bool) {
	var cfg *ini.File
	var libmuttonSection *ini.Section

	if append {
		// load existing ini file
		cfg = loadConfig()

		// acquire LIBMUTTON section
		libmuttonSection, _ = cfg.GetSection("LIBMUTTON")
	} else {
		// create empty ini file
		cfg = ini.Empty()

		// create LIBMUTTON section
		libmuttonSection, _ = cfg.NewSection("LIBMUTTON")

		// set default textEditor value
		if configFileMap["textEditor"] == "" {
			configFileMap["textEditor"] = textEditorFallback()
		}
	}

	// write provided configFileMap key-value pairs to the LIBMUTTON section
	for key, value := range configFileMap {
		libmuttonSection.Key(key).SetValue(value)
	}

	// save the new config file
	err := cfg.SaveTo(ConfigPath)
	if err != nil {
		fmt.Println(AnsiError + "Failed to save libmutton.ini: " + err.Error() + AnsiReset)
		os.Exit(1)
	}
}

// libmutton.ini layout
// [LIBMUTTON]
// gpgID = <gpg key id>
// textEditor = <editor command> TODO move to MUTN section heading, as it only applies to the CLI implementation
// sshUser = <remote user>
// sshIP = <remote ip>
// sshPort = <remote ssh port>
// sshKey = <ssh private key identity file path>
// sshKeyProtected = <true/false>
// netPinEnabled = <true/false> TODO netPin functionality not yet implemented
// sshEntryRoot = <remote entry root>
// sshIsWindows = <true/false>

// Developers of alternative clients:
// If you are adding additional settings to the config file,
// please create a new section heading for your app-specific settings.
