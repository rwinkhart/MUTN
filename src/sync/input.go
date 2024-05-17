package sync

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

// inputKeyFilePassphrase prompts the user for a passphrase for an SSH key file
// TODO support non-CLI implementations
func inputKeyFilePassphrase() []byte {
	fmt.Print("\nEnter passphrase for your SSH keyfile: ")
	passphrase, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return passphrase
}
