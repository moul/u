package u

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/multierr"
)

// UnzipBytes is similar to Unzip but takes a zip archive as bytes instead of looking for a real file.
func UnzipBytes(src []byte, dest string) ([]string, error) {
	buf := bytes.NewReader(src)
	r, err := zip.NewReader(buf, int64(len(src)))
	if err != nil {
		return nil, err
	}
	return unzip(r, dest)
}

// Unzip decompresses a zip archive, moving all files and folders within the zip file to an output directory.
// Based on https://golangcode.com/unzip-files-in-go/ (MIT).
func Unzip(src string, dest string) ([]string, error) {
	r, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return unzip(&r.Reader, dest)
}

func unzip(r *zip.Reader, dest string) ([]string, error) {
	var (
		filenames = make([]string, 0)
		errs      error
	)
	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		// check for ZipSlip. more Info: https://snyk.io/research/zip-slip-vulnerability#go
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			errs = multierr.Append(errs, fmt.Errorf("%s: illegal file path", fpath))
			continue
		}

		if f.FileInfo().IsDir() {
			// dir
			if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
				errs = multierr.Append(errs, err)
				continue
			}
		} else {
			// file

			if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				errs = multierr.Append(errs, err)
				continue
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				errs = multierr.Append(errs, err)
				continue
			}

			rc, err := f.Open()
			if err != nil {
				errs = multierr.Append(errs, err)
				outFile.Close()
				continue
			}

			if f.Mode()&os.ModeSymlink != 0 {
				buff, _ := ioutil.ReadAll(rc)
				rc.Close()
				err = os.Symlink(string(buff), fpath)
				if err != nil {
					errs = multierr.Append(errs, err)
					continue
				}
			} else {
				_, err = io.Copy(outFile, rc)

				outFile.Close()
				rc.Close()

				if err != nil {
					errs = multierr.Append(errs, err)
					continue
				}
			}
		}
		filenames = append(filenames, fpath)
	}
	return filenames, errs
}
