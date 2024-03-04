//go:build !windows && !darwin

package offline

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// TODO Handle cmd.Run errors, especially those due to missing clipboard utilities
// TODO Handle index out of range errors when copying fields that do not exist
// TODO Implement support for MacOS via pbcopy, Termux via termux-clipboard-set (in separate files)

// CopyField copies a field from an entry to the clipboard
func CopyField(targetLocation string, field uint8, executableName string) {
	if isFile, _ := TargetIsFile(targetLocation, true); isFile {
		copySubject := DecryptGPG(targetLocation)[field]

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
		cmd.Run()

		// TODO Strongly encourage other implementations to support a clipclear argument
		// TODO If an implementation opts out of doing this, this may error out
		cmd = exec.Command(executableName, "clipclear")
		writeToStdin(cmd, copySubject)
		cmd.Start()

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
		cmd.Run()
	}
	os.Exit(0)
}
