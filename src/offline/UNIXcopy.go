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

// CopyField copies a field from an entry to the clipboard
// TODO Implement support for MacOS via pbcopy, Termux via termux-clipboard-set, X11 via xclip
func CopyField(targetLocation string, field uint8, executableName string) {
	if isFile, _ := TargetIsFile(targetLocation, true); isFile {
		copySubject := DecryptGPG(targetLocation)[field]

		cmd := exec.Command("wl-copy")
		writeToStdin(cmd, copySubject)
		cmd.Run()

		// TODO Unpair from mutn, as this breaks library functionality
		// TODO Maybe use this method for mutn and strongly encourage other implementations to support a clipclear argument
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
// TODO Implement support for MacOS via pbcopy, Termux via termux-clipboard-set, X11 via xclip
// TODO Make clear time period configurable
func ClipClear(oldContents string) {
	time.Sleep(30 * time.Second)

	cmd := exec.Command("wl-paste")
	newContents, _ := cmd.Output()

	if oldContents == strings.TrimRight(string(newContents), "\r\n") {
		cmd = exec.Command("wl-copy", "-c")
		cmd.Run()
	}
	os.Exit(0)
}
