package u

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ExpandPath performs various expansions on a given path.
//
// - Replaces ~/ with $HOME/.
// - Returns absolute path.
// - Expands env vars.
// TODO: - Follow symlinks.
func ExpandPath(path string) (string, error) {
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

	// compute absolute path
	{
		result, err := filepath.Abs(path)
		if err != nil {
			return "", fmt.Errorf("absolute path: %q: %w", path, err)
		}
		path = result
	}

	// expand env vars
	{
		path = os.ExpandEnv(path)
	}

	// eval symlinks
	// TODO: do not fail if path does not exist
	/*
		{
			result, err := filepath.EvalSymlinks(path)
			if err != nil {
				return "", fmt.Errorf("eval symlinks: %q: %w", path, err)
			}
			path = result
		}
	*/

	return path, nil
}

// MustExpandPath wraps ExpandPath and panics if initialization fails.
func MustExpandPath(path string) string {
	ret, err := ExpandPath(path)
	if err != nil {
		panic(err)
	}
	return ret
}
