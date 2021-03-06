package u_test

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"moul.io/u"
)

func ExampleUnzip() {
	// create zipfile on fs
	f, cleanup, err := u.TempfileWithContent(zipdata_simple)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// create tempdir for dest
	tempdir, err := ioutil.TempDir("", "u")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tempdir)

	// unzip to dest
	files, err := u.Unzip(f.Name(), tempdir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		relPath := "." + strings.TrimPrefix(file, tempdir)
		fmt.Println(relPath)
	}
	// Output:
	// ./test.txt
	// ./gophercolor16x16.png
}

func ExampleUnzipBytes() {
	// create tempdir for dest
	tempdir, err := ioutil.TempDir("", "u")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tempdir)

	// unzip to dest
	files, err := u.UnzipBytes(zipdata_simple, tempdir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		relPath := "." + strings.TrimPrefix(file, tempdir)
		fmt.Println(relPath)
	}
	// Output:
	// ./test.txt
	// ./gophercolor16x16.png
}

func ExampleUnzip_zipslip() {
	// create zipfile on fs
	f, cleanup, err := u.TempfileWithContent(zipdata_zipline)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// create tempdir for dest
	tempdir, err := ioutil.TempDir("", "u")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tempdir)

	// unzip to dest
	_, err = u.Unzip(f.Name(), tempdir)
	fmt.Println(err)

	// Output:
	// /tmp/evil.txt: illegal file path
}

func ExampleUnzipBytes_createAndUnzip() {
	// create a custom zip
	buf := new(bytes.Buffer)
	{
		w := zip.NewWriter(buf)
		// a.txt
		{
			hdr := zip.FileHeader{
				Name:   "a.txt",
				Method: zip.Deflate,
			}
			hdr.SetMode(0o755)
			f, err := w.CreateHeader(&hdr)
			if err != nil {
				panic(err)
			}
			_, err = f.Write([]byte("hello world!"))
			if err != nil {
				panic(err)
			}
		}

		// b.txt -> a.txt
		{
			hdr := zip.FileHeader{
				Name:    "b.txt",
				Comment: "c",
				Method:  zip.Deflate,
			}
			hdr.SetMode(0o755 | os.ModeSymlink)
			f, err := w.CreateHeader(&hdr)
			if err != nil {
				panic(err)
			}
			_, err = f.Write([]byte("a.txt"))
			if err != nil {
				panic(err)
			}
		}

		err := w.Close()
		if err != nil {
			panic(err)
		}
	}

	// unzip it
	{
		// create tempdir for dest
		tempdir, err := ioutil.TempDir("", "u")
		if err != nil {
			panic(err)
		}
		defer os.RemoveAll(tempdir)

		// unzip to dest
		files, err := u.UnzipBytes(buf.Bytes(), tempdir)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			relPath := "." + strings.TrimPrefix(file, tempdir)
			stat, err := os.Lstat(file)
			if err != nil {
				panic(err)
			}
			fmt.Println(relPath, stat.Mode()&os.ModeSymlink != 0)
		}
	}
	// Output:
	// ./a.txt false
	// ./b.txt true
}

func ExampleUnzipBytes_createAndUnzipZipSlip() {
	// create second temp dir
	var victim string
	{
		var err error
		victim, err = ioutil.TempDir("", "u")
		if err != nil {
			panic(err)
		}
		defer os.RemoveAll(victim)
	}

	// create a custom zip
	buf := new(bytes.Buffer)
	{
		w := zip.NewWriter(buf)
		// a.txt
		{
			hdr := zip.FileHeader{
				Name:   "a",
				Method: zip.Deflate,
			}
			hdr.SetMode(0o755 | os.ModeSymlink)
			f, err := w.CreateHeader(&hdr)
			if err != nil {
				panic(err)
			}
			_, err = f.Write([]byte(victim))
			if err != nil {
				panic(err)
			}
		}

		// b.txt -> a.txt
		{
			hdr := zip.FileHeader{
				Name:   "a/b.txt",
				Method: zip.Deflate,
			}
			hdr.SetMode(0o755)
			f, err := w.CreateHeader(&hdr)
			if err != nil {
				panic(err)
			}
			_, err = f.Write([]byte("hello world!"))
			if err != nil {
				panic(err)
			}
		}

		err := w.Close()
		if err != nil {
			panic(err)
		}
	}

	// unzip it
	{
		// create tempdir for dest
		tempdir, err := ioutil.TempDir("", "u")
		if err != nil {
			panic(err)
		}
		defer os.RemoveAll(tempdir)

		// unzip to dest
		_, err = u.UnzipBytes(buf.Bytes(), tempdir)
		errStr := err.Error()
		errStr = strings.Replace(errStr, tempdir, "TEMPDIR", -1)
		fmt.Println(errStr)
	}
	// Output:
	// TEMPDIR/a/b.txt: illegal file path
}

var zipdata_simple = []byte{
	0x50, 0x4b, 0x03, 0x04, 0x14, 0x00, 0x00, 0x00, 0x08, 0x00, 0x81, 0x61, 0x25, 0x3d, 0xc0, 0xd7,
	0xed, 0xc3, 0x19, 0x00, 0x00, 0x00, 0x1a, 0x00, 0x00, 0x00, 0x08, 0x00, 0x1c, 0x00, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x74, 0x78, 0x74, 0x55, 0x54, 0x09, 0x00, 0x03, 0x71, 0xfc, 0x82, 0x4c, 0x76,
	0xfc, 0x82, 0x4c, 0x75, 0x78, 0x0b, 0x00, 0x01, 0x04, 0xf5, 0x01, 0x00, 0x00, 0x04, 0x14, 0x00,
	0x00, 0x00, 0x0b, 0xc9, 0xc8, 0x2c, 0x56, 0x00, 0xa2, 0x44, 0x85, 0x92, 0xd4, 0xe2, 0x12, 0x20,
	0x51, 0x51, 0xa2, 0x90, 0x96, 0x99, 0x93, 0xaa, 0xc7, 0x05, 0x00, 0x50, 0x4b, 0x03, 0x04, 0x0a,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x9d, 0x7e, 0x25, 0x3d, 0xfe, 0x31, 0xd5, 0x54, 0x11, 0x03, 0x00,
	0x00, 0x11, 0x03, 0x00, 0x00, 0x14, 0x00, 0x1c, 0x00, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x31, 0x36, 0x78, 0x31, 0x36, 0x2e, 0x70, 0x6e, 0x67, 0x55, 0x54, 0x09,
	0x00, 0x03, 0x3a, 0x30, 0x83, 0x4c, 0x3b, 0x30, 0x83, 0x4c, 0x75, 0x78, 0x0b, 0x00, 0x01, 0x04,
	0xf5, 0x01, 0x00, 0x00, 0x04, 0x14, 0x00, 0x00, 0x00, 0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a,
	0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00,
	0x0f, 0x08, 0x06, 0x00, 0x00, 0x00, 0xed, 0x73, 0x4f, 0x2f, 0x00, 0x00, 0x00, 0x04, 0x67, 0x41,
	0x4d, 0x41, 0x00, 0x00, 0xd6, 0xd8, 0xd4, 0x4f, 0x58, 0x32, 0x00, 0x00, 0x00, 0x19, 0x74, 0x45,
	0x58, 0x74, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x00, 0x41, 0x64, 0x6f, 0x62, 0x65,
	0x20, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61, 0x64, 0x79, 0x71, 0xc9, 0x65, 0x3c, 0x00,
	0x00, 0x02, 0xa3, 0x49, 0x44, 0x41, 0x54, 0x78, 0xda, 0x74, 0x93, 0x5f, 0x48, 0x53, 0x51, 0x18,
	0xc0, 0xbf, 0x7b, 0xcf, 0xee, 0xc6, 0x75, 0x5b, 0x56, 0xda, 0x2a, 0x25, 0x6d, 0x6c, 0x06, 0x51,
	0x56, 0xf4, 0x5f, 0x23, 0xb7, 0x22, 0x22, 0xb0, 0x7a, 0x88, 0x41, 0x0f, 0x4b, 0x0c, 0x0a, 0xa3,
	0x82, 0x02, 0x1f, 0x62, 0x1a, 0x45, 0x2f, 0xa1, 0x83, 0x7a, 0x49, 0x2a, 0x86, 0x15, 0x35, 0xff,
	0x64, 0x98, 0x90, 0xa2, 0x91, 0x2e, 0x10, 0x9c, 0x2f, 0x53, 0x92, 0xc2, 0x07, 0xe9, 0xb6, 0x86,
	0x28, 0xda, 0x5d, 0xb6, 0xed, 0x5e, 0x75, 0x0e, 0xd9, 0x9f, 0xd3, 0x39, 0x9b, 0x8d, 0xad, 0xf2,
	0x83, 0xef, 0x9e, 0x7b, 0xbf, 0x3f, 0xbf, 0xf3, 0x7d, 0xdf, 0x39, 0x97, 0xc1, 0x18, 0x43, 0xa6,
	0xd8, 0xed, 0x76, 0xb0, 0xd9, 0x6c, 0x97, 0xcf, 0x9d, 0x3a, 0x7a, 0x23, 0xaf, 0xd0, 0x58, 0x14,
	0x94, 0xe4, 0x65, 0xdf, 0xc4, 0x97, 0x65, 0x25, 0x87, 0x14, 0x9c, 0x56, 0x57, 0xe1, 0x76, 0xbb,
	0xbd, 0x99, 0xf1, 0x8a, 0xcc, 0x8f, 0x96, 0x96, 0x56, 0xf4, 0xba, 0xbd, 0xcd, 0xe1, 0x68, 0xa8,
	0xbd, 0x64, 0x34, 0x1a, 0xe0, 0xb8, 0xe5, 0x1a, 0xdc, 0x6f, 0xb0, 0x43, 0x69, 0x91, 0x16, 0xa4,
	0x90, 0x04, 0xf3, 0x6c, 0xfe, 0x26, 0x12, 0x96, 0x05, 0x60, 0x68, 0x05, 0x35, 0x35, 0x35, 0x10,
	0x5e, 0x8a, 0xc0, 0xf0, 0xe0, 0x40, 0x67, 0xf3, 0x0b, 0xa7, 0xc5, 0xeb, 0x15, 0x40, 0xfc, 0x3a,
	0x02, 0xe6, 0xb3, 0x55, 0xa0, 0xe2, 0x10, 0xc4, 0x7f, 0x8e, 0x83, 0x82, 0xc1, 0x70, 0xfb, 0x51,
	0x57, 0xdd, 0x56, 0xc3, 0xb6, 0xc6, 0x58, 0x2c, 0x4a, 0x36, 0x6b, 0x49, 0x11, 0x28, 0x80, 0x65,
	0x99, 0x63, 0x05, 0xba, 0x75, 0x1d, 0x77, 0x6f, 0x56, 0x61, 0x59, 0x0a, 0x61, 0xcf, 0xe8, 0x27,
	0x7c, 0xe4, 0xd0, 0x5e, 0x7c, 0xa2, 0x6c, 0x27, 0x7e, 0xd3, 0x54, 0x8f, 0x27, 0x5c, 0xcf, 0xb0,
	0x6f, 0xc8, 0x89, 0x2f, 0x5a, 0x4e, 0x0a, 0xeb, 0xf3, 0x36, 0xb0, 0x66, 0xb3, 0x39, 0x99, 0x47,
	0x35, 0xd9, 0x82, 0x26, 0x87, 0x2f, 0x2b, 0xd1, 0x6f, 0x31, 0xed, 0x30, 0x16, 0xc1, 0xc8, 0x87,
	0x56, 0x00, 0xa4, 0x84, 0xca, 0x8a, 0xdd, 0x10, 0x8c, 0xaa, 0x41, 0xa3, 0xd5, 0x42, 0x3c, 0x91,
	0x80, 0x70, 0x64, 0x19, 0xce, 0x57, 0x56, 0x94, 0x08, 0x33, 0xf3, 0xd7, 0x37, 0x17, 0x14, 0x34,
	0x65, 0xb5, 0xc0, 0x30, 0x0c, 0x47, 0xde, 0x4b, 0xaf, 0x5e, 0x38, 0xd3, 0x7f, 0xeb, 0x8a, 0x25,
	0x7f, 0x61, 0x71, 0x09, 0x18, 0x16, 0xc1, 0xb4, 0x5f, 0x06, 0xdd, 0xda, 0x1c, 0x50, 0xa9, 0x38,
	0x60, 0x48, 0x80, 0x4a, 0xc9, 0x81, 0xa3, 0x73, 0x10, 0x1c, 0x6d, 0x3d, 0x8d, 0xf3, 0x52, 0xa8,
	0x8e, 0x02, 0xd8, 0x15, 0x50, 0x94, 0xe8, 0xd8, 0xdb, 0xf7, 0x43, 0x1d, 0x41, 0x39, 0x4c, 0xb0,
	0xa4, 0xef, 0x78, 0x02, 0x8a, 0x37, 0xe6, 0x26, 0x93, 0x02, 0x52, 0x38, 0x09, 0xa3, 0x92, 0xa7,
	0x51, 0xc2, 0x82, 0x2c, 0xf5, 0xfd, 0xa9, 0x80, 0xcd, 0x9c, 0xe8, 0x5c, 0x50, 0xbe, 0xf3, 0xc4,
	0xd9, 0x33, 0xc9, 0xab, 0x94, 0x29, 0x2a, 0x81, 0x90, 0xf9, 0x80, 0x77, 0x4a, 0x04, 0x61, 0x46,
	0x82, 0x68, 0x2c, 0x06, 0xba, 0xfc, 0x5c, 0xea, 0x0a, 0xfc, 0x17, 0x40, 0x44, 0x7a, 0xd5, 0xe5,
	0xb2, 0xf9, 0xa6, 0x45, 0x40, 0x28, 0xe5, 0x8a, 0xc6, 0xe2, 0xb0, 0x6f, 0x7b, 0x31, 0x98, 0x76,
	0x15, 0x26, 0xc3, 0x57, 0xe0, 0xea, 0xd5, 0x00, 0x10, 0x8b, 0xc7, 0x5d, 0xee, 0xd1, 0xf1, 0x25,
	0x5a, 0x7a, 0x7a, 0x50, 0x2c, 0x4b, 0x1f, 0x69, 0x20, 0x0d, 0x5b, 0x15, 0x40, 0x44, 0x16, 0xe7,
	0x82, 0x0b, 0x2c, 0xfb, 0xaf, 0x4b, 0xa1, 0x40, 0xf0, 0xc3, 0x1f, 0xa0, 0x04, 0x7f, 0xd6, 0x4d,
	0x34, 0x18, 0x0c, 0x29, 0x1a, 0x49, 0xfa, 0xee, 0x9b, 0x04, 0x44, 0x56, 0xe6, 0xaf, 0x64, 0x6a,
	0x63, 0x01, 0x43, 0x30, 0x82, 0x91, 0x5a, 0xb3, 0xe6, 0x34, 0x31, 0x35, 0xa7, 0x01, 0x1e, 0x8f,
	0x27, 0x19, 0xa4, 0x25, 0x67, 0x6e, 0xab, 0xab, 0x57, 0x7b, 0x06, 0x7b, 0x79, 0xba, 0x1b, 0x47,
	0x14, 0x21, 0xb2, 0x72, 0x0a, 0xa0, 0x47, 0xfb, 0xb8, 0xb5, 0xef, 0xd7, 0xc7, 0x51, 0xe1, 0xe5,
	0xc1, 0x03, 0xfb, 0x67, 0xb3, 0x2a, 0x28, 0x2f, 0x2f, 0x4f, 0x57, 0x20, 0x8a, 0x62, 0x42, 0x92,
	0xa4, 0xda, 0xa7, 0xed, 0xfd, 0x0f, 0x4d, 0x87, 0xf7, 0xe4, 0xce, 0xce, 0x4c, 0xc3, 0xb8, 0x30,
	0x15, 0x78, 0x37, 0x30, 0xdc, 0x8b, 0x11, 0x7f, 0xaf, 0xba, 0xba, 0x7a, 0xd2, 0x6a, 0xb5, 0x66,
	0xff, 0x4c, 0x82, 0x20, 0x64, 0x56, 0xbb, 0x48, 0xf4, 0xb9, 0xcb, 0x33, 0x61, 0x1d, 0xfb, 0xe6,
	0x8f, 0xf0, 0x3c, 0xff, 0xa0, 0xbb, 0xbb, 0xfb, 0x33, 0xb1, 0x85, 0xf4, 0x7a, 0x3d, 0xbd, 0x74,
	0x59, 0xad, 0xfd, 0x16, 0x60, 0x00, 0xa8, 0x08, 0x0d, 0x8b, 0x03, 0xe7, 0x62, 0xea, 0x00, 0x00,
	0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82, 0x50, 0x4b, 0x01, 0x02, 0x1e, 0x03,
	0x14, 0x00, 0x00, 0x00, 0x08, 0x00, 0x81, 0x61, 0x25, 0x3d, 0xc0, 0xd7, 0xed, 0xc3, 0x19, 0x00,
	0x00, 0x00, 0x1a, 0x00, 0x00, 0x00, 0x08, 0x00, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00,
	0x00, 0x00, 0xa4, 0x81, 0x00, 0x00, 0x00, 0x00, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x74, 0x78, 0x74,
	0x55, 0x54, 0x05, 0x00, 0x03, 0x71, 0xfc, 0x82, 0x4c, 0x75, 0x78, 0x0b, 0x00, 0x01, 0x04, 0xf5,
	0x01, 0x00, 0x00, 0x04, 0x14, 0x00, 0x00, 0x00, 0x50, 0x4b, 0x01, 0x02, 0x1e, 0x03, 0x0a, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x9d, 0x7e, 0x25, 0x3d, 0xfe, 0x31, 0xd5, 0x54, 0x11, 0x03, 0x00, 0x00,
	0x11, 0x03, 0x00, 0x00, 0x14, 0x00, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0xa4, 0x81, 0x5b, 0x00, 0x00, 0x00, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x31, 0x36, 0x78, 0x31, 0x36, 0x2e, 0x70, 0x6e, 0x67, 0x55, 0x54, 0x05, 0x00, 0x03, 0x3a,
	0x30, 0x83, 0x4c, 0x75, 0x78, 0x0b, 0x00, 0x01, 0x04, 0xf5, 0x01, 0x00, 0x00, 0x04, 0x14, 0x00,
	0x00, 0x00, 0x50, 0x4b, 0x05, 0x06, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x02, 0x00, 0xa8, 0x00,
	0x00, 0x00, 0xba, 0x03, 0x00, 0x00, 0x1a, 0x00, 0x54, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20,
	0x61, 0x20, 0x7a, 0x69, 0x70, 0x66, 0x69, 0x6c, 0x65, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2e,
}

var zipdata_zipline = []byte{
	0x50, 0x4b, 0x03, 0x04, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8f, 0xb0, 0x8f, 0x4c, 0x0f,
	0x6f, 0x4f, 0xf3, 0x13, 0x00, 0x00, 0x00, 0x13, 0x00, 0x00, 0x00, 0x08, 0x00, 0x1c, 0x00,
	0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x74, 0x78, 0x74, 0x55, 0x54, 0x09, 0x00, 0x03, 0x3d, 0xa2,
	0xd3, 0x5a, 0x3e, 0xa2, 0xd3, 0x5a, 0x75, 0x78, 0x0b, 0x00, 0x01, 0x04, 0xf6, 0x01, 0x00,
	0x00, 0x04, 0x14, 0x00, 0x00, 0x00, 0x74, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20, 0x61,
	0x20, 0x67, 0x6f, 0x6f, 0x64, 0x20, 0x6f, 0x6e, 0x65, 0x0a, 0x50, 0x4b, 0x03, 0x04, 0x14,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x95, 0xb0, 0x8f, 0x4c, 0x60, 0x41, 0x7b, 0x39, 0x14, 0x00,
	0x00, 0x00, 0x14, 0x00, 0x00, 0x00, 0x84, 0x00, 0x00, 0x00, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x74, 0x6d, 0x70, 0x2f, 0x65,
	0x76, 0x69, 0x6c, 0x2e, 0x74, 0x78, 0x74, 0x74, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20,
	0x61, 0x6e, 0x20, 0x65, 0x76, 0x69, 0x6c, 0x20, 0x6f, 0x6e, 0x65, 0x0a, 0x50, 0x4b, 0x01,
	0x02, 0x1e, 0x03, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8f, 0xb0, 0x8f, 0x4c, 0x0f, 0x6f,
	0x4f, 0xf3, 0x13, 0x00, 0x00, 0x00, 0x13, 0x00, 0x00, 0x00, 0x08, 0x00, 0x18, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xa4, 0x81, 0x00, 0x00, 0x00, 0x00, 0x67, 0x6f,
	0x6f, 0x64, 0x2e, 0x74, 0x78, 0x74, 0x55, 0x54, 0x05, 0x00, 0x03, 0x3d, 0xa2, 0xd3, 0x5a,
	0x75, 0x78, 0x0b, 0x00, 0x01, 0x04, 0xf6, 0x01, 0x00, 0x00, 0x04, 0x14, 0x00, 0x00, 0x00,
	0x50, 0x4b, 0x01, 0x02, 0x14, 0x03, 0x14, 0x00, 0x00, 0x00, 0x00, 0x00, 0x95, 0xb0, 0x8f,
	0x4c, 0x60, 0x41, 0x7b, 0x39, 0x14, 0x00, 0x00, 0x00, 0x14, 0x00, 0x00, 0x00, 0x84, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xa4, 0x81, 0x55, 0x00, 0x00,
	0x00, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e,
	0x2f, 0x74, 0x6d, 0x70, 0x2f, 0x65, 0x76, 0x69, 0x6c, 0x2e, 0x74, 0x78, 0x74, 0x50, 0x4b,
	0x05, 0x06, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x02, 0x00, 0x00, 0x01, 0x00, 0x00, 0x0b,
	0x01, 0x00, 0x00, 0x00, 0x00,
}

func BenchmarkUnzip(b *testing.B) {
	// create zipfile on fs
	f, cleanup, err := u.TempfileWithContent(zipdata_simple)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// create tempdir for dest
	tempdir, err := ioutil.TempDir("", "u")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tempdir)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := u.Unzip(f.Name(), tempdir)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkUnzipBytes(b *testing.B) {
	// create tempdir for dest
	tempdir, err := ioutil.TempDir("", "u")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tempdir)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := u.UnzipBytes(zipdata_simple, tempdir)
		if err != nil {
			b.Error(err)
		}
	}
}
