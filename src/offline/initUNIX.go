//go:build !windows

package offline

import (
	"fmt"
	"os"
)

const fallbackEditor = "vi" // vi is pre-installed on most UNIX systems

func dirInit() {
	// create EntryRoot
	err := os.MkdirAll(EntryRoot, 0700)
	if err != nil {
		fmt.Println(AnsiError + "Failed to create \"" + EntryRoot + "\":" + err.Error() + AnsiReset)
		os.Exit(1)
	}

	// create config directory
	err = os.MkdirAll(ConfigDir, 0700)
	if err != nil {
		fmt.Println(AnsiError + "Failed to create \"" + ConfigDir + "\":" + err.Error() + AnsiReset)
		os.Exit(1)
	}
}
