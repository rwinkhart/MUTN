//go:build windows

package offline

import (
	"fmt"
	"os"
)

const fallbackEditor = "nvim" // since there is no pre-installed CLI editor on Windows, default to the most popular one

func dirInit() {
	// create EntryRoot (includes config directory on Windows)
	err := os.MkdirAll(EntryRoot, 0700)
	if err != nil {
		fmt.Println(AnsiError + "Failed to create \"" + EntryRoot + "\":" + err.Error() + AnsiReset)
		os.Exit(1)
	}
}
