package main

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"github.com/rwinkhart/MUTN/src/sync"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(backend.AnsiError + "No arguments supplied" + backend.AnsiReset) // TODO add server help menu
		os.Exit(0)
	}

	switch args[1] {
	case "fetch":
		sync.GetRemoteDataFromServer()
	case "shear":
		deviceIDTargetLocation := strings.Split(args[2], "\x1d")
		targetLocationIncomplete := strings.ReplaceAll(strings.ReplaceAll(deviceIDTargetLocation[1], "\x1f", " "), "\x1e", backend.PathSeparator)
		sync.Shear(targetLocationIncomplete, deviceIDTargetLocation[0])
	default:
		fmt.Println(backend.AnsiError + "Invalid argument" + backend.AnsiReset) // TODO add server help menu
	}
}
