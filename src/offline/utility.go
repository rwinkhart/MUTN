package offline

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// TargetIsFile TargetStatusCheck checks if the targetLocation is a file, directory, or is inaccessible
// returns: isFile, isAccessible
func TargetIsFile(targetLocation string, errorOnFail bool) (bool, bool) {
	targetInfo, err := os.Stat(targetLocation)
	if err != nil {
		if errorOnFail {
			fmt.Println(AnsiError + "Failed to access \"" + targetLocation + "\" - ensure it exists and has the correct permissions" + AnsiReset)
			os.Exit(1)
		} else {
			return false, false
		}
	}
	if targetInfo.IsDir() {
		return false, true
	} else {
		return true, true
	}
}

// writeToStdin writes a string to a command's stdin
func writeToStdin(cmd *exec.Cmd, input string) {
	stdin, _ := cmd.StdinPipe()
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, input)
	}()
}
