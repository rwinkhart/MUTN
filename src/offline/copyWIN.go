//go:build windows

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
func CopyField(targetLocation string, field uint8, executableName string) {
	if isFile, _ := TargetIsFile(targetLocation, true); isFile {
		copySubject := DecryptGPG(targetLocation)[field]

		cmd := exec.Command("powershell.exe", "-c", fmt.Sprintf("echo '%s' | Set-Clipboard", copySubject))
		cmd.Run()

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

	cmd := exec.Command("powershell.exe", "-c", "Get-Clipboard")
	newContents, _ := cmd.Output()

	if oldContents == strings.TrimRight(string(newContents), "\r\n") {
		cmd = exec.Command("powershell.exe", "-c", "Set-Clipboard")
		cmd.Run()
	}
	os.Exit(0)
}
