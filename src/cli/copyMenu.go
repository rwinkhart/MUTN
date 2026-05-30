package cli

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/rwinkhart/go-boilerplate/back"
	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
	"github.com/rwinkhart/libmutton/clip"
	"github.com/rwinkhart/libmutton/crypt"
	"github.com/rwinkhart/libmutton/global"
	"github.com/rwinkhart/libmutton/syncclient"
)

// CopyMenu decrypts an entry and allows the user to interactively copy
// fields without having to re-decrypt each time. decSlice can be left nil
// to decrypt the entry specified by vanityPath. If the entry was provided
// decrypted, CopyMenu assumes it was edited and triggers a sync.
func CopyMenu(vanityPath string, decSlice []string, oldPassword []byte) {
	realPath := global.GetRealPath(vanityPath)
	var err error
	if decSlice == nil {
		// decrypt entry
		decSlice, err = crypt.DecryptFileToSlice(realPath, nil)
		if err != nil {
			other.PrintError("Failed to decrypt entry: "+err.Error(), global.ErrorDecryption)
		}
	} else {
		fmt.Print("\n")
		_, err := syncclient.RunJob()
		if err != nil {
			other.PrintError("Failed to sync on copy menu entry: "+err.Error(), global.ErrorSyncProcess)
		}
	}

	// determine populated fields in entry
	var allStrings = [5]string{"Username", "Password", "TOTP Code", "URL", "Note (first line)"}
	var allIndices = [5]int{1, 0, 2, 3, 4}
	var allRunes = [5]rune{'u', 'p', 't', 'l', 'n'}
	var fieldStrings []string
	var fieldRunes []rune
	for i := range allIndices {
		if len(decSlice) > allIndices[i] && decSlice[allIndices[i]] != "" {
			fieldStrings = append(fieldStrings, allStrings[i])
			fieldRunes = append(fieldRunes, allRunes[i])
			if allIndices[i] == 0 && oldPassword != nil {
				fieldStrings = append(fieldStrings, "Old Password")
				fieldRunes = append(fieldRunes, 'b')
			}
		}
	}

	// if notes are included, preview them
	if len(decSlice) > 4 {
		EntryReader(vanityPath, append([]string{"", "", "", ""}, decSlice[4:]...), true)
		if len(fieldStrings) < 1 { // if notes are the only thing included, exit
			fmt.Printf("\r%sNo copyable fields present, exiting copy menu...%s\n", back.AnsiWarning, back.AnsiReset)
			os.Exit(0)
		}
	}

	// set up signal handling for ctrl+c to clear clipboard
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		if err = clip.ClearProcess(nil); err != nil {
			other.PrintError("Failed to clear clipboard on exit: "+err.Error(), global.ErrorClipboard)
		}
		os.Exit(0)
	}()

	// set up timer to close copy menu if an item is
	// not selected (useful to avoid extra user input
	// when the copy menu is displayed automatically
	// and is not desired)
	const timeoutSeconds uint8 = 5
	selectedChan := make(chan bool, 1)
	go func() {
		for i := uint8(1); i <= timeoutSeconds; i++ {
			select {
			case <-selectedChan:
				return
			case <-time.After(1 * time.Second):
				if i == timeoutSeconds {
					fmt.Printf("\r%sNo field selected, exiting copy menu...%s\n", back.AnsiWarning, back.AnsiReset)
					os.Exit(0)
				}
				fmt.Printf("\rField to copy (exiting in %d seconds): ", timeoutSeconds-i)
			}
		}
	}()

	// copy selected field to clipboard
	var runesToIndices = map[rune]int{'u': 1, 'p': 0, 't': 2, 'l': 3, 'n': 4}
	var selectedField rune
	for {
		fmt.Println()
		if selectedField != 0 {
			selectedField = front.InputMenuGenWithRuneInputs("Field to copy:", fieldStrings, fieldRunes)
		} else {
			selectedField = front.InputMenuGenWithRuneInputs("Field to copy (exiting in 5 seconds):", fieldStrings, fieldRunes)
			selectedChan <- true
		}

		switch selectedField {
		case 'b': // old password
			if err = clip.CopyBytes(false, oldPassword); err != nil {
				other.PrintError("Failed to copy old password to clipboard: "+err.Error(), global.ErrorClipboard)
			}
		case 't': // TOTP
			fmt.Println(back.AnsiWarning + "[Starting]" + back.AnsiReset + " TOTP clipboard refresher")
			errorChan := make(chan error, 1)
			done := make(chan bool)
			go clip.TOTPCopier(decSlice[2], errorChan, done)
			// block until first successful copy
			if err = <-errorChan; err != nil {
				other.PrintError("Failed to copy TOTP token: "+err.Error(), global.ErrorClipboard)
			}
			// block until the user sends "q"
			for {
				input := front.Input("Service will run until 'q' is entered:")
				if input == "q" {
					break
				}
			}
			close(done)
			fmt.Println(back.AnsiBlue + "[Stopped]" + back.AnsiReset + " TOTP clipboard refresher")
		case 'n': // notes
			decSlice[4] = strings.TrimRight(decSlice[4], " ")
			fallthrough
		default:
			if err = clip.CopyBytes(false, []byte(decSlice[runesToIndices[selectedField]])); err != nil {
				other.PrintError("Failed to copy field to clipboard: "+err.Error(), global.ErrorClipboard)
			}
		}
	}
}
