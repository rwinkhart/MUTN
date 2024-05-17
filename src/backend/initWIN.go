//go:build windows

package backend

const FallbackEditor = "nvim" // since there is no pre-installed CLI editor on Windows, default to the most popular one

// textEditorFallback returns FallbackEditor
func textEditorFallback() string {
	return FallbackEditor
}
