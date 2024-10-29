//go:build !noMarkdown && BEAN

package cli

import (
	"fmt"

	bean "github.com/Trojan2021/BEAN/render"
)

// renderNote renders the notes section of an entry (in Markdown) to stdout.
func renderNote(noteLines []string) {
	fmt.Println(bean.RenderMarkdown(noteLines, width))
}
