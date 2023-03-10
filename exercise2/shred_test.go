package shred

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

func TestFileDoesNotExist(t *testing.T) {
	err := Shred("nonexistent_file.txt")
	if !os.IsNotExist(err) {
		t.Errorf("Expected file not found error, but got: %v", err)
	}
}

func TestShred(t *testing.T) {
	// Create a temporary file and write some data to it
	file, err := ioutil.TempFile("", "shredtest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name()) // Clean up the temporary file
	defer file.Close()

	data := []byte("this is some test data")
	_, err = file.Write(data)
	if err != nil {
		t.Fatal(err)
	}

	// Shred the file
	err = Shred(file.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Verify that the file has been deleted
	_, err = os.Stat(file.Name())
	if !os.IsNotExist(err) {
		t.Errorf("Expected file to be deleted, but it still exists")
	}
}

func TestShredLargeFile(t *testing.T) {
	// Create a temporary file with 1 GB of random data
	file, err := ioutil.TempFile("", "shredtest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name()) // Clean up the temporary file
	defer file.Close()

	// Generate random data for the file
	dataSize := 1024 * 1024 // 1 MB
	data := make([]byte, dataSize)
	_, err = rand.Read(data)
	if err != nil {
		t.Fatal(err)
	}

	// Write the random data to the file
	_, err = file.Write(data)
	if err != nil {
		t.Fatal(err)
	}

	// Shred the file
	err = Shred(file.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Verify that the file has been deleted
	_, err = os.Stat(file.Name())
	if !os.IsNotExist(err) {
		t.Errorf("Expected file to be deleted, but it still exists")
	}
}

func TestShredDir(t *testing.T) {
	// Create a temporary directory
	dir, err := ioutil.TempDir("", "shredtest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// Try to shred the directory instead of a regular file
	err = Shred(dir)
	if err == nil {
		t.Errorf("Expected an error, but Shred succeeded on directory")
	}
}

func TestShredNonRegularFile(t *testing.T) {
	// Create a temporary named pipe
	pipeName := filepath.Join(os.TempDir(), "shredtestpipe")
	err := syscall.Mkfifo(pipeName, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(pipeName)

	// Try to shred the named pipe
	err = Shred(pipeName)
	if err == nil {
		t.Errorf("Expected an error, but Shred succeeded on named pipe")
	}
}

func TestShredFilePermissions(t *testing.T) {
	// Create a temporary file with read-only permissions
	f, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	if err := f.Chmod(0400); err != nil {
		t.Fatal(err)
	}

	// Attempt to shred the file
	err = Shred(f.Name())

	// Verify that the function returned an appropriate error
	if err == nil {
		t.Errorf("Expected an error when shredding a read-only file, but no error was returned")
	}

}
