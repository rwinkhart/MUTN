package offline

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

// TODO Handle cmd.Run errors, especially those due to missing clipboard utilities

// CopyField copies a field from an entry to the clipboard
// TODO Implement support for Windows via powershell, MacOS via pbcopy, Termux via termux-clipboard-set, X11 via xclip
// TODO Monitor Go support for Haiku, if possible implement support via clipboard command
func CopyField(targetLocation string, field uint8, executableName string) {
	if isFile, _ := TargetIsFile(targetLocation, true); isFile {
		copySubject := DecryptGPG(targetLocation)[field]

		//hash := getSha512Sum(copySubject)

		cmd := exec.Command("wl-copy")
		writeToStdin(cmd, copySubject)
		cmd.Run()

		// TODO Unpair from mutn, as this breaks library functionality
		// TODO Maybe use this method for mutn and strongly encourage other implementations to support a clipclear argument
		// TODO If an implementation opts out of doing this, this may error out
		//cmd = exec.Command(executableName, "clipclear", hash)
		cmd = exec.Command(executableName, "clipclear")
		writeToStdin(cmd, copySubject)
		cmd.Start()

	} else {
		fmt.Println(AnsiError + "Failed to read \"" + targetLocation + "\" - it is a directory" + AnsiReset)
	}
	os.Exit(0)
}

// ClipClear is called in a separate process to clear the clipboard after 30 seconds
// TODO Implement support for Windows via powershell, MacOS via pbcopy, Termux via termux-clipboard-set, X11 via xclip
// TODO Make clear time period configurable
func ClipClear(oldContents string) {
	time.Sleep(30 * time.Second)

	cmd := exec.Command("wl-paste")
	newContents, _ := cmd.Output()

	if oldContents == strings.TrimRight(string(newContents), "\n") {
		cmd := exec.Command("wl-copy", "-c")
		cmd.Run()
	}
	os.Exit(0)
}

// writeToStdin writes a string to a command's stdin
func writeToStdin(cmd *exec.Cmd, input string) {
	stdin, _ := cmd.StdinPipe()
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, input)
	}()
}
