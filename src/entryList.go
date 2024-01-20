package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// global variables used only in this file
var (
	fileList []string
	dirList  []string
)

// global constants used only in this file
const (
	ansiAlternateEntryColor = "\033[38;5;8m"
)

// processing for printing file entries (determines color, line wrapping, and prints)
func printFileEntry(entry string, lastSlash int, charCounter int, colorAlternator int8) (int, int8) {
	// determine color to print fileEntryName (alternate each time function is run)
	var colorCode string
	if colorAlternator > 0 {
		colorCode = ""
	} else {
		colorCode = ansiAlternateEntryColor
	}
	colorAlternator = -colorAlternator

	// trim the containing directory and file extension from the entry to determine fileEntryName
	fileEntryName := entry[lastSlash:]
	fileEntryName = fileEntryName[:len(fileEntryName)-4]

	// determine whether to wrap to a new line (+1 is to account for trailing spaces)
	charCounter += len(fileEntryName) + 1
	if charCounter >= width {
		charCounter = len(fileEntryName) + 1
		fmt.Println()
	}

	// print fileEntryName to screen
	fmt.Printf("%s%s%s ", colorCode, fileEntryName, ansiReset)

	return charCounter, colorAlternator
}

// EntryListGen generates and displays full libmutton entry list
func EntryListGen() {
	fmt.Print("\n" + ansiBlackOnWhite + "libmutton entries:" + ansiReset)

	// walk entry directory
	_ = filepath.WalkDir(EntryRoot,
		func(fullPath string, entry fs.DirEntry, err error) error {

			// check for errors encountered while walking directory
			if err != nil {
				// create EntryRoot if the error is the result of it not existing on the system
				if os.IsNotExist(err) {
					_ = os.Mkdir(EntryRoot, 0700)
					dirList = append(dirList, "")
				} else {
					// otherwise, print the source of the error
					fmt.Print("\n\n\033[38;5;9mAn unexpected error occurred while generating the entry list: " + err.Error() + ansiReset)
				}
				// quit walking EntryRoot and return nil to allow the program to continue
				return nil
			}

			// trim root path from each path before storing
			trimmedPath := fullPath[rootLength:]

			// create three separate slices for root-level entries, all other entries, and all subdirectories
			// root-level entries get their own slice so that they can be alphabetically sorted without the chance of directories being placed in from of them
			if !entry.IsDir() {
				fileList = append(fileList, trimmedPath)
			} else {
				dirList = append(dirList, trimmedPath)
			}

			return nil
		})

	// dirList iteration
	dirListLength := len(dirList) // save length for multiple references below
	charCounter := 0              // track whether to line-wrap based on character count in line
	var colorAlternator int8 = 1  // track alternating colors for each printed entry name
	var containsSubdirectory bool // indicates whether the current directory contains a subdirectory
	var indent int                // visual indentation multiplier
	for i, directory := range dirList {

		// reset formatting variables for new directory
		charCounter = 0
		colorAlternator = 1

		// determine directory's indentation multiplier based on PathSeparator occurrences - only run if last directory contained a subdirectory (indicating that the current directory is a subdirectory)
		indent = strings.Count(directory, PathSeparator) - 1 // subtract 1 to account for trailing PathSeparator

		// check if next directory is within the current one
		if dirListLength > i+1 {
			nextDir := dirList[i+1]
			if directory == nextDir[:strings.LastIndex(nextDir, PathSeparator)] {
				containsSubdirectory = true
			} else {
				containsSubdirectory = false
			}
		} else {
			containsSubdirectory = false
		}

		// fileList iteration
		containsFiles := false // indicates whether the current directory contains files (entries)
		for _, file := range fileList {

			// get index of last occurrence of pathSeparator in trimmed entry path (used to split entry's containing directory and entry's name)
			lastSlash := strings.LastIndex(file, PathSeparator) + 1

			// print the current file if it belongs in the current directory - otherwise, break the loop and move on to the next directory
			if file[:lastSlash-1] == directory {

				// print directory header if this is the first run of the loop
				if !containsFiles {
					containsFiles = true
					if !Windows { // for consistency, format directories with UNIX-style path separators on all platforms
						fmt.Printf("\n\n"+strings.Repeat(" ", indent*2)+"\033[38;5;7;48;5;8m%s/"+ansiReset+"\n", directory)
					} else {
						fmt.Printf("\n\n"+strings.Repeat(" ", indent*2)+"\033[38;5;7;48;5;8m%s/"+ansiReset+"\n", strings.ReplaceAll(directory, PathSeparator, "/"))
					}
				}

				if indent > 0 { // indent entry list to the same degree as the directory header
					fmt.Print(strings.Repeat(" ", indent*2))
				}
				charCounter, colorAlternator = printFileEntry(file, lastSlash, charCounter, colorAlternator)
			}
		}

		if !containsFiles { // if the current directory contains no files...
			if !containsSubdirectory { // nor does it contain any subdirectories...
				if dirListLength > 1 { // and directories besides the root-level exist... display directory header and empty directory warning
					if !Windows { // for consistency, format directories with UNIX-style path separators on all platforms
						fmt.Printf("\n\n"+strings.Repeat(" ", indent*2)+"\033[38;5;7;48;5;8m%s/"+ansiReset+"\n", directory)
					} else {
						fmt.Printf("\n\n"+strings.Repeat(" ", indent*2)+"\033[38;5;7;48;5;8m%s/"+ansiReset+"\n", strings.ReplaceAll(directory, PathSeparator, "/"))
					}
					fmt.Print(strings.Repeat(" ", indent*2) + "\033[38;5;11m-empty directory-" + ansiReset)
				} else { // warn if the only thing that exists is the root-level directory
					fmt.Print("\n\nNothing's here! For help creating your first entry, run \"mutn help\".")
				}
			}
		}
	}

	// print trailing new lines for proper spacing after entry list is complete
	fmt.Print("\n\n")
}
