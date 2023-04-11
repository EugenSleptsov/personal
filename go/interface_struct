package main

import (
	"fmt"
	"io"
	"os"
)

// ReaderWriter is an interface that combines the io.Reader and io.Writer interfaces.
type ReaderWriter interface {
	io.Reader
	io.Writer
}

// FileProcessor is a struct that will process a file.
type FileProcessor struct {
	filepath string
}

// Process accepts an interface (ReaderWriter) and performs reading and writing operations.
func (fp *FileProcessor) Process(rw ReaderWriter) error {
	// Perform read and write operations using the ReaderWriter interface.
	// ...
	return nil
}

// NewFileProcessor is a constructor that returns a concrete type (FileProcessor).
func NewFileProcessor(filepath string) *FileProcessor {
	return &FileProcessor{filepath: filepath}
}

func main() {
	// Create a new FileProcessor instance.
	fp := NewFileProcessor("example.txt")

	// Open a file that implements the ReaderWriter interface.
	file, err := os.OpenFile("example.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Call the Process method with the file (which implements the ReaderWriter interface).
	err = fp.Process(file)
	if err != nil {
		fmt.Println("Error processing file:", err)
		return
	}

	fmt.Println("File processing completed successfully.")
}
