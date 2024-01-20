package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var (
	fileListRoot []string
	fileListMain []string
	dirListMain  []string
)

// processing for printing file entries (determines color, line wrapping, and prints)
func printFileEntry(entry string, lastSlash int, charCounter int, colorAlternator int8) (int, int8) {
	// determine color to print fileEntryName (alternate each time function is run)
	var colorCode string
	if colorAlternator > 0 {
		colorCode = ""
	} else {
		colorCode = "\033[38;5;8m"
	}
	colorAlternator = -colorAlternator

	// trim the containing directory and file extension from the entry to determine fileEntryName
	fileEntryName := entry[lastSlash:]
	fileEntryName = fileEntryName[:len(fileEntryName)-4]

	// determine whether to wrap to a new line (+1 is to account for trailing spaces)
	charCounter += len(fileEntryName) + 1
	if charCounter >= Width {
		charCounter = len(fileEntryName) + 1
		fmt.Println()
	}

	// print fileEntryName to screen
	fmt.Printf("%s%s\033[0m ", colorCode, fileEntryName)

	return charCounter, colorAlternator
}

// EntryListGen generates and displays full entry list
func EntryListGen() {
	fmt.Print("\n\033[38;5;0;48;5;15mlibmutton entries:\033[0m")

	// walk entry directory
	_ = filepath.WalkDir(EntryRoot,
		func(fullPath string, entry fs.DirEntry, err error) error {

			// check for errors encountered while walking directory
			if err != nil {
				// create EntryRoot if the error is the result of it not existing on the system
				if err.Error() == "lstat "+EntryRoot+": no such file or directory" {
					_ = os.Mkdir(EntryRoot, 0700)
				} else {
					// otherwise, print the source of the error
					fmt.Print("\n\n\033[38;5;9mAn unexpected error occurred while generating the entry list: " + err.Error() + "\033[0m")
				}
				// quit walking EntryRoot and return nil to allow the program to continue
				return nil
			}

			// trim root path from each path before storing
			trimmedPath := fullPath[RootLength:]

			// create three separate slices for root-level entries, all other entries, and all subdirectories
			// root-level entries get their own slice so that they can be alphabetically sorted without the chance of directories being placed in from of them
			if !entry.IsDir() && strings.Count(trimmedPath, PathSeparator) == 1 {
				fileListRoot = append(fileListRoot, trimmedPath)
			} else if !entry.IsDir() {
				fileListMain = append(fileListMain, trimmedPath)
			} else if trimmedPath != "" {
				dirListMain = append(dirListMain, trimmedPath)
			}

			return nil
		})

	// fileListRoot iteration
	charCounter := 0             // set to track whether to line-wrap based on character count in line
	var colorAlternator int8 = 1 // set to allow alternating colors for each printed entry name
	ran := false                 // set to track whether the loop has been run yet
	for _, file := range fileListRoot {

		// print root ("/") header if this is the first run of the loop
		// must be done within loop in order to allow not printing the header if there are no root-level entries
		if !ran {
			ran = true
			fmt.Print("\n\n\033[38;5;7;48;5;8m/\033[0m\n")
		}

		// get index of last occurrence of pathSeparator in trimmed entry path (used to split entry's containing directory and entry's name) and pass to printFileEntry
		charCounter, colorAlternator = printFileEntry(file, strings.LastIndex(file, PathSeparator)+1, charCounter, colorAlternator)

	}

	dirListMainLength := len(dirListMain) // save length for multiple references below
	if dirListMainLength > 0 {            // DirListMain iteration - only run if non-root-level directories are present TODO is this still necessary?
		var containsSubdirectory bool // track whether the current directory contains a subdirectory
		var indent int                // visual indentation multiplier
		for i, directory := range dirListMain {

			// reset formatting variables for new directory
			charCounter = 0
			colorAlternator = 1

			// determine directory's indentation multiplier based on PathSeparator occurrences - only run if last directory contained a subdirectory (indicating that the current directory is a subdirectory)
			indent = strings.Count(directory, PathSeparator) - 1 // subtract 1 to account for trailing PathSeparator

			// check if next directory is within the current one
			if dirListMainLength > i+1 {
				nextDir := dirListMain[i+1]
				if directory == nextDir[:strings.LastIndex(nextDir, PathSeparator)] {
					containsSubdirectory = true
				} else {
					containsSubdirectory = false
				}
			} else {
				containsSubdirectory = false
			}

			// fileListMain iteration
			ran = false // set to track whether the loop has been run yet
			for _, file := range fileListMain {

				// get index of last occurrence of pathSeparator in trimmed entry path (used to split entry's containing directory and entry's name)
				lastSlash := strings.LastIndex(file, PathSeparator) + 1

				// print the current file if it belongs in the current directory - otherwise, break the loop and move on to the next directory
				if file[:lastSlash-1] == directory {

					// print directory header if this is the first run of the loop
					// must be done within loop in order to allow not printing the header if its directory contains no entries
					if !ran {
						ran = true
						// for consistency, format directories with UNIX-style path separators on all platforms
						if !Windows {
							fmt.Printf("\n\n"+strings.Repeat(" ", indent*2)+"\033[38;5;7;48;5;8m%s/\033[0m\n", directory)
						} else {
							fmt.Printf("\n\n"+strings.Repeat(" ", indent*2)+"\033[38;5;7;48;5;8m%s/\033[0m\n", strings.ReplaceAll(directory, PathSeparator, "/"))
						}
					}

					if indent > 0 { // indent entry list to the same degree as the directory header
						fmt.Print(strings.Repeat(" ", indent*2))
					}
					charCounter, colorAlternator = printFileEntry(file, lastSlash, charCounter, colorAlternator)
				}
			}

			if !ran { // if the current directory contains no files...
				if !containsSubdirectory { // display directory header and empty directory warning if it also contains no subdirectories
					if !Windows { // for consistency, format directories with UNIX-style path separators on all platforms
						fmt.Printf("\n\n"+strings.Repeat(" ", indent*2)+"\033[38;5;7;48;5;8m%s/\033[0m\n", directory)
					} else {
						fmt.Printf("\n\n"+strings.Repeat(" ", indent*2)+"\033[38;5;7;48;5;8m%s/\033[0m\n", strings.ReplaceAll(directory, PathSeparator, "/"))
					}
					fmt.Print(strings.Repeat(" ", indent*2) + "\033[38;5;11m-empty directory-\033[0m")
				}
			}
		}
	} else if !ran {
		fmt.Print("\n\nNothing's here! For help creating your first entry, run \"mutn help\".")
	}

	// print trailing new lines for proper spacing after entry list is complete
	fmt.Print("\n\n")

}
