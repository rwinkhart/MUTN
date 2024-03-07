//go:build windows

package cli

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/offline"
	"strings"
)

func printDirectoryHeader(vanityDirectory string, indent int) {
	fmt.Printf("\n\n"+strings.Repeat(" ", indent*2)+ansiDirectoryHeader+"%s/"+offline.AnsiReset+"\n", strings.ReplaceAll(vanityDirectory, offline.PathSeparator, "/"))
}
