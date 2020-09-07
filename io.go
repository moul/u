package u

import (
	"io"
	"io/ioutil"
	"os"
)

// SilentClose calls an io.Closer.Close() function and ignore potential errors.
//
// You can use it as `defer SilenceClose(f)`
func SilentClose(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}

// CaptureStdout temporarily pipes os.Stdout into a buffer.
func CaptureStdout() (func() string, error) {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	os.Stdout = w

	closer := func() string {
		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = old
		return string(out)
	}
	return closer, nil
}

// CaptureStderr temporarily pipes os.Stderr into a buffer.
func CaptureStderr() (func() string, error) {
	old := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	os.Stderr = w

	closer := func() string {
		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = old
		return string(out)
	}
	return closer, nil
}

// CaptureStdoutAndStderr temporarily pipes os.Stdout and os.Stderr into a buffer.
func CaptureStdoutAndStderr() (func() string, error) {
	oldErr := os.Stderr
	oldOut := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	os.Stderr = w
	os.Stdout = w

	closer := func() string {
		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stderr = oldErr
		os.Stdout = oldOut
		return string(out)
	}
	return closer, nil
}

// MustCaptureStdout wraps CaptureStdout and panics if initialization fails.
func MustCaptureStdout() func() string {
	closer, err := CaptureStdout()
	if err != nil {
		panic(err)
	}
	return closer
}

// MustCaptureStderr wraps CaptureStderr and panics if initialization fails.
func MustCaptureStderr() func() string {
	closer, err := CaptureStderr()
	if err != nil {
		panic(err)
	}
	return closer
}

// MustCaptureStdoutAndStderr wraps CaptureStdoutAndStderr and panics if initialization fails.
func MustCaptureStdoutAndStderr() func() string {
	closer, err := CaptureStdoutAndStderr()
	if err != nil {
		panic(err)
	}
	return closer
}
