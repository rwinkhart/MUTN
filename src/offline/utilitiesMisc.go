package offline

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// TargetIsFile TargetStatusCheck checks if the targetLocation is a file, directory, or is inaccessible
// failCondition: 0 = fail on inaccessible, 1 = fail on inaccessible/file, 2 = fail on inaccessible/directory
// returns: isFile, isAccessible
func TargetIsFile(targetLocation string, errorOnFail bool, failCondition uint8) (bool, bool) {
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
		if errorOnFail && failCondition == 2 {
			fmt.Println(AnsiError + "\"" + targetLocation + "\" is a directory" + AnsiReset)
			os.Exit(1)
		}
		return false, true
	} else {
		if errorOnFail && failCondition == 1 {
			fmt.Println(AnsiError + "\"" + targetLocation + "\" is a file" + AnsiReset)
			os.Exit(1)
		}
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
