package cli

import (
	"bufio"
	"fmt"
	"github.com/rwinkhart/MUTN/src/offline"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/exec"
)

// input prompts the user for input and returns the input as a string
func input(prompt string) string {
	fmt.Print("\n" + prompt + " ")
	var input string
	fmt.Scanln(&input)
	return input
}

// inputHidden prompts the user for input and returns the input as a string, hiding the input from the terminal
func inputHidden(prompt string) string {
	fmt.Print("\n" + prompt + " ")
	byteInput, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	password := string(byteInput)
	fmt.Println()
	return password
}

// inputBinary prompts the user with a yes/no question and returns the response as a boolean
func inputBinary(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n" + prompt + " (y/N) ")
	char, _, _ := reader.ReadRune()
	if char == 'y' || char == 'Y' {
		return true
	}
	return false
}

// newNote uses the user-specified text editor to create a new note and returns the note as a slice of strings
func newNote() []string {
	tempFile := offline.CreateTempFile()
	defer os.Remove(tempFile.Name())
	editor := offline.ReadConfig([]string{"textEditor"})[0]
	cmd := exec.Command(editor, tempFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(offline.AnsiError + "Failed to write note with " + editor + offline.AnsiReset)
		os.Exit(1)
	}

	file, err := os.Open(tempFile.Name())
	if err != nil {
		fmt.Println(offline.AnsiError + "Failed to open temporary file (\"" + tempFile.Name() + "\") " + err.Error() + offline.AnsiReset)
		os.Exit(1)
	}
	defer file.Close()

	var note []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		note = append(note, scanner.Text())
	}

	return offline.RemoveTrailingEmptyStrings(note)
}
