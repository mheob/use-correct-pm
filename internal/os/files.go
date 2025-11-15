// Package os holds some OS related functions
package os

import "os"

// FileExists checks if a file exists
func FileExists(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}
