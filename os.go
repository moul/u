package u

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

// TempfileWithContent creates a tempfile with specified content written in it, it also seeks the file pointer so you can read it directly.
// The second returned parameter is a cleanup function that closes and removes the temp file.
func TempfileWithContent(content []byte) (*os.File, func(), error) {
	// create temp file
	tmpfile, err := ioutil.TempFile("", "u")
	if err != nil {
		return nil, nil, err
	}

	// write content
	_, err = tmpfile.Write(content)
	if err != nil {
		return nil, nil, err
	}

	// seek at the beginning of file
	_, err = tmpfile.Seek(0, io.SeekStart)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		_ = tmpfile.Close()
		_ = os.Remove(tmpfile.Name())
	}

	return tmpfile, cleanup, nil
}

// MustTempfileWithContent wraps TempfileWithContent and panics if initialization fails.
func MustTempfileWithContent(content []byte) (*os.File, func()) {
	f, cleanup, err := TempfileWithContent(content)
	if err != nil {
		panic(err)
	}
	return f, cleanup
}

func ExpandUser(path string) (string, error) {
	// expand variables
	path = os.ExpandEnv(path)

	// replace ~ with homedir
	if len(path) > 1 && path[:2] == "~/" {
		home := os.Getenv("HOME") // *nix
		if home == "" {
			home = os.Getenv("USERPROFILE") // windows
		}
		if home == "" {
			return "", errors.New("user home directory not found")
		}

		return strings.Replace(path, "~", home, 1), nil
	}

	return path, nil
}

// MustExpandUser wraps ExpandUser and panics if initialization fails.
func MustExpandUser(path string) string {
	ret, err := ExpandUser(path)
	if err != nil {
		panic(err)
	}
	return ret
}

// PathExists checks whether a path exists or not.
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// DirExists checks whether a path exists and is a directory.
func DirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

// FileExists checks whether a path exists and is a regular file.
func FileExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.Mode().IsRegular()
}

// TempFileName returns a valid temporary file name (the file is not created).
func TempFileName(dir, pattern string) (string, error) {
	tmpfile, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		return "", err
	}
	tmpfile.Close()
	os.Remove(tmpfile.Name())
	return tmpfile.Name(), nil
}

// MustTempFileName wraps TempFileName and panics if initialization fails.
func MustTempFileName(dir, pattern string) string {
	ret, err := TempFileName(dir, pattern)
	if err != nil {
		panic(err)
	}
	return ret
}

// CreateEmptyFileWithSize creates a new file of the desired size, filled with zeros.
func CreateEmptyFileWithSize(path string, size uint) error {
	fd, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create failed: %v", err)
	}
	if size > 0 {
		_, err = fd.Seek(int64(size)-1, 0)
		if err != nil {
			return fmt.Errorf("seek failed: %v", err)
		}
		_, err = fd.Write([]byte{0})
		if err != nil {
			return fmt.Errorf("write failed: %v", err)
		}
	}
	err = fd.Close()
	if err != nil {
		return fmt.Errorf("close failed: %v", err)
	}
	return nil
}

// CurrentUsename returns the current user's username.
// If username cannot be retrieved, it returns the passed fallback.
func CurrentUsername(fallback string) string {
	current, err := user.Current()
	if err == nil && current.Username != "" {
		return current.Username
	}
	if name := os.Getenv("USER"); name != "" {
		return name
	}
	return fallback
}
