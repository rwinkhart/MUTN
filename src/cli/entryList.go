package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/rwinkhart/go-boilerplate/back"
	"github.com/rwinkhart/go-boilerplate/other"
	"github.com/rwinkhart/libmutton/age"
	"github.com/rwinkhart/libmutton/synccommon"
)

// ANSI color constants used only in this file
const (
	ansiAlternateEntryColor = "\033[38;5;8m"
	ansiDirectoryHeader     = "\033[38;5;7;48;5;8m"
)

// determineIndentation calculates and returns the final visual indentation multiplier (needed to adjust indentation for skipped parent directories); also subtracts "old" text from directory header.
func determineIndentation(skippedDirList []bool, dirList []string, currentDirIndex int) (int, string) {
	var subtractor int      // tracks how much to subtract from expected indentation multiplier
	var lastPrefixIndex int // tracks the index (in both skippedDirList and dirList) of the last displayed parent directory
	var trimmedDirectory = dirList[currentDirIndex]

	// determine initial indentation multiplier based on "/" occurrences
	indent := strings.Count(trimmedDirectory, "/") - 1 // subtract 1 to avoid indenting root-level directories

	for i, skipped := range skippedDirList[:currentDirIndex] { // checks each skipped directory to determine if it is a parent to the current directory
		if strings.HasPrefix(trimmedDirectory, dirList[i]+"/") { // if the current directory is the child of this iteration's directory...
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

// printFileEntry handles processing for printing file entries (determines color, wraps lines, and prints).
func printFileEntry(entry string, lastSlash, charCounter, indent int, colorAlternator int8, agingTimestamp int64) (int, int8) {
	// determine color to print fileEntryName (alternate each time function is run)
	var colorCode string
	if colorAlternator > 0 {
		colorCode = ""
	} else {
		colorCode = ansiAlternateEntryColor
	}
	colorAlternator = -colorAlternator

	// determine password aging dot
	var agingDot string
	switch age.TranslateAgeTimestamp(agingTimestamp) {
	case 1:
		agingDot = back.AnsiGreen + "⁍" + back.AnsiReset
	case 2:
		agingDot = back.AnsiWarning + "⁍" + back.AnsiReset
	case 3:
		agingDot = back.AnsiError + "⁍" + back.AnsiReset
	}

	// trim the containing directory from the entry to determine fileEntryName
	fileEntryName := entry[lastSlash:]

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
	fmt.Printf("%s%s%s%s ", agingDot, colorCode, fileEntryName, back.AnsiReset)

	return charCounter, colorAlternator
}

// EntryListGen generates and displays the full libmutton entry list.
func EntryListGen() {
	fileList, dirList, err := synccommon.WalkEntryDir()
	if err != nil {
		other.PrintError("Failed to generate entry list: "+err.Error(), back.ErrorRead)
	}
	var vanityPathsToTimestamps map[string]int64
	vanityPathsToTimestamps, err = synccommon.GetEntryAges()
	if err != nil {
		other.PrintError("Failed to retrieve entry aging data: "+err.Error(), back.ErrorRead)
	}

	// print header bar w/total entry count
	fmt.Print("\n"+ansiBlackOnWhite, len(fileList), " libmutton entries:"+back.AnsiReset)

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

		// check if next directory is within the current one
		if dirListLength > i+1 {
			if nextDir := dirList[i+1]; directory == nextDir[:strings.LastIndex(nextDir, "/")] {
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
			if lastSlash := strings.LastIndex(file, "/") + 1; file[:lastSlash-1] == directory {

				// print directory header if this is the first run of the loop
				if !containsFiles {
					containsFiles = true
					skippedDirList[i] = false                                                  // the directory header is being printed, indicate that it is not being skipped
					indent, vanityDirectory = determineIndentation(skippedDirList, dirList, i) // calculate the final indentation multiplier
					fmt.Printf("\n\n"+strings.Repeat(" ", indent*2)+ansiDirectoryHeader+"%s/"+back.AnsiReset+"\n", vanityDirectory)
				}

				charCounter, colorAlternator = printFileEntry(file, lastSlash, charCounter, indent, colorAlternator, vanityPathsToTimestamps[file])
			}
		}

		if !containsFiles { // if the current directory contains no files...
			if !containsSubdirectory { // nor does it contain any subdirectories...
				if dirListLength > 1 { // and directories besides the root-level exist... display directory header and empty directory warning
					skippedDirList[i] = false                                                  // the directory header is being printed, indicate that it is not being skipped
					indent, vanityDirectory = determineIndentation(skippedDirList, dirList, i) // calculate the final indentation multiplier
					fmt.Printf("\n\n"+strings.Repeat(" ", indent*2)+ansiDirectoryHeader+"%s/"+back.AnsiReset+"\n", vanityDirectory)
					fmt.Print(strings.Repeat(" ", indent*2) + back.AnsiWarning + "-empty directory-" + back.AnsiReset)
				} else { // warn if the only thing that exists is the root-level directory
					fmt.Print("\n\nNothing's here! For help creating your first entry, run \"mutn help\".")
				}
			}
		}
	}

	// print trailing new lines for proper spacing after entry list is complete
	fmt.Print("\n\n")

	os.Exit(0)
}
