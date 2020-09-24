package u

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
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
