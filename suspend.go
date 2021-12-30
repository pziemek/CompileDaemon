package main

import (
	"os"
)

// checkIfSuspended return true if not empty
func checkIfSuspended() bool {
	cnt, err := os.ReadFile(*flagSuspend)
	if err != nil {
		return false
	}
	if len(cnt) > 0 {
		return true
	}
	return false
}
