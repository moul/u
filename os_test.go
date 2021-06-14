package u_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"moul.io/u"
)

func ExampleTempfileWithContent() {
	f, cleanup, err := u.TempfileWithContent([]byte("AAA\nBBB\nCCC"))
	if err != nil {
		panic(err)
	}
	defer cleanup()

	out, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	// Output:
	// AAA
	// BBB
	// CCC
}

func ExampleMustTempfileWithContent() {
	f, cleanup := u.MustTempfileWithContent([]byte("AAA\nBBB\nCCC"))
	defer cleanup()

	out, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	// Output:
	// AAA
	// BBB
	// CCC
}

func ExamplePathExists() {
	file, err := ioutil.TempFile("", "bar")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(file.Name())
	fmt.Println(u.PathExists("/laksjdflkasdjflaksdjfalskdfjasdlfkj")) // should not exist
	fmt.Println(u.PathExists(file.Name()))
	fmt.Println(u.PathExists(filepath.Dir(file.Name())))
	// Output:
	// false
	// true
	// true
}

func ExampleDirExists() {
	file, err := ioutil.TempFile("", "bar")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(file.Name())
	fmt.Println(u.DirExists("/laksjdflkasdjflaksdjfalskdfjasdlfkj")) // should not exist
	fmt.Println(u.DirExists(file.Name()))
	fmt.Println(u.DirExists(filepath.Dir(file.Name())))
	// Output:
	// false
	// false
	// true
}

func ExampleFileExists() {
	file, err := ioutil.TempFile("", "bar")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(file.Name())
	fmt.Println(u.FileExists("/laksjdflkasdjflaksdjfalskdfjasdlfkj")) // should not exist
	fmt.Println(u.FileExists(file.Name()))
	fmt.Println(u.FileExists(filepath.Dir(file.Name())))
	// Output:
	// false
	// true
	// false
}

func ExampleTempFileName() {
	tempname, err := u.TempFileName("", "u")
	if err != nil {
		panic(err)
	}
	if u.FileExists(tempname) {
		panic("there is already one file with tempname")
	}
	f, err := os.Create(tempname)
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempname)
	f.Close()
	fmt.Println("Everything is OK!")
	// Output: Everything is OK!
}

func ExampleMustTempFileName() {
	tempname := u.MustTempFileName("", "u")
	f, _ := os.Create(tempname)
	f.Close()
	os.Remove(tempname)
}

func ExampleCreateEmptyFileWithSize() {
	tempname := u.MustTempFileName("", "u")

	err := u.CreateEmptyFileWithSize(tempname, 42)
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempname)

	fi, err := os.Stat(tempname)
	if err != nil {
		panic(err)
	}
	fmt.Println(fi.Size())

	b, err := ioutil.ReadFile(tempname)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)

	// Output:
	// 42
	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}

func ExampleCurrentUsername() {
	fmt.Println(u.CurrentUsername("fallback"))
}

func TestCurrentUsername(t *testing.T) {
	username := u.CurrentUsername("fallback")
	if username == "fallback" || username == "" {
		t.Errorf("Expected username to set, got %q.", username)
	}
}
