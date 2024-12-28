package loader

import "io"

type ReadWriteResetCloser interface {
	io.ReadWriteCloser

	// Reset truncates the file and seeks to the beginning of the file.
	Reset() error
}

type Loader interface {
	Load() ([]ReadWriteResetCloser, error)
}
