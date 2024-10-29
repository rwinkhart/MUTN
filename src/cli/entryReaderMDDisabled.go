//go:build noMarkdown

package cli

import (
	"fmt"
	"strings"
)

// renderNote prints the notes section of an entry to stdout (when Markdown rendering is disabled).
func renderNote(noteLines []string) {
	fmt.Print("\n" + strings.Join(noteLines, "\n") + "\n\n")
}
