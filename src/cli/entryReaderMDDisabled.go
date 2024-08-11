//go:build noMarkdown

package cli

import (
	"fmt"
)

// renderNote prints the notes section of an entry to stdout (when Markdown rendering is disabled).
func renderNote(note *string) {
	fmt.Print("\n" + *note + "\n\n")
}
