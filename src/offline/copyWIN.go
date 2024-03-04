//go:build windows

package offline

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// TODO Avoid index out of range errors when copying fields that do not exist

// CopyField copies a field from an entry to the clipboard
func CopyField(targetLocation string, field uint8, executableName string) {
	if isFile, _ := TargetIsFile(targetLocation, true); isFile {
		copySubject := DecryptGPG(targetLocation)[field]

		cmd := exec.Command("powershell.exe", "-c", fmt.Sprintf("echo '%s' | Set-Clipboard", copySubject))
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

	cmd := exec.Command("powershell.exe", "-c", "Get-Clipboard")
	newContents, _ := cmd.Output()

	if oldContents == strings.TrimRight(string(newContents), "\r\n") {
		cmd = exec.Command("powershell.exe", "-c", "Set-Clipboard")
		err := cmd.Run()
		if err != nil {
			fmt.Println(AnsiError + "Failed to clear clipboard: " + err.Error() + AnsiReset)
			os.Exit(1)
		}
	}
	os.Exit(0)
}
