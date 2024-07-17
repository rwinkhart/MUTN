//go:build windows

package cli

const fallbackEditor = "nvim" // since there is no pre-installed CLI editor on Windows, default to the most popular one
