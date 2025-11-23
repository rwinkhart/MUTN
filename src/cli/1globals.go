package cli

import (
	"os"

	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/libmutton/global"
	"golang.org/x/term"
)

var (
	width, _, _ = term.GetSize(int(os.Stdout.Fd()))
)

const (
	MUTNVersion = "0.D.0" // untagged releases feature a letter suffix corresponding to the eventual release version, e.g "0.B.0" -> "0.2.0", "0.2.A" -> "0.2.1"

	ansiBlackOnWhite = "\033[38;5;0;48;5;15m"
)

func init() {
	global.GetPassword = front.InputHidden
}
