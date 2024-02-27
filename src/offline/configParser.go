package offline

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var Config = make(map[string]interface{})

func ReadConfig() {
	cfg, err := ini.Load(home + "/.config/libmutton/libmutton.ini")
	if err != nil {
		fmt.Println(AnsiError + "\033[38;5;9mFailed to load libmutton.ini" + AnsiReset)
		os.Exit(1)
	}

	// OFFLINE
	Config["gpgID"] = cfg.Section("OFFLINE").Key("gpgID").String()
	Config["textEditor"] = cfg.Section("OFFLINE").Key("textEditor").String()

	// ONLINE (conditional)
	onlineMode, err := cfg.Section("OFFLINE").Key("onlineMode").Bool()
	if err != nil {
		fmt.Println(AnsiError + "\033[38;5;9mFailed to enable online (synced) mode - [OFFLINE]/onlineMode in libmutton.ini must be a boolean value - continuing in offline mode" + AnsiReset)
		onlineMode = false
	}
	if onlineMode {
		Config["sshError"], err = cfg.Section("ONLINE").Key("sshError").Bool()
		if err != nil {
			cfg.Section("ONLINE").Key("sshError").SetValue("true")
		}
		Config["netPinEnabled"], err = cfg.Section("ONLINE").Key("netPinEnabled").Bool()
		if err != nil {
			cfg.Section("ONLINE").Key("netPinEnabled").SetValue("false")
		}
		Config["remoteUser"] = cfg.Section("ONLINE").Key("remoteUser").String()
		Config["remoteIP"] = cfg.Section("ONLINE").Key("remoteIP").String()
		Config["remotePort"], err = cfg.Section("ONLINE").Key("remotePort").Int()
		if err != nil {
			fmt.Println(AnsiError + "\033[38;5;9mFailed to enable online (synced) mode - [ONLINE]/remotePort in libmutton.ini must be an integer value- continuing in offline mode" + AnsiReset)
			// remotePort == 0 represents offline-only mode
			Config["remotePort"] = 0
		}
		Config["identityFile"] = cfg.Section("ONLINE").Key("identityFile").String()
	} else {
		// remotePort == 0 represents offline-only mode
		Config["remotePort"] = 0
	}

	// DEBUG config
	//for key, value := range Config {
	//	fmt.Print(key + ": ")
	//	fmt.Println(value)
	//}

}
