//go:build windows || (linux && wsl)

package backend

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// copyField copies a field from an entry to the clipboard
func copyField(executableName, copySubject string) {
	cmd := exec.Command("powershell.exe", "-c", fmt.Sprintf("echo '%s' | Set-Clipboard", strings.ReplaceAll(copySubject, "'", "''")))
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

	cmd := exec.Command("powershell.exe", "-c", "Get-Clipboard")
	newContents, err := cmd.Output()
	if err != nil {
		fmt.Println(AnsiError + "Failed to read clipboard contents: " + err.Error() + AnsiReset)
		os.Exit(1)
	}

	if oldContents == strings.TrimRight(string(newContents), "\r\n") {
		cmd = exec.Command("powershell.exe", "-c", "Set-Clipboard")
		err = cmd.Run()
		if err != nil {
			fmt.Println(AnsiError + "Failed to clear clipboard: " + err.Error() + AnsiReset)
			os.Exit(1)
		}
	}
	os.Exit(0)
}
