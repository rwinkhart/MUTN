//go:build !windows

package backend

import (
	"os"
)

const FallbackEditor = "vi" // vi is pre-installed on most UNIX systems

// textEditorFallback returns the value of the $EDITOR environment variable, or FallbackEditor if it is not set
func textEditorFallback() string {
	// ensure textEditor is set
	textEditor := os.Getenv("EDITOR")
	if textEditor == "" {
		textEditor = FallbackEditor
	}
	return textEditor
}
