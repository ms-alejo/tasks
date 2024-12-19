package cmd

import (
	"fmt"
	"os"

	"github.com/gofrs/flock"
)

var lock *flock.Flock

// opens a file and uses a lock to prevent concurrent access
func loadFile(filepath string) (*os.File, error) {
	lock = flock.New(filepath + ".lock") // create a file lock

	// acquire lock
	locked, err := lock.TryLock()
	if err != nil {
		return nil, fmt.Errorf("failed to acquire file lock: %v", err)
	}
	if !locked {
		return nil, fmt.Errorf("file is already locked by another process")
	}

	// open file or create when it doesn't exist
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		lock.Unlock() // release on error
		return nil, fmt.Errorf("failed to open file: %v", err)
	}

	return f, nil
}

func closeFile(f *os.File) error {
	// unlock file and close
	if lock != nil {
		lock.Unlock()
	}
	return f.Close()
}
