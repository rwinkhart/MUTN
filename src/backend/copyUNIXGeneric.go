//go:build !windows && !darwin && !termux

package backend

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// copyField copies a field from an entry to the clipboard
func copyField(copySubject string, executableName string) {
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

	// launch clipboard clearing process if executableName is provided
	if executableName != "" {
		cmd = exec.Command(executableName, "clipclear")
		writeToStdin(cmd, copySubject)
		err = cmd.Start()
		if err != nil {
			fmt.Println(AnsiError + "Failed to launch automated clipboard clearing process - does this libmutton implementation support the \"clipclear\" argument?" + AnsiReset)
			os.Exit(1)
		}
		os.Exit(0) // only exit if clipboard clearing process is launched, otherwise assume continuous clipboard refresh
	}
}

// clipClear is called in a separate process to clear the clipboard after 30 seconds
func clipClear(oldContents string) {
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
