package u_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"moul.io/u"
)

func TestExpandPath(t *testing.T) {
	os.Setenv("HOME", "/home/foo")
	os.Setenv("USER", "foo")
	os.Unsetenv("FOOBAR")
	workdir, err := os.Getwd()
	require.NoError(t, err)

	tests := []struct {
		input       string
		expected    string
		shouldFails bool
	}{
		{"/home/foo", "/home/foo", false},
		{"/home/foo/", "/home/foo", false},
		{"~", filepath.Join(workdir, "~"), false},
		{"~/", "/home/foo/", false},
		{"$HOME/hello", "/home/foo/hello", false},
		{"/home/$USER/hello", "/home/foo/hello", false},
		{"/tmp/$FOOBAR/hello", "/tmp/hello", false},
	}
	for _, tc := range tests {
		name := strings.Replace(tc.input, "/", "-", -1)
		t.Run(name, func(t *testing.T) {
			result, err := u.ExpandPath(tc.input)
			if tc.shouldFails {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, result, tc.expected)
			}
		})
	}
}

func ExampleExpandPath() {
	os.Setenv("HOME", "/home/foo") // just for example
	ret, err := u.ExpandPath("~/hello-world.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
	// Output: /home/foo/hello-world.txt
}

func ExampleMustExpandPath() {
	os.Setenv("HOME", "/home/foo") // just for example
	ret := u.MustExpandPath("~/hello-world.txt")
	fmt.Println(ret)
	// Output: /home/foo/hello-world.txt
}
