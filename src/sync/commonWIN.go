//go:build windows

package sync

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/backend"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// WalkEntryDir walks the entry directory and returns lists of all files and directories found (two separate lists)
// regardless of platform, all paths are stored with forward slashes (UNIX-style)
func WalkEntryDir() ([]string, []string) {
	// define file/directory containing slices so that they may be accessed by the anonymous WalkDir function
	var fileList []string
	var dirList []string

	// walk entry directory
	_ = filepath.WalkDir(backend.EntryRoot,
		func(fullPath string, entry fs.DirEntry, err error) error {

			// check for errors encountered while walking directory
			if err != nil {
				if os.IsNotExist(err) {
					fmt.Println(backend.AnsiError+"The entry directory does not exist - run \""+os.Args[0], "init"+"\" to create it"+backend.AnsiReset)
				} else {
					// otherwise, print the source of the error
					fmt.Println(backend.AnsiError + "An unexpected error occurred while generating the entry list: " + err.Error() + backend.AnsiReset)
				}
				os.Exit(1)
			}

			// trim root path from each path before storing and replace backslashes with forward slashes
			trimmedPath := strings.ReplaceAll(fullPath[rootLength:], "\\", "/")

			// append the path to the appropriate slice
			if !entry.IsDir() {
				fileList = append(fileList, trimmedPath)
			} else {
				dirList = append(dirList, trimmedPath)
			}

			return nil
		})

	return fileList, dirList
}
