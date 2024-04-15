//go:build !windows

package cli

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"strings"
)

func printDirectoryHeader(vanityDirectory string, indent int) {
	fmt.Printf("\n\n"+strings.Repeat(" ", indent*2)+ansiDirectoryHeader+"%s/"+backend.AnsiReset+"\n", vanityDirectory)
}
