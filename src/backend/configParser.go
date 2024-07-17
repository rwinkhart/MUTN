package backend

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
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
// requires requestedValues: a slice of arrays (length 2) each containing a section and a key name
// requires missingValueError: an error message to display if a key is missing a value, set to "" for auto-generated or "0" to exit/return silently with code 0
// returns config: a slice of values for the specified keys
func ParseConfig(requestedValues [][2]string, missingValueError string) []string {
	cfg := loadConfig()

	var config []string

	for _, pair := range requestedValues {
		value := cfg.Section(pair[0]).Key(pair[1]).String()

		// ensure specified key has a value
		if value == "" {
			switch missingValueError {
			case "":
				fmt.Println(AnsiError + "Failed to find value for key \"" + pair[1] + "\" in section \"[" + pair[0] + "]\" in libmutton.ini" + AnsiReset)
			case "0":
				Exit(0)
			default:
				fmt.Println(AnsiError + missingValueError + AnsiReset)
			}
			os.Exit(1)
		}

		config = append(config, value)
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
