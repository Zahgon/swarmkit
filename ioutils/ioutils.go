package ioutils

import (
	"os"
)

// todo: split docker/pkg/ioutils into a separate repo

// AtomicWriteFile atomically writes data to a file specified by filename.
func AtomicWriteFile(filename string, data []byte, perm os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}
