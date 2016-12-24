package file

import (
	"io/ioutil"
	"os"
)

// Function to read the contents of a file and return contents as array of bytes
// with error as nil. It will return error if error occur while reading
func Read(path string) ([]byte, error) {
	var (
		content []byte
		err     error
	)
	if content, err = ioutil.ReadFile(path); err != nil {
		return content, err
	}
	return content, nil
}

// Creates a new file and returns the file pointer with error as nil if the
// file is succesfully created else return file pointer as nil with the error
func Create(filePath string) (*os.File, error) {
	newFile, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	return newFile, nil
}

// Close the file
func Close(f *os.File) error {
	err := f.Close()
	if err != nil {
		return err
	}
	return nil
}

// Function to write array of bytes to the file and return number of bytes
// written with error as nil if successful, otherwise it will return 0 with error
func Write(f *os.File, data []byte) (int, error) {
	var (
		n   int
		err error
	)
	n, err = f.Write(data)
	if err != nil {
		return 0, err
	}

	return n, nil
}

// Function to issue a Sync to flush writes to stable storage
// It will return error as nil if successful otherwise error
func Sync(f *os.File) error {
	var err error
	err = f.Sync()
	if err != nil {
		return err
	}
	return nil
}

// Open an existing file and returns the file pointer with error as nil if the
// file is succesfully opened else return file pointer as nil with the error
func Open(filePath string) (*os.File, error) {
	newFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}
	return newFile, nil
}
