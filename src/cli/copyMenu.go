package cli

import (
	"fmt"
	"os"
	"os/signal"
	"slices"
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

// CopyMenu decrypts an entry and allows the user to
// interactively copy fields without having to re-decrypt each time.
// decSlice can be left nil to decrypt the entry specified by vanityPath.
func CopyMenu(vanityPath string, decSlice []string, oldPassword string) {
	realPath := global.GetRealPath(vanityPath)
	var err error
	if decSlice == nil {
		// decrypt entry
		decSlice, err = crypt.DecryptFileToSlice(realPath)
		if err != nil {
			other.PrintError("Failed to decrypt entry: "+err.Error(), global.ErrorDecryption)
		}
	} else {
		fmt.Print("\n\n")
		_, err := syncclient.RunJob(false)
		if err != nil {
			other.PrintError("Failed to sync on copy menu entry: "+err.Error(), global.ErrorSyncProcess)
		}
	}

	// determine populated fields in entry
	var fieldStrings = []string{"Username", "Password", "TOTP Code", "URL", "Note (first line)"}
	var indices = []int{1, 0, 2, 3, 4}
	var fieldOptions []string
	for i := range indices {
		if len(decSlice) > indices[i] && decSlice[indices[i]] != "" {
			fieldOptions = append(fieldOptions, fieldStrings[i])
			if indices[i] == 0 && oldPassword != "" {
				fieldOptions = append(fieldOptions, "Old Password")
			}
		}
	}

	// if no copyable lines are populated, render notes and exit
	if len(fieldOptions) < 1 {
		EntryReader(vanityPath, decSlice, true)
		fmt.Printf("\r%sNo copyable fields present, exiting copy menu...%s\n", back.AnsiWarning, back.AnsiReset)
		os.Exit(0)
	}

	// set up signal handling for ctrl+c to clear clipboard
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		if err = clip.ClearProcess(""); err != nil {
			other.PrintError("Failed to clear clipboard on exit", global.ErrorClipboard)
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
	var choice int
	var selectedField string
	for {
		fmt.Println()
		if selectedField != "" {
			choice = front.InputMenuGen("Field to copy:", fieldOptions)
		} else {
			choice = front.InputMenuGen("Field to copy (exiting in 5 seconds):", fieldOptions)
			selectedChan <- true
		}
		selectedField = fieldOptions[choice-1]
		if selectedField == "Old Password" {
			err = clip.CopyString(false, oldPassword)
			if err != nil {
				other.PrintError("Failed to copy old password to clipboard: "+err.Error(), global.ErrorClipboard)
			}
			continue
		}
		choice = indices[slices.Index(fieldStrings, selectedField)]
		if choice == 2 {
			fmt.Println(back.AnsiWarning + "[Starting]" + back.AnsiReset + " TOTP clipboard refresher")
			errorChan := make(chan error)
			done := make(chan bool)
			go clip.TOTPCopier(decSlice[2], errorChan, done)
			// block until first successful copy
			err = <-errorChan
			if err != nil {
				other.PrintError("Failed to copy TOTP token: "+err.Error(), global.ErrorClipboard)
			}
			fmt.Println(back.AnsiGreen + "[Started]" + back.AnsiReset + " TOTP clipboard refresher")
			// block until the user sends "q"
			for {
				input := front.Input("Service will run until 'q' is entered:")
				if input == "q" {
					break
				}
			}
			close(done)
			fmt.Println(back.AnsiBlue + "\n[Stopped]" + back.AnsiReset + " TOTP clipboard refresher")
		} else {
			if err = clip.CopyString(false, decSlice[choice]); err != nil {
				other.PrintError("Failed to copy field to clipboard: "+err.Error(), global.ErrorClipboard)
			}
		}
	}
}
