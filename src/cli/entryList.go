package cli

import (
	"fmt"
	"github.com/rwinkhart/MUTN/src/offline"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// global constants used only in this file
const (
	ansiAlternateEntryColor = "\033[38;5;8m"
)

// calculates and returns the final visual indentation multiplier (needed to adjust indentation for skipped parent directories) - also subtracts "old" text from directory header
func indentSubtractor(skippedDirList []bool, dirList []string, currentDirIndex int, indent int) (int, string) {
	var subtractor int      // tracks how much to subtract from expected indentation multiplier
	var lastPrefixIndex int // tracks the index (in both skippedDirList and dirList) of the last displayed parent directory
	var trimmedDirectory = dirList[currentDirIndex]

	for i, skipped := range skippedDirList[:currentDirIndex] { // checks each skipped directory to determine if it is a parent to the current directory
		if strings.HasPrefix(trimmedDirectory, dirList[i]+offline.PathSeparator) { // if the current directory is the child of this iteration's directory...
			if skipped { // ...and this iteration's directory was skipped...
				subtractor++ // increment the subtractor to indicate that the visual indentation should be reduced
			} else {
				lastPrefixIndex = i
			}
		}
	}

	indent = indent - subtractor // calculates final visual indentation multiplier

	if indent < 0 { // disallow negative indentation multipliers
		indent = 0
	} else if indent > 0 { // trim the most recently displayed parent directory from the directory header to avoid displaying redundant information
		trimmedDirectory = strings.Replace(trimmedDirectory, dirList[lastPrefixIndex], "", 1)
	}

	return indent, trimmedDirectory
}

// processing for printing file entries (determines color, line wrapping, and prints)
func printFileEntry(entry string, lastSlash int, charCounter int, colorAlternator int8, indent int) (int, int8) {
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

	if charCounter == 0 { // indent first line of entries for each directory header
		fmt.Print(strings.Repeat(" ", indent*2))
	}

	// determine whether to wrap to a new line (+1 is to account for trailing spaces)
	charCounter += len(fileEntryName) + 1
	if indentation := indent * 2; charCounter+(indentation) >= width {
		charCounter = len(fileEntryName) + 1
		fmt.Print("\n" + strings.Repeat(" ", indentation)) // indent each line
	}

	// print fileEntryName to screen
	fmt.Printf("%s%s%s ", colorCode, fileEntryName, ansiReset)

	return charCounter, colorAlternator
}

// EntryListGen generates and displays full libmutton entry list
func EntryListGen() {
	fmt.Print("\n" + ansiBlackOnWhite + "libmutton entries:" + ansiReset)

	// define file/directory containing slices so that they may be accessed by the anonymous WalkDir function
	var fileList []string
	var dirList []string

	// walk entry directory
	_ = filepath.WalkDir(offline.EntryRoot,
		func(fullPath string, entry fs.DirEntry, err error) error {

			// check for errors encountered while walking directory
			if err != nil {
				// create EntryRoot if the error is the result of it not existing on the system
				if os.IsNotExist(err) {
					_ = os.Mkdir(offline.EntryRoot, 0700)
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
	dirListLength := len(dirList)                    // save length for multiple references below
	var skippedDirList = make([]bool, dirListLength) // stores whether each directory was skipped during printout (later used to determine appropriate visual indentation)
	charCounter := 0                                 // track whether to line-wrap based on character count in line
	var colorAlternator int8 = 1                     // track alternating colors for each printed entry name
	var containsSubdirectory bool                    // indicates whether the current directory contains a subdirectory
	var indent int                                   // visual indentation multiplier
	var vanityDirectory string                       // directory header printed to end-user - visual only, not used in any processing
	for i, directory := range dirList {

		// reset formatting variables for new directory
		charCounter = 0
		colorAlternator = 1

		// default to assuming this directory will be skipped (unless it is the root)
		if i == 0 {
			skippedDirList[i] = false
		} else {
			skippedDirList[i] = true
		}

		// determine directory's indentation multiplier based on PathSeparator occurrences
		indent = strings.Count(directory, offline.PathSeparator) - 1 // subtract 1 to avoid indenting root-level directories

		// check if next directory is within the current one
		if dirListLength > i+1 {
			if nextDir := dirList[i+1]; directory == nextDir[:strings.LastIndex(nextDir, offline.PathSeparator)] {
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

			// print the current file if it belongs in the current directory - otherwise, break the loop and move on to the next directory
			if lastSlash := strings.LastIndex(file, offline.PathSeparator) + 1; file[:lastSlash-1] == directory {

				// print directory header if this is the first run of the loop
				if !containsFiles {
					containsFiles = true
					skippedDirList[i] = false                                                      // the directory header is being printed, indicate that it is not being skipped
					indent, vanityDirectory = indentSubtractor(skippedDirList, dirList, i, indent) // calculate the final indentation multiplier
					if !offline.Windows {                                                          // for consistency, format directories with UNIX-style path separators on all platforms
						fmt.Printf("\n\n"+strings.Repeat(" ", indent*2)+"\033[38;5;7;48;5;8m%s/"+ansiReset+"\n", vanityDirectory)
					} else {
						fmt.Printf("\n\n"+strings.Repeat(" ", indent*2)+"\033[38;5;7;48;5;8m%s/"+ansiReset+"\n", strings.ReplaceAll(vanityDirectory, offline.PathSeparator, "/"))
					}
				}

				charCounter, colorAlternator = printFileEntry(file, lastSlash, charCounter, colorAlternator, indent)
			}
		}

		if !containsFiles { // if the current directory contains no files...
			if !containsSubdirectory { // nor does it contain any subdirectories...
				if dirListLength > 1 { // and directories besides the root-level exist... display directory header and empty directory warning
					skippedDirList[i] = false                                                      // the directory header is being printed, indicate that it is not being skipped
					indent, vanityDirectory = indentSubtractor(skippedDirList, dirList, i, indent) // calculate the final indentation multiplier
					if !offline.Windows {                                                          // for consistency, format directories with UNIX-style path separators on all platforms
						fmt.Printf("\n\n"+strings.Repeat(" ", indent*2)+"\033[38;5;7;48;5;8m%s/"+ansiReset+"\n", vanityDirectory)
					} else {
						fmt.Printf("\n\n"+strings.Repeat(" ", indent*2)+"\033[38;5;7;48;5;8m%s/"+ansiReset+"\n", strings.ReplaceAll(vanityDirectory, offline.PathSeparator, "/"))
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