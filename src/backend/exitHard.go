//go:build !returnOnExit

package backend

import "os"

func Exit(code int) {
	os.Exit(code)
}
