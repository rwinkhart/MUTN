package backend

import (
	"fmt"
	"os"
)

// AddFolder creates a new directory at targetLocation
func AddFolder(targetLocation string) {
	// create the directory specified by targetLocation
	err := os.Mkdir(targetLocation, 0700)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println(AnsiError + "Directory already exists" + AnsiReset)
			os.Exit(1)
		} else {
			fmt.Println(AnsiError + "Failed to create directory: " + err.Error() + AnsiReset)
			os.Exit(1)
		}
	}
	// TODO If in online mode, create the directory on the server
	os.Exit(0)
}
