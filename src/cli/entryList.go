package cli

import (
	"fmt"
	"os"
	"sort"
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

func writeIndent(sb *strings.Builder, indent int) { sb.WriteString(strings.Repeat(" ", indent*2)) }

// determineIndentation calculates the final visual indentation multiplier (adjusts for skipped parents)
// and trims the parent directory name from the directory header (if applicable).
func determineIndentation(skippedDirList []bool, dirList []string, currentDirIndex int) (int, string) {
	trimmed := dirList[currentDirIndex]
	indent := strings.Count(trimmed, "/") - 1 // avoid indenting root-level dirs

	subtractor, lastPrefixIndex := 0, 0
	for i, skipped := range skippedDirList[:currentDirIndex] {
		if strings.HasPrefix(trimmed, dirList[i]+"/") {
			if skipped {
				subtractor++
			} else {
				lastPrefixIndex = i
			}
		}
	}
	if indent -= subtractor; indent < 0 {
		return 0, trimmed
	}
	if indent > 0 {
		trimmed = strings.TrimPrefix(trimmed, dirList[lastPrefixIndex])
	}
	return indent, trimmed
}

// printFileEntry handles processing for printing file entries (determines color, adds aging indicator, wraps lines, and prints).
func printFileEntry(sb *strings.Builder, entry string, lastSlash int, charCounter *int, indent int, colorAlternator int8, agingTimestamp *int64) int8 {
	colorCode := ""
	if colorAlternator <= 0 {
		colorCode = ansiAlternateEntryColor
	}
	colorAlternator = -colorAlternator

	agingDot := ""
	switch age.TranslateAgeTimestamp(agingTimestamp) {
	case 1:
		agingDot = back.AnsiGreen + "⁍" + back.AnsiReset
	case 2:
		agingDot = back.AnsiWarning + "⁍" + back.AnsiReset
	case 3:
		agingDot = back.AnsiError + "⁍" + back.AnsiReset
	}

	name := entry[lastSlash:]
	if *charCounter == 0 {
		writeIndent(sb, indent)
	}

	// wrap if needed; +1 accounts for trailing space, +1 more if aging dot is present
	inc := len(name) + 1
	if agingDot != "" {
		inc++
	}
	if *charCounter += inc; *charCounter+indent*2 >= width {
		*charCounter = len(name) + 1
		if agingDot != "" {
			*charCounter++
		}
		sb.WriteString("\n")
		writeIndent(sb, indent)
	}

	sb.WriteString(agingDot)
	sb.WriteString(colorCode)
	sb.WriteString(name)
	sb.WriteString(back.AnsiReset)
	sb.WriteString(" ")
	return colorAlternator
}

// EntryListGen generates and displays the full libmutton entry list.
func EntryListGen() {
	_, dirList, err := synccommon.WalkEntryDir()
	if err != nil {
		other.PrintError("Failed to generate entry list: "+err.Error(), back.ErrorRead)
	}
	entryMap, err := synccommon.GetAllEntryData()
	if err != nil {
		other.PrintError("Failed to retrieve entry aging data: "+err.Error(), back.ErrorRead)
	}

	var sb strings.Builder

	// print header bar w/total entry count
	sb.WriteString("\n")
	sb.WriteString(ansiBlackOnWhite)
	fmt.Fprint(&sb, len(entryMap))
	sb.WriteString(" libmutton entries:")
	sb.WriteString(back.AnsiReset)

	entriesByDir := make(map[string][]string, len(dirList))
	for vanityPath := range entryMap {
		if lastSlash := strings.LastIndex(vanityPath, "/"); lastSlash >= 0 {
			entriesByDir[vanityPath[:lastSlash]] = append(entriesByDir[vanityPath[:lastSlash]], vanityPath)
		}
	}
	for _, entries := range entriesByDir {
		sort.Strings(entries)
	}

	// precompute whether each directory contains subdirectories;
	// handle root-level child dirs by treating "" as their parent.
	hasChildDir := make(map[string]bool, len(dirList))
	for _, d := range dirList {
		switch parentEnd := strings.LastIndex(d, "/"); {
		case parentEnd > 0:
			hasChildDir[d[:parentEnd]] = true
		case parentEnd == -1:
			hasChildDir[""] = true
		}
	}

	dirListLength := len(dirList)
	skippedDirList := make([]bool, dirListLength)
	charCounter := 0
	var colorAlternator int8 = 1

	printDirHeader := func(i int) int {
		skippedDirList[i] = false
		indent, vanityDirectory := determineIndentation(skippedDirList, dirList, i)
		sb.WriteString("\n\n")
		writeIndent(&sb, indent)
		sb.WriteString(ansiDirectoryHeader)
		sb.WriteString(vanityDirectory)
		sb.WriteString("/" + back.AnsiReset + "\n")
		return indent
	}

	for i, directory := range dirList {
		charCounter, colorAlternator = 0, 1
		skippedDirList[i] = i != 0

		if entries := entriesByDir[directory]; len(entries) > 0 {
			indent := printDirHeader(i)
			for _, vanityPath := range entries {
				lastSlash := strings.LastIndex(vanityPath, "/") + 1
				colorAlternator = printFileEntry(&sb, vanityPath, lastSlash, &charCounter, indent, colorAlternator, entryMap[vanityPath].AgeTimestamp)
			}
			continue
		}

		if hasChildDir[directory] {
			continue
		}

		if dirListLength > 1 {
			indent := printDirHeader(i)
			writeIndent(&sb, indent)
			sb.WriteString(back.AnsiWarning + "-empty directory-" + back.AnsiReset)
			continue
		}

		sb.WriteString("\n\nNothing's here! For help creating your first entry, run \"mutn help\".")
	}

	sb.WriteString("\n\n")
	fmt.Print(sb.String())
	os.Exit(0)
}
