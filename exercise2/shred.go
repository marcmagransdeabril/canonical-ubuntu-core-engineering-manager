package shred

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var ErrPermission = errors.New("permission denied")

func Shred(path string) error {
	const BUFFER_SIZE = 4096 // 4k buffer size

	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	// Check that the path points to a regular file
	if !fileInfo.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", path)
	}

	// Check file permissions
	if fileInfo.Mode().Perm()&0600 == 0 {
		return ErrPermission
	}
	// Open
	file, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	fileSize := fileInfo.Size()
	buffer := make([]byte, BUFFER_SIZE)

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 3; i++ {

		// Write random data to the file using a limited-size buffer
		offset := int64(0)

		for offset < fileSize {

			// Generate random data and write to buffer
			n, err := rand.Read(buffer)
			if err != nil {
				return err
			}

			// Write buffer to file
			_, err = file.WriteAt(buffer[:n], offset)
			if err != nil {
				return err
			}

			// Update the offset and repeat
			offset += int64(n)
		}

		// Flush file data to disk
		err := file.Sync()
		if err != nil {
			return err
		}
	}

	// Remove the file
	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
