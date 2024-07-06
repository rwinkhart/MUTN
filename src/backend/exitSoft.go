//go:build returnOnExit

package backend

func Exit(code int) {
	return code
}
