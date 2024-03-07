package offline

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

// ReadConfig reads the libmutton.ini file and returns a map of the requested values
// requires readMap: a map of section names to key names (indicates requested values)
// returns configMap: a map of key names to values (sections are irrelevant)
func ReadConfig(readKeys []string) []string {
	cfg, err := ini.Load(ConfigPath)
	if err != nil {
		fmt.Println(AnsiError + "Failed to load libmutton.ini" + AnsiReset)
		os.Exit(1)
	}

	var config []string

	for _, key := range readKeys {
		config = append(config, cfg.Section("LIBMUTTON").Key(key).String())
	}

	return config
}

// libmuttn.ini layout
// [LIBMUTTON]
// gpgID = <gpg key id>
// textEditor = <editor command>
// onlineMode = <true/false>
// sshError = <true/false>
// netPinEnabled = <true/false>
// remoteUser = <ssh user>
// remoteIP = <ssh ip>
// remotePort = <ssh port>
// identityFile = <path to private key>
