package logpher

import (
	"fmt"
	"os"
	"sync"
)

// fileWriter defines a basic logger that writes to a file
type fileWriter struct {
	lock   *sync.Mutex
	closed bool
	file   *os.File
}

// newFileWriter creates a new file based logger
func newFileWriter(path string) *fileWriter {
	file, err := openFile(toAbsolutePath(path))
	panicOnError(err)

	return &fileWriter{
		lock: &sync.Mutex{},
		file: file,
	}
}

// colourEnabled determines if colour output is enabled for this writer
func (f *fileWriter) colourEnabled() bool {
	return false
}

// write writes a line to the file
func (f *fileWriter) write(line string) {
	f.lock.Lock()
	defer f.lock.Unlock()

	if f.closed {
		return
	}

	_, err := f.file.WriteString(line + "\n")
	if err != nil {
		fmt.Println("Failed to write log line:", err)
	}
}

// close closes the file writer
func (f *fileWriter) close() {
	f.lock.Lock()
	defer f.lock.Unlock()

	f.file.Close()
	f.closed = true
}
