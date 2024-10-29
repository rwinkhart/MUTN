//go:build !noMarkdown && !BEAN

package cli

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/glamour"
)

// renderNote renders the notes section of an entry (in Markdown) to stdout.
func renderNote(noteLines []string) {
	note := strings.Join(noteLines, "\n")
	r, _ := glamour.NewTermRenderer(glamour.WithStylesFromJSONBytes(glamourStyle()), glamour.WithPreservedNewLines(), glamour.WithWordWrap(width))
	markdownNotesString, _ := r.Render(note)

	// print markdown-rendered notes
	fmt.Print(markdownNotesString)
}
