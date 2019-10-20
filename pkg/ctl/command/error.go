package command

import (
	"fmt"
	"os"
)

// http://tldp.org/LDP/abs/html/exitcodes.html
const (
	ExitSuccess = iota
	ExitError
	ExitBadConnection
	ExitInterrupted
	ExitIO
	ExitBadArgs = 128
)

// ExitWithError exits with error
func ExitWithError(code int, err error) {
	fmt.Fprintln(os.Stderr, "Error:", err)
	os.Exit(code)
}

// NormalExit exits normally
func NormalExit(msg string) {
	fmt.Fprintln(os.Stdout, msg)
	os.Exit(ExitSuccess)
}
