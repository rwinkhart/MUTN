package offline

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func CopyField(targetLocation string, field uint8) {
	if isFile, _ := TargetIsFile(targetLocation, true); isFile {
		copySubject := DecryptGPG(targetLocation)[field]
		cmd := exec.Command("wl-copy")
		stdin, _ := cmd.StdinPipe()
		go func() {
			defer stdin.Close()
			io.WriteString(stdin, copySubject)
		}()
		cmd.Run()
	} else {
		fmt.Println(AnsiError + "Failed to read \"" + targetLocation + "\" - it is a directory" + AnsiReset)
	}
	os.Exit(0)
}

func clipClear() {

}
