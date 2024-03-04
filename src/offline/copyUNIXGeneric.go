//go:build !windows && !darwin

package offline

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// TODO Implement support for MacOS via pbcopy, Termux via termux-clipboard-set (in separate files)

// CopyField copies a field from an entry to the clipboard
func CopyField(targetLocation string, field int, executableName string) {
	if isFile, _ := TargetIsFile(targetLocation, true); isFile {
		decryptedEntry := DecryptGPG(targetLocation)
		var copySubject string // will store data to be copied

		// ensure field exists in entry
		if len(decryptedEntry) > field {
			copySubject = decryptedEntry[field]
		} else {
			fmt.Println(AnsiError + "Field does not exist in entry" + AnsiReset)
			os.Exit(1)
		}

		// ensure field is not blank
		if copySubject == "" {
			fmt.Println(AnsiError + "Field is empty" + AnsiReset)
			os.Exit(1)
		}

		var envSet bool // track whether environment variables are set
		var cmd *exec.Cmd
		// determine whether to use wl-copy (Wayland) or xclip (X11)
		if _, envSet = os.LookupEnv("WAYLAND_DISPLAY"); envSet {
			cmd = exec.Command("wl-copy")
		} else if _, envSet = os.LookupEnv("DISPLAY"); envSet {
			cmd = exec.Command("xclip", "-sel", "c")
		} else {
			fmt.Println(AnsiError + "Clipboard platform could not be determined - note that the clipboard does not function in a raw TTY" + AnsiReset)
			os.Exit(1)
		}

		writeToStdin(cmd, copySubject)
		err := cmd.Run()
		if err != nil {
			fmt.Println(AnsiError + "Failed to copy to clipboard: " + err.Error() + AnsiReset)
			os.Exit(1)
		}

		cmd = exec.Command(executableName, "clipclear")
		writeToStdin(cmd, copySubject)
		err = cmd.Start()
		if err != nil {
			fmt.Println(AnsiError + "Failed to launch automated clipboard clearing process - does this libmutton implementation support the \"clipclear\" argument?" + AnsiReset)
			os.Exit(1)
		}

	} else {
		fmt.Println(AnsiError + "Failed to read \"" + targetLocation + "\" - it is a directory" + AnsiReset)
	}
	os.Exit(0)
}

// ClipClear is called in a separate process to clear the clipboard after 30 seconds
// TODO Make clear time period configurable
func ClipClear(oldContents string) {
	time.Sleep(30 * time.Second)

	var envSet bool   // track whether environment variables are set
	var platform bool // track clipboard platform, false for Wayland, true for X11
	var cmd *exec.Cmd
	// determine whether to use wl-copy (Wayland) or xclip (X11)
	if _, envSet = os.LookupEnv("WAYLAND_DISPLAY"); envSet {
		cmd = exec.Command("wl-paste")
	} else if _, envSet = os.LookupEnv("DISPLAY"); envSet {
		cmd = exec.Command("xclip", "-o", "-sel", "c")
		platform = true
	}
	newContents, _ := cmd.Output()

	if oldContents == strings.TrimRight(string(newContents), "\r\n") {
		switch platform {
		case false:
			cmd = exec.Command("wl-copy", "-c")
		case true:
			cmd = exec.Command("xclip", "-i", "/dev/null", "-sel", "c")
		}
		err := cmd.Run()
		if err != nil {
			fmt.Println(AnsiError + "Failed to clear clipboard: " + err.Error() + AnsiReset)
			os.Exit(1)
		}
	}
	os.Exit(0)
}
