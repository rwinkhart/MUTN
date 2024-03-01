package offline

import (
	"os/exec"
	"strings"
)

// TODO GPG support is a temporary feature - it will be replaced with a different encryption scheme in the future
// TODO These functions may continue to exist after that point, but consider them deprecated

// DecryptGPG decrypts a GPG-encrypted file and returns the contents as a slice of (trimmed) strings
func DecryptGPG(targetLocation string) []string {
	cmd := exec.Command("gpg", "--pinentry-mode", "loopback", "-q", "-d", targetLocation)
	output, _ := cmd.CombinedOutput()
	outputSlice := strings.Split(string(output), "\n")
	for i, lineData := range outputSlice {
		outputSlice[i] = strings.TrimRight(lineData, " \t") // remove trailing whitespace
	}
	return outputSlice
}
