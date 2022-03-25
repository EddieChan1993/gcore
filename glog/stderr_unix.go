// Log the panic under unix to the log file
//go:build darwin || unix || linux
// +build darwin unix linux

package glog

import (
	Log "log"
	"os"
	"syscall"
)

// redirectStderr to the file passed in
func redirectStderr(f *os.File) {
	err := syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		Log.Fatalf("Failed to redirect stderr to file: %v", err)
	}
}
